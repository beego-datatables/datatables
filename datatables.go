package datatables

import (
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
)

type DataTables interface {

	Table(ctx url.Values, tb_name string,columns[] string, search_filter []string) (interface{},error)

}

var RegisterColumns map[string]interface{} = map[string]interface{}{

}

type Data struct {
	Ctx				url.Values //get args
	TableName   	string	//table name
	Columns			[]string //select column
	SearchFilter	[]string //where filter
	//Model            interface{}
}



func (p *Data)Table() (rs interface{}, err error){
	start,_ := strconv.Atoi(p.Ctx.Get("start"))
	length,_ := strconv.Atoi(p.Ctx.Get("length"))
	search := p.Ctx.Get("search[value]")
	order_column,_ := strconv.Atoi(p.Ctx.Get("order[0][column]"))
	order_dir := p.Ctx.Get("order[0][dir]")
	draws,_ := strconv.Atoi(p.Ctx.Get("draw"))


	//query field
	var selectStr string
	for k,v := range p.Columns{
		if k != 0{ selectStr += ","}
		selectStr += v
	}

	//search
	var whereStr  string
	search_len := len(search)

	//offset
	colOffset := start
	qb, _ := orm.NewQueryBuilder("mysql")

	//search
	if search_len >0 {
		for k,v := range p.SearchFilter{
			if k != 0 {
				whereStr += " OR "}
			whereStr +=  v + " LIKE " + "\"%"+search+"%\"" //like
		}
		if order_dir == "asc"{
			qb.Select(selectStr).From(p.TableName).Where(whereStr).OrderBy(p.Columns[order_column]).Asc().Limit(length).Offset(colOffset)
		}else{
			qb.Select(selectStr).From(p.TableName).Where(whereStr).OrderBy(p.Columns[order_column]).Desc().Limit(length).Offset(colOffset)
		}
	}else{
		if order_dir == "asc"{
			qb.Select(selectStr).From(p.TableName).OrderBy(p.Columns[order_column]).Asc().Limit(length).Offset(colOffset)
		}else{
			qb.Select(selectStr).From(p.TableName).OrderBy(p.Columns[order_column]).Desc().Limit(length).Offset(colOffset)
		}
	}
	sql := qb.String()
	o := orm.NewOrm()

	cl := RegisterColumns[p.TableName]
	num, err := o.Raw(sql).QueryRows(cl)



	recordTotal, err := o.QueryTable(p.TableName).Count() //data sum
	var recordsFiltered int32 //search data sum
	type rFilterCont struct {
		CntNumber 		int
	}
	var rcount rFilterCont
	if search_len >0 {
		qb2, _ := orm.NewQueryBuilder("mysql")
		qb2.Select("Count(*) AS cnt_number ").From(p.TableName).Where(whereStr)
		sqlFilter := qb2.String()
		o.Raw(sqlFilter).QueryRow(&rcount)
		recordsFiltered = int32(rcount.CntNumber)
	}else{
		recordsFiltered = int32(recordTotal)
	}

	return map[string]interface{}{
		"draw": int32(draws),
		"recordsTotal": recordTotal,
		"recordsFiltered": recordsFiltered,
		"data":  func() (interface{}) {
		if num == 0{
			rest := []string {}
			return rest
		}else{
			return cl
		}}()},
	err

}
