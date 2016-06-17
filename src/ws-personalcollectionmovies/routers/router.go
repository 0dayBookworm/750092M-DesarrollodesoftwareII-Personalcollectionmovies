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
    // Modulo: Errores
    initErrorModule()
    // Modulo: Peliculas
    initMovieModule()
    
    initCollectionsModule()
}


func initRegistrationModule() {
    // Formulario de registro.
    beego.Router("/register", &controllers.RegisterController{})
    // Crear un usuario.
    beego.Router("/register/create", &controllers.RegisterController{}, "post:CreateUseraccount")
}

func initProfileModule() {
    // Formulario
    beego.Router("/profile", &controllers.ProfileController{})
    // Actualizar información
    beego.Router("profile/update", &controllers.ProfileController{}, "post:Update")
    // Cambiar contraseña.
    beego.Router("/profile/password", &controllers.ProfileController{}, "post:ChangePassword")
    // Eliminar cuenta de usuario.
    beego.Router("/profile/delete", &controllers.ProfileController{}, "post:Delete")
}

func initLoginModule() {
    // Iniciar sesión.
    beego.Router("/login", &controllers.SessionController{}, "post:Login")
    // Cerrar sesión.
    beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
}

func initErrorModule() {
    beego.Router("/error", &controllers.ErrorController{}, "get:HandleError")
}

func initMovieModule() {
    beego.Router("/movie", &controllers.MovieController{}, "get:GetMovie")
}

func initCollectionsModule() {
    beego.Router("/collection", &controllers.CollectionController{})
}