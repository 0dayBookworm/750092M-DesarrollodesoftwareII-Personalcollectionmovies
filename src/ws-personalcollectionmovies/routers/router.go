package routers

import (
	"ws-personalcollectionmovies/controllers"
	"github.com/astaxie/beego"
)

// Aquí inicializamos todas las rutas del servicio.
func init() {
    // Index or home.
    beego.Router("/", &controllers.HomeController{})
}
