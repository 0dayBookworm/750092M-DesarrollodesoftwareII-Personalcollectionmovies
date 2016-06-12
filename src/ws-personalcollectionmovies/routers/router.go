package routers

import (
	"ws-personalcollectionmovies/controllers"
	"github.com/astaxie/beego"
)

// Aquí inicializamos todas las rutas del servicio.
func init() {
    // Index or home.
    beego.Router("/", &controllers.HomeController{})
    // Modulo: Inicio y cierre de sesión.
    initLoginModule()
    // Modulo: Registro de cuentas de usuario.
    initRegistrationModule()
    // Modulo: Perfil de usuario.
    initProfileModule()
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

func initLoginModule() {
    // Iniciar sesión.
    beego.Router("/login", &controllers.SessionController{}, "post:Login")
}
