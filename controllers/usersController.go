package controllers

import (
	"encoding/json"
	"beego/models"
	"beego/utils"
	beego "github.com/beego/beego/v2/server/web"
)

// UserController represents controller for user API
type UsersController struct {
	beego.Controller
}

func (uc *UsersController) AddNewUser() {
	result := models.Response{}
	u := models.Users{}
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	u.Password, _ = utils.HashPassword(u.Password)
	user := models.InsertOneUser(u)
	result.Code = 0
	result.Info = "Berhasil"
	result.Data = user
	uc.Data["json"] = result
	uc.ServeJSON()
}