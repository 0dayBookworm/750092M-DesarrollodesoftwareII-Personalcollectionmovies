package routers

import (
    "net/http"
    "text/template"
	"ws-personalcollectionmovies/controllers"
	"ws-personalcollectionmovies/log"
	"github.com/astaxie/beego"
)

// Aquí inicializamos todas las rutas del servicio.
func init() {
    beego.ErrorHandler("404", page_not_found)
    // Index or home.
    beego.Router("/", &controllers.HomeController{})
    // Modulo: Registro de cuentas de usuario.
    initRegistrationModule()
    // Modulo: Perfil de usuario.
    initProfileModule()
    // Modulo: Inicio y cierre de sesión.
    initLoginModule()
    // Modulo: Peliculas
    initMovieModule()
    // Modulo: Colecciones.
    initCollectionsModule()
    // Modulo: Contacto.
    initContactModule()
    // ErrorHandler
   
}
// Modulo: Registro de cuentas de usuario.
func initRegistrationModule() {
    // Formulario de registro.
    beego.Router("/register", &controllers.RegisterController{})
    // Crear un usuario.
    beego.Router("/register/create", &controllers.RegisterController{}, "post:CreateUseraccount")
}
// Modulo: Perfil de usuario.
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
// Modulo: Inicio y cierre de sesión.
func initLoginModule() {
    // Iniciar sesión.
    beego.Router("/login", &controllers.SessionController{}, "post:Login")
    // Cerrar sesión.
    beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
}
// Modulo: Peliculas
func initMovieModule() {
    // Ver una pelicula
    beego.Router("/movie", &controllers.MovieController{}, "get:GetMovie")
    // Busqueda de peliculas.
    beego.Router("/movie/search", &controllers.MovieController{}, "get:Search")
    // Estrenos
    beego.Router("/movie/upcoming", &controllers.MoviesController{}, "get:GetUpcomingMovies")
    // Cartelera
    beego.Router("/movie/nowplaying", &controllers.MoviesController{}, "get:GetNowPLayingMovies")
}
// Modulo: Colecciones.
func initCollectionsModule() {
    beego.Router("/collection", &controllers.CollectionController{})
}
// Modulo: Contacto.
func initContactModule() {
    beego.Router("/contact", &controllers.ContactController{})
}
// page_not_found Sobreescribimos este método para que beego muestre la pagina de error 404 personalizada.
func page_not_found(rw http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles(beego.BConfig.WebConfig.ViewsPath + "404.tpl")
    err := t.Execute(rw, nil)
    if err!=nil {
        log.Info("router.go "+err.Error())
        panic(err)
    }
}