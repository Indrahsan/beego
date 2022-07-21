package controllers

import (
	"beego/models"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type BukuController struct {
	beego.Controller
}

func (u *BukuController) GetAll() {
	buku, err := models.GetAllBuku()
	result := models.Response{}

	if err == nil {
		result.Code = 0
		result.Info = "Berhasil"
		result.Data = buku
		u.Data["json"] = result
		u.ServeJSON()
	} else {
		result.Code = 0
		result.Info = err.Error()
	}
	
u.Data["json"] = result
u.ServeJSON()

}
