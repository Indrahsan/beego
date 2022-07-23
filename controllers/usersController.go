package controllers

import (
	"encoding/json"
	"beego/models"
	"beego/utils"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego/validation"
	"log"
	"strconv"
)

// UserController represents controller for user API
type UsersController struct {
	beego.Controller
}

func (uc *UsersController) AddNewUser() {
	result := models.Response{}
	u := models.Users{}
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	valid := validation.Validation{}
	valid.Required(u.Nama, "Name")
	valid.Required(u.Password, "Password")
	valid.Required(u.Username, "Username")
	if valid.HasErrors() {
        // validation does not pass
        // print invalid message
		
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
			result.Code = 10	
			result.Info = err.Key + err.Message
			uc.Data["json"] = result
			uc.ServeJSON()
        }
	} else {
	u.Password, _ = utils.HashPassword(u.Password)
	user := models.InsertOneUser(u)
	result.Code = 0
	result.Info = "Berhasil"
	result.Data = user
	uc.Data["json"] = result
	uc.ServeJSON()
	}
}

func (uc *UsersController) DeleteUser() {
	result := models.Response{}
	// get id from query string and convert it to int
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))

	// delete user
	deleted := models.DeleteUsers(id)

	result.Code = 0
	result.Info = "Berhasil delete!"
	result.Data = deleted
	// generate response
	uc.Data["json"] =  result
	uc.ServeJSON()
}