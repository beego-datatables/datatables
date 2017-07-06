#### beego  datatables

[beego](https://github.com/astaxie/beego/) MVC  [datatables](http://datatables.net/examples/server_side/pipeline.html) plugins

###### Download and install
`go get "github.com/beego-datatables/datatables"`

###### Usage

routers
```go
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
```go
package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"github.com/beego-datatables/datatables"
	".../models"
)

type Example struct {

	beego.Controller
	
}

func (c *Example) AjaxData(){

	var Qtab datatables.Data
	Qtab.Ctx = c.Input() //datatables get
	Qtab.DBName = "default"
	Qtab.TableName = "example_record" //modles tables name
	Qtab.Columns = []string{"id","user_name","operation","action","result","create_time"} //datatables columns arrange
	Qtab.Order = []string{"","user_name","","action","result","create_time"} 
	Qtab.SearchFilter = []string{"user_name","operation","action","result"} //datatables filter
	datatables.RegisterColumns[Qtab.TableName] = new([]models.ExampleRecord) //register result 

	rs , _ := Qtab.Table()


	c.Data["json"] = rs

	c.ServeJSON()
	
}

```


models

```go
type ExampleRecord struct {

	Id				int
	User				*User 		`orm:"rel(fk);null;on_delete(set_null)"`
	Operation 			string
	Action 				string
	Result 				string		`orm:"type(text)"`
	CreateTime 			time.Time 	`orm:"auto_now_add;type(datetime)"`
	
}
```

datatables ajax

```javascript
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
            { "data": "operation" },
            { "data": "action" },
            { "data": "result" }
        ],
        ...
       })
```
#### LICENSE

MIT
