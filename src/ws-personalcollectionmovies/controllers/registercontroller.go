package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
func (pController *RegisterController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	
	// Servimos la pagina.
	pController.TplName = "register.html"
}