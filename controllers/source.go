package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type SourceController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
//func (u *SourceController) Post() {
//	var user models.User
//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
//	uid := models.AddUser(user)
//	u.Data["json"] = map[string]string{"uid": uid}
//	u.ServeJSON()
//}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
//func (u *SourceController) GetAll() {
//	sources := models.GetAllSources()
//	u.Data["json"] = sources
//	u.ServeJSON()
//}


//@router /:sid [get]
func (u *SourceController) Get() {
	sid := u.GetString(":sid")
	u.Data["json"] = sid
	u.ServeJSON()
	//sid := u.GetString(":sid")
	//if sid != "" {
	//	source := models.GetSource(1)
	//	//if err != nil {
	//	//	u.Data["json"] = err.Error()
	//	//} else {
	//	//	u.Data["json"] = source
	//	//}
	//	u.Data["json"] = source
	//}
	//u.ServeJSON()
}

//@router /index/:sid [get]
func (u *SourceController) Index() {
	sid := u.GetString(":sid")
	u.Data["json"] = sid
	u.ServeJSON()
}

//@router /lock [get]
func (u *SourceController) Lock()  {
	u.Data["json"] = "lock"
	u.ServeJSON()
}


