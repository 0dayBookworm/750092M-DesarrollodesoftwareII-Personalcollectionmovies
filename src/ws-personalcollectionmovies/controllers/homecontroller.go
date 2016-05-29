package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}
func (pController *HomeController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	
	// Servimos la pagina.
	pController.TplName = "home.html"
}