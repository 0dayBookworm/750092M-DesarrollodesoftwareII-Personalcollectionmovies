package controllers

import (
	"github.com/astaxie/beego"
)


const PROFILE = "profile.tpl"
const PROFILE_TITLE ="Perfil"

type ProfileController struct {
	beego.Controller
}

var (
	_isSessionActive   bool       = false
)

func (pController *ProfileController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	// Damos titulo a la pagina.
	pController.Data["Title"] = PROFILE_TITLE
	// Verificamos la sesión.
	pController.VerifySession()
	
	if (_isSessionActive) {
	
		// Servimos la pagina.
		pController.TplName = "profile.tpl"
	} else {
		// Redireccionamos a inicio.
		pController.Redirect("/", 302)
	}
}

func (pController *ProfileController) VerifySession() {
    pController.Layout = PROFILE
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = "logincontrol.tpl"
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = "logoutcontrol.tpl"
    	 pController.Data["Username"] = sessionVal.(session.UserSession).Username
    	 _isSessionActive = true
    }
}