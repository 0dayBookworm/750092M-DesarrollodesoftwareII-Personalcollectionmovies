package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/model/session"
)

const INDEX = "index.tpl"
const INDEX_TITLE ="Inicio"

type HomeController struct {
	beego.Controller
}
func (pController *HomeController) Get() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = INDEX_TITLE
  
	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = INDEX
}

func (pController *HomeController) VerifySession() {
    pController.Layout = INDEX
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