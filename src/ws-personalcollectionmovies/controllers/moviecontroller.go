package controllers

import (
	// "fmt"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	
	// gotmdb "github.com/ryanbradynd05/go-tmdb"
	
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/model/session"
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/model/tmdb"
	"ws-personalcollectionmovies/model/google"
	"ws-personalcollectionmovies/model/database/domain"
	"ws-personalcollectionmovies/model/database/connection"
)

type MovieController struct {
	beego.Controller
}

func (pController *MovieController) GetMovie() {
	// Verificamos la sesión.
	pController.Layout = MOVIE
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = LOGIN_CONTROL
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = LOGOUT_CONTROL
    	 pController.Data["Username"] = sessionVal.(session.UserSession).Username
    }
	
	request := wsinterface.GetMovieRequest{}
	err := pController.ParseForm(&request)
	// Verificamos que no haya ocurrido un error al parsear.
	// En caso de que hayan ocurrido errores debemos retornar un tpl informando que no hay resultados para la busqueda.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	
 	// Realizamos la busqueda primero en cache.
	
	
 	// Realizamos la busqueda en la API.
 	res, err := tmdb.GetMovieInfo(request.ID)
 	
 	if err != nil {
 		log.Error(error_.ERR_0040) 
 		pController.Redirect("404", 302)
     	pController.StopRun()
 	}
 	// No es necesario verificar si hubieron resultados para la busqueda.
	
 	// Una vez obtenidos los datos editamos el template.
 	pController.Data["Title"]=strings.ToUpper(res.Title)
 	pController.Data["OriginalTitle"]=strings.ToUpper(res.OriginalTitle)
	
 	if strings.Compare(res.PosterPath, "") == 0 {
		pController.Data["Poster"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/image_no_available.png"
 	} else {
 		pController.Data["Poster"]="https://image.tmdb.org/t/p/w500"+res.PosterPath
 	}
	
	pController.Data["Runtime"]=res.Runtime
	
 	releaseDateData := strings.Split(res.ReleaseDate, "-")
 	if len(releaseDateData) < 3 {
 		pController.Data["Released"]=res.ReleaseDate
 	} else {
 		pController.Data["Day"]=releaseDateData[2]
 		month, _ := strconv.Atoi(releaseDateData[1])
 		pController.Data["Month"]=util.GetMonth(month)
 		pController.Data["Year"]=strings.ToUpper(releaseDateData[0])
 	}
	
	// Ponemos las compañias productoras.
	productionCompanies := "<br>"
	for i := range res.ProductionCompanies {
		productionCompanies += res.ProductionCompanies[i].Name+"<br>"
	}
 	pController.Data["ProductionCompanies"]=productionCompanies

 	// Ponemos los generos.
	genres := "<br>"
	for i := range res.Genres {
		genres += `<span class="label label-default">`+res.Genres[i].Name+`</span><br>`
	}
 	pController.Data["Genres"]=genres
 	
 	pController.Data["Website"]=res.Homepage
 	
 	pController.Data["Plot"]=res.Overview
 	// Obtenemos el trailer.
 	// Verificamos que haya trailer.
 	if (len(res.Videos.Results) > 0) {
 		// Obtenemos el primer video.
 		trailerSrc := res.Videos.Results[0].Key
 		pController.Data["Trailer"]=trailerSrc
 	}
 	
 	// Actualizamos los controles de compartición.
 	pController.Data["PageUrl"] =PAGE_URI+strconv.Itoa(res.ID)
	
	// Actualizamos los controles de las listas y colecciones.
	if sessionVal != nil {
		// Verificamos si la pelicula ya ha sido añadida a la lista de peliculas por ver.
		_, err = domain.UserNextviewsMovieByUsernameID(connection.GetConn(), sessionVal.(session.UserSession).Username, strconv.Itoa(res.ID))
		if err != nil {
			pController.Data["WatchListContent"]=`<a id="WatchListAdd"> <small> <i class="glyphicon glyphicon-eye-open"></i> Añadir a Peliculas por ver</small> </a>`
		} else {
			pController.Data["WatchListContent"]=`<a id="WatchListRemove"> <small style="color: #3F81BF"> <i class="glyphicon glyphicon-eye-open"></i> Remover de Peliculas por ver</small> </a>`
		}
		// Verificamos si la pelicula ya ha sido añadida a la lista de películas vistas.
		viewMovie, err := domain.UserViewMovieByUsernameID(connection.GetConn(), sessionVal.(session.UserSession).Username, strconv.Itoa(res.ID))
		if err != nil {
			pController.Data["WasView"]=false
			pController.Data["ViewListContent"]=`<a id="ViewListAdd"> <small> <i class="glyphicon glyphicon-eye-close"></i>  Añadir a Vistas </small> </a>`
		} else {
			pController.Data["WasView"]=true
			pController.Data["ViewListContent"]=`<a id="ViewListRemove"> <small style="color: #3F81BF"> <i class="glyphicon glyphicon-eye-close"></i>  Remover de Vistas </small> </a>`
			pController.Data["PlaceDescription"]=viewMovie.Place
			pController.Data["CinemaIframe"]=`<iframe id="CinemaIframe" width="100%" height="300px" frameborder="0" style="border:0" src="`+google.GetEmbedSource(viewMovie.Place)+`" allowfullscreen></iframe>`
		}
	} else {
		pController.Data["WatchListContent"]=`<a id="WatchListAdd"> <small id="WatchListIndicator"> <i class="glyphicon glyphicon-eye-open"></i> Añadir a Peliculas por ver</small> </a>`
		pController.Data["ViewListContent"]=`<a id="ViewListAdd"> <small> <i class="glyphicon glyphicon-eye-close"></i>  Añadir a Vistas </small> </a>`
	}
	
	// Servimos la pagina.
	pController.TplName = MOVIE
}

func (pController *MovieController) Search() {
	request := wsinterface.SearchMovieRequest{}
	err := pController.ParseForm(&request)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	pController.StopRun()
	}
	res, err := tmdb.SearchMovie(request.Title, request.Year)
	
	if err != nil {
		log.Error(error_.ERR_0040) 
		pController.ServeMessage(error_.KO, error_.ERR_0040+err.Error())
    	pController.StopRun()
	}
	pController.Data["json"] =  res
    pController.ServeJSON()
}

func (pController *MovieController) ServeMessage(pErrorCode, pErrorMessage string) {
	searchMovieResponse := wsinterface.SearchMovieResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &searchMovieResponse
    pController.ServeJSON()
}
