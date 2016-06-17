package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/utils/pagination"
	"ws-personalcollectionmovies/model/session"
)

const COLLECTION = "collection.tpl"
const COLLECTION_TITLE ="Colecciones"

type CollectionController struct {
	beego.Controller
}
func (pController *CollectionController) Get() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = COLLECTION_TITLE
  
	pController.VerifySession()
	

	
	// Servimos la pagina.
	pController.TplName = COLLECTION
}

func (pController *CollectionController) VerifySession() {
    pController.Layout = COLLECTION
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