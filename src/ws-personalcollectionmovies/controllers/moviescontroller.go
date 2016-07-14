package controllers

import (
	"fmt"
 	"strconv"
 	"strings"
 	"github.com/astaxie/beego"
 	
	gotmdb "github.com/ryanbradynd05/go-tmdb"
	
 	"ws-personalcollectionmovies/log"
 	"ws-personalcollectionmovies/error_"
 	"ws-personalcollectionmovies/model/session"
 	"ws-personalcollectionmovies/model/tmdb"
)

type MoviesController struct {
	beego.Controller
}

func (pController *MoviesController) GetUpcomingMovies() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = UPCOMING
	// Verificamos la sesión.
	pController.VerifySession()
	// Obtenemos la pagina solicitada.
	page := pController.Input().Get("Page")
	// Verificamos que sea un entero.
	_, err := strconv.Atoi(page)
	// Verificamos que no haya ocurrido un error al parsear.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Realizamos la busqueda en la API.
 	res, err := tmdb.GetUpcomingMovies(page)
 	
 	if err != nil {
		log.Error(error_.ERR_0040) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Si todo marcho con exito.
	pController.Data["PageHeader"] = UPCOMING_HEAD_TITLE
	// Pagina actual.
	pController.Data["Page"] = res.Page
	// Total de paginas.
	pController.Data["TotalPages"] = res.TotalPages
	// Total de resultados.
	pController.Data["TotalResults"] = res.TotalResults
	// Resultados
	movieShorts := res.Results
	pController.SetMovieDateResults(movieShorts)

	// Actualizamos el paginador.
	pController.SetPaginator(1, res.Page, res.TotalPages, "//personalcollectionmovies-alobaton.c9users.io/movie/upcoming")
	
 	// Servimos la pagina.
	pController.TplName = MOVIES
}

func (pController *MoviesController) GetNowPLayingMovies() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = NOW_PLAYING
	// Verificamos la sesión.
	pController.VerifySession()
	// Obtenemos la pagina solicitada.
	page := pController.Input().Get("Page")
	// Verificamos que sea un entero.
	_, err := strconv.Atoi(page)
	// Verificamos que no haya ocurrido un error al parsear.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Realizamos la busqueda en la API.
 	res, err := tmdb.GetMovieNowPlaying(page)
 	
 	if err != nil {
		log.Error(error_.ERR_0040) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Si todo marcho con exito.
	pController.Data["PageHeader"] = NOW_PLAYING_HEAD_TITLE
	// Pagina actual.
	pController.Data["Page"] = res.Page
	// Total de paginas.
	pController.Data["TotalPages"] = res.TotalPages
	// Total de resultados.
	pController.Data["TotalResults"] = res.TotalResults
	// Resultados
	movieShorts := res.Results
	pController.SetMovieDateResults(movieShorts)

	// Actualizamos el paginador.
	pController.SetPaginator(1, res.Page, res.TotalPages, "//personalcollectionmovies-alobaton.c9users.io/movie/nowplaying")
	
 	// Servimos la pagina.
	pController.TplName = MOVIES
}

func (pController *MoviesController) GetPopularMovies() {
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = POPULAR
	// Verificamos la sesión.
	pController.VerifySession()
	// Obtenemos la pagina solicitada.
	page := pController.Input().Get("Page")
	// Verificamos que sea un entero.
	_, err := strconv.Atoi(page)
	// Verificamos que no haya ocurrido un error al parsear.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Realizamos la busqueda en la API.
 	res, err := tmdb.GetMovieNowPlaying(page)
 	
 	if err != nil {
		log.Error(error_.ERR_0040) 
		pController.Redirect("404", 302)
    	pController.StopRun()
	}
	// Si todo marcho con exito.
	pController.Data["PageHeader"] = POPULAR_HEAD_TITLE
	// Pagina actual.
	pController.Data["Page"] = res.Page
	// Total de paginas.
	pController.Data["TotalPages"] = res.TotalPages
	// Total de resultados.
	pController.Data["TotalResults"] = res.TotalResults
	// Resultados
	movieShorts := res.Results
	pController.SetMovieDateResults(movieShorts)

	// Actualizamos el paginador.
	pController.SetPaginator(1, res.Page, res.TotalPages, "//personalcollectionmovies-alobaton.c9users.io/movie/popular")
	
 	// Servimos la pagina.
	pController.TplName = MOVIES
}


// VerifySession [Metodo encargado de verificar la sesión y actualizar el control de login/logout]
func (pController *MoviesController) VerifySession() {
    pController.Layout = MOVIES
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
// SetPaginator [Metodo encargado de actualizar el paginador del template movies.tpl]
func (pController *MoviesController) SetPaginator(pMinPage, pPage, pTotalPages int, pUrl string) {
	before := `<li><a href="`+pUrl+`?Page=%d"><i class="glyphicon glyphicon-chevron-left"></i> Anterior</a></li>`
	next := `<li><a href="`+pUrl+`?Page=%d">Siguiente <i class="glyphicon glyphicon-chevron-right"></i></a></li>`
	var paginator string
	if (pPage == pMinPage) {
		paginator = fmt.Sprintf(next, pPage+1)
	} else if (pPage == pTotalPages) {
		paginator = fmt.Sprintf(before, pPage-1)
	} else {
		paginator = fmt.Sprintf(before, pPage-1)+fmt.Sprintf(next, pPage+1)
	}
	pController.Data["Paginator"] = paginator
}

func (pController *MoviesController) SetMovieDateResults(pMovieShorts []gotmdb.MovieShort) {
	movieShortsContent := ""
	for i := range pMovieShorts {
	    // Por cada resultado armamos el div html.
	    // Creamos nuestros contenedores
	    item :=  `<a href="%s" class="item col-xs-6 col-lg-6"> %s </a>`
	    thumbnail := `<div class="thumbnail row"> %s </div>`
	    img := `<img class="pull-left" src="%s" alt="Generic placeholder thumbnail" />`
	    caption := `<div class="caption"> %s </div>`
	    title := `<h3 class="list-group-item-heading"> %s </h3>`
	    originalTitle := `<h5 class=""> %s </h5>`
	    releaseDate := `<p class="list-group-item-text"> %s <i class="glyphicon glyphicon-calendar"></i></p>`
	    voteAverage := `<p class=""> %g <i class="glyphicon glyphicon-star"></i> </p>`
	    
	    title = fmt.Sprintf(title, pMovieShorts[i].Title)
	    originalTitle = fmt.Sprintf(originalTitle, pMovieShorts[i].OriginalTitle)
	    releaseDate = fmt.Sprintf(releaseDate, pMovieShorts[i].ReleaseDate)
	    voteAverage = fmt.Sprintf(voteAverage, pMovieShorts[i].VoteAverage)
	    
	    caption = fmt.Sprintf(caption, title+originalTitle+releaseDate+voteAverage)
	    
	    if strings.Compare(pMovieShorts[i].PosterPath, "") == 0 {
	    	img = fmt.Sprintf(img, "http://personalcollectionmovies-alobaton.c9users.io/public/images/image_no_available.png")
	 	} else {
	 		img = fmt.Sprintf(img, "https://image.tmdb.org/t/p/w500"+pMovieShorts[i].PosterPath)
	 	}
	 	
	    thumbnail = fmt.Sprintf(thumbnail, img+caption)
	    
	    item = fmt.Sprintf(item, "http://personalcollectionmovies-alobaton.c9users.io/movie?ID="+strconv.Itoa(pMovieShorts[i].ID), thumbnail)
	    
	    movieShortsContent += item
	}
	// Ponemos los resultados en la pagina.
	pController.Data["MovieShorts"]=movieShortsContent
}