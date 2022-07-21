package utils

import(
	"golang.org/x/crypto/bcrypt"
	"github.com/beego/beego/v2/client/orm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConnectionDB() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/golang") 
}