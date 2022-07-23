package models

import(
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"beego/utils"
)

type Users struct {
	Id 		int
	Nama    string `json:"nama" orm:"null" form:"nama"`
	Username string `json:"username" orm:"null;unique" form:"username"`
	Password string `json:"password" orm:"null" form:"password"`
}

func init() {
	utils.ConnectionDB()
	orm.RegisterModel(new(Users))
}

func (u *Users) TableName() string {
    // db table name
    return "user"
}

func InsertOneUser(user Users) *Users {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	// get prepared statement
	i, _ := qs.PrepareInsert()

	var u Users

	// Insert
	id, err := i.Insert(&user)
	if err == nil {
		// successfully inserted
		u = Users{Id: int(id)}
		err := o.Read(&u)
		if err == orm.ErrNoRows {
			return nil
		}
	} else {
		return nil
	}

	return &u
}

func DeleteUsers(id int) bool {
	o := orm.NewOrm()
	_, err := o.Delete(&Users{Id: id})
	if err == nil {
		// successfull
		return true
	}

	return false
}