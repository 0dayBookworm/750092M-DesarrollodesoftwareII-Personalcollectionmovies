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
    // ErrorHandler
    beego.ErrorHandler("404", page_not_found)
    // Index or home.
    beego.Router("/", &controllers.HomeController{})
    // Terminos y condiciones.
    beego.Router("/terms", &controllers.TermsController{})
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
    beego.Router("/account/profile", &controllers.ProfileController{})
    
    beego.Router("/account/security", &controllers.ProfileController{}, "get:Security")
    // Actualizar información
    beego.Router("/account/profile/update", &controllers.ProfileController{}, "post:Update")
    // Cambiar contraseña.
    beego.Router("/account/profile/password", &controllers.ProfileController{}, "post:ChangePassword")
    // Eliminar cuenta de usuario.
    beego.Router("/account/profile/delete", &controllers.ProfileController{}, "post:Delete")
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
    // Populares
    beego.Router("/movie/popular", &controllers.MoviesController{}, "get:GetPopularMovies")
}
// Modulo: Colecciones.
func initCollectionsModule() {
    // Peliculas vistas.
    // Peliculas por Ver.
    beego.Router("/account/watchlist", &controllers.CollectionController{}, "get:Watchlist")
    // Añadir y remover de peliculas por ver.
    beego.Router("/account/watchlist/add", &controllers.CollectionController{}, "get:WatchListAdd")
    // beego.Router("/account/watchlist/remove", &controllers.CollectionController{}, "post:WatchlistRemove")
    // Mis colecciones.
    beego.Router("/account/collection", &controllers.CollectionController{}, "get:Get")
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