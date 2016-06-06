package routers

import (
	"ws-personalcollectionmovies/controllers"
	"github.com/astaxie/beego"
)

// Aquí inicializamos todas las rutas del servicio.
func init() {
    // Index or home.
    beego.Router("/", &controllers.HomeController{})
    // Modulo: Perfil de usuario.
    initProfileModule()
    // Modulo: Registro de cuentas de usuario.
    initRegistrationModule()
}


func initRegistrationModule() {
    // Formulario de registro.
    beego.Router("/register", &controllers.RegisterController{})
    // Crear un usuario.
    beego.Router("/register/create", &controllers.RegisterController{}, "post:CreateUseraccount")
}

func initProfileModule() {
    // Datos de usuario.
    beego.Router("/profile", &controllers.ProfileController{})
}
