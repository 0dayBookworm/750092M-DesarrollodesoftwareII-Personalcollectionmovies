package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/model/session"
)

const MOVIE = "movie.tpl"
const MOVIE_TITLE ="Pelicula"

type MovieController struct {
	beego.Controller
}
func (pController *MovieController) GetMovie() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = MOVIE_TITLE
  
	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = MOVIE
}

func (pController *MovieController) VerifySession() {
    pController.Layout = MOVIE
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