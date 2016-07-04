package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/model/session"
)

type TermsController struct {
	beego.Controller
}
func (pController *TermsController) Get() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = TERMS_TITLE
  
	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = TERMS
}

func (pController *TermsController) VerifySession() {
    pController.Layout = TERMS
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