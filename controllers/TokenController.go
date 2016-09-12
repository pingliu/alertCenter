package controllers

import (
	"alertCenter/core/db"
	"alertCenter/core/service"
	"alertCenter/models"
	"alertCenter/util"
)

type TokenController struct {
	BaseController
}

func (e *TokenController) AddToken() {
	user := e.GetSession("user")
	if user == nil {
		e.Data["json"] = util.GetErrorJson("please certification")
		e.ServeJSON()
	} else {
		u := user.(*models.User)
		projectName := e.GetString("projectName")
		if projectName == "" {
			e.Data["json"] = util.GetErrorJson("projectName can't be empty")
			e.ServeJSON()
		} else {
			service := &service.TokenService{
				Session: db.GetMongoSession(),
			}
			token := service.CreateToken(projectName, u.Name)
			e.Data["json"] = util.GetSuccessReJson(token)
			e.ServeJSON()
		}
	}
}