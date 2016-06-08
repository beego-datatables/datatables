#### beego  datatables

[beego](https://github.com/astaxie/beego/) MVC  [datatables](http://datatables.net/examples/server_side/pipeline.html) plugins

###### Download adn install
`go get  https://github.com/beego-datatables/datatables`

###### Example

routers
```
package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/datatables/data/json", &controllers.OperationRecord{},"*:AjaxData")
	
}
```

beego controllers:
```
package controllers

type Example struct {

	beego.Controller
	
}

func (c *Example) AjaxData(){

	var Qtab datatables.Data
	Qtab.Ctx = c.Input() //datatables get

	Qtab.TableName = "example" //modles tables name
	Qtab.Columns = []string{"id","user_name","operation","action","result","create_time"} //datatables columns arrange
	Qtab.SearchFilter = []string{"user_name","operation","action","result"} //datatables filter
	datatables.RegisterColumns[Qtab.TableName] = new([]models.ExampleRecord) //register result 

	rs , rserr := Qtab.Table()

	c.Data["json"] = rs

	c.ServeJSON()
	
}

```


models

```
type ExampleRecord struct {

	Id					int
	User				*User 		`orm:"rel(fk);null;on_delete(set_null)"`
	Operation 			string
	Action 				string
	Result 				string		`orm:"type(text)"`
	CreateTime 			time.Time 	`orm:"auto_now_add;type(datetime)"`
	
}
```

datatables ajax

```
    var example_table = $('#table').DataTable({
        ...
        "processing": true,
        "serverSide": true,
        "ajax": '/datatables/data/json',
        "aLengthMenu": [[5,15, 25, 50, 100, 200, 500], ['5','15', '25', '50', '100', '200', '500']],
        "iDisplayLength": 15,
        "columns": [
            { "data": "id" },
            { "data": "username" },
            { "data": "execute" },
            { "data": "action" },
            { "data": "result" }
        ],
        ...
       })
```
