package controllers

import (
	"github.com/astaxie/beego"
)

type ProfileController struct {
	beego.Controller
}
func (pController *ProfileController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	// Damos titulo a la pagina.
	pController.Data["Title"] = "Perfil"
	
	pController.Data["Username"] = "alobaton"
	pController.Data["Email"] = "alobaton.restrepo@gmail.com"
	// Servimos la pagina.
	pController.TplName = "profile.html"
}