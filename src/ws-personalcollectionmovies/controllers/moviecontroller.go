package controllers

import (
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
)

const MOVIE = "movie.tpl"
const MOVIE_TITLE ="Pelicula"

type MovieController struct {
	beego.Controller
}

func (pController *MovieController) GetMovie() {
	// Verificamos la sesión.
	pController.VerifySession()
	
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

func (pController *MovieController) VerifySession() {
    pController.Layout = MOVIE
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = "logincontrol.tpl"
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = "logoutcontrol.tpl"
    	 pController.Data["Username"] = sessionVal.(session.UserSession).Username
    }
}
func (pController *MovieController) ServeMessage(pErrorCode, pErrorMessage string) {
	searchMovieResponse := wsinterface.SearchMovieResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &searchMovieResponse
    pController.ServeJSON()
}
