package controllers


import (
// 	"strings"
// 	"fmt"
	"github.com/astaxie/beego"
	
// 	"ws-personalcollectionmovies/log"
// 	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/session"
	"ws-personalcollectionmovies/model/wsinterface"
)

const CONTACT = "contact.tpl"
const CONTACT_TITLE ="Contactanos"

type ContactController struct {
	beego.Controller
}

func (pController *ContactController) Get() {
    // Inertamos el titulo a la pagina.
	pController.Data["Title"] = CONTACT_TITLE
	// Verificamos la sesión.
	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = CONTACT
}


func (pController *ContactController) VerifySession() {
    pController.Layout = CONTACT
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
func (pController *ContactController) ServeMessage(pErrorCode, pErrorMessage string) {
	contactResponse := wsinterface.ContactResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &contactResponse
    pController.ServeJSON()
}