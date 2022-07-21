package utils

import(
	"github.com/beego/beego/v2/client/orm"
)

func ConnectionDB() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/golang") 
}