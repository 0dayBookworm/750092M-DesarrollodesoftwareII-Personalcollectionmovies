package controllers

import(
    "fmt"
	"strings"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/model/wsinterface"
    )

// Templates
const INDEX = "index.tpl"
const TERMS ="termsandconditions.tpl"
const REGISTER = "register.tpl"
const ACCOUNT = "account.tpl"
const PROFILE = "profile.tpl"
const MOVIE = "movie.tpl"
const MOVIES = "movies.tpl"
const COLLECTIONS = "collections.tpl"
const COLLECTION = "collection.tpl"
const CONTACT = "contact.tpl"
const SECURITY="password.tpl"
const COMMON_ACCOUNT_CONTROL="commonaccountcontrol.tpl"
const AUDIT_ACCOUNT_CONTROL="auditaccountcontrol.tpl"
const COMMON_NAV="commonnav.tpl"
const AUDIT_NAV="auditnav.tpl"
const LOGOUT_CONTROL="logoutcontrol.tpl"
const LOGIN_CONTROL="logincontrol.tpl"
const LIST ="list.tpl"
// Titulos
const INDEX_TITLE ="Inicio"
const REGISTER_TITLE = "Registrate"
const MOVIE_TITLE ="Pelicula"
const COLLECTION_TITLE ="Mis colecciones"
const UPCOMING ="Estrenos"
const UPCOMING_HEAD_TITLE = "PROXIMOS ESTRENOS"
const NOW_PLAYING = "Cartelera"
const NOW_PLAYING_HEAD_TITLE = "EN CARTELERA HOY"
const POPULAR = "Populares"
const POPULAR_HEAD_TITLE = "PELÍCULAS POPULARES"
const CONTACT_TITLE ="Contactanos"
const TERMS_TITLE ="Términos y Condicones"
// Peliculas por ver.
const WATCHLIST_TITLE ="PELÍCULAS POR VER"
// Peliculas vistas
const VIEWLIST_TITLE = "PELÍCULAS VISTAS"
// Paginas
const PAGE_URI = "https://personalcollectionmovies-alobaton.c9users.io/movie?ID="


func ConcatenateMovieShortResult(pMyMovieShort wsinterface.MyMovieShort) string{
	// Parseamos a un MovieShort personalizado con el fin de reutilizar código.
	item :=  `<a href="%s" class="item col-xs-6 col-lg-6"> %s </a>`
	thumbnail := `<div class="thumbnail row"> %s </div>`
	img := `<img class="pull-left" src="%s" alt="Generic placeholder thumbnail" />`
	caption := `<div class="caption"> %s </div>`
	title := `<h3 class="list-group-item-heading"> %s </h3>`
	originalTitle := `<h5 class=""> %s </h5>`
	releaseDate := `<p class="list-group-item-text"> %s <i class="glyphicon glyphicon-calendar"></i></p>`
	voteAverage := `<p class=""> %g <i class="glyphicon glyphicon-star"></i> </p>`
	
	title = fmt.Sprintf(title, pMyMovieShort.Title)
	originalTitle = fmt.Sprintf(originalTitle, pMyMovieShort.OriginalTitle)
	releaseDate = fmt.Sprintf(releaseDate, pMyMovieShort.ReleaseDate[:10])
	voteAverage = fmt.Sprintf(voteAverage, util.Round(pMyMovieShort.VoteAverage, 2))
	
	caption = fmt.Sprintf(caption, title+originalTitle+releaseDate+voteAverage)
	    
	if strings.Compare(pMyMovieShort.PosterPath, "") == 0 {
		img = fmt.Sprintf(img, "http://personalcollectionmovies-alobaton.c9users.io/public/images/image_no_available.png")
	} else {
		img = fmt.Sprintf(img, "https://image.tmdb.org/t/p/w500"+pMyMovieShort.PosterPath)
	}
	 	
	thumbnail = fmt.Sprintf(thumbnail, img+caption)
	    
	item = fmt.Sprintf(item, "http://personalcollectionmovies-alobaton.c9users.io/movie?ID="+pMyMovieShort.ID, thumbnail)
	    
	return item
}