package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}
func (pController *HomeController) Get() {
	// Servimos la pagina.
	pController.TplName = "index.html"
}

