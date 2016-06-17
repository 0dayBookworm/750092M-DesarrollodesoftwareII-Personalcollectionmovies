package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/model/session"
)

const ERROR = "error.tpl"
const ERROR_TITLE ="Error"

type ErrorController struct {
	beego.Controller
}
func (pController *ErrorController) HandleError() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = ERROR_TITLE
  
	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = ERROR
}

func (pController *ErrorController) VerifySession() {
    pController.Layout = ERROR
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = "logincontrol.tpl"
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = "logoutcontrol.tpl"
    	 pController.Data["Username"] = sessionVal.(session.UserSession).Username
    }
}