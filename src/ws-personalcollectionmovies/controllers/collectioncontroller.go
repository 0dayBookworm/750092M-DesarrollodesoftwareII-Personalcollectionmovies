package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/session"
	
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/model/database/connection"
	"ws-personalcollectionmovies/model/database/domain"
	"ws-personalcollectionmovies/model/tmdb"
)

type CollectionController struct {
	beego.Controller
}
func (pController *CollectionController) Get() {
	pController.Layout = ACCOUNT
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = LOGIN_CONTROL
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	// Inertamos el titulo a la pagina.
		pController.Data["Title"] = COLLECTION_TITLE
    	// Actualizamos los layouts.
    	pController.LayoutSections["SessionControl"]=LOGOUT_CONTROL
    	// Actualizamos el layout de control de cuenta.
		pController.LayoutSections["AccountControl"] = COMMON_ACCOUNT_CONTROL
		// Actualizamos la sección de reportes.
		pController.LayoutSections["Reports"] = COMMON_NAV
		
    	pController.LayoutSections["AccountContent"] = COLLECTIONS
    	 
    	_username := sessionVal.(session.UserSession).Username
    	email := sessionVal.(session.UserSession).Email
    	gender := sessionVal.(session.UserSession).Gender
    	
    	// Actualizamos los datos del formulario de registro.
		pController.Data["Username"] = _username
		pController.Data["Email"] = email
		// Actualizamos el avatar.
		if (gender == "male") {
			pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_male.jpg"
		} else {
			pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_female.jpg"
		}
    	
    	// Servimos la pagina.
		pController.TplName = ACCOUNT
    } else{
    	pController.Redirect("/", 302)
    }
}

func (pController *CollectionController) Watchlist() {
	pController.Layout = ACCOUNT
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = LOGIN_CONTROL
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	// Actualizamos información general del perfil.
    	pController.LayoutSections["SessionControl"]=LOGOUT_CONTROL
    	// Actualizamos el layout de control de cuenta.
		pController.LayoutSections["AccountControl"] = COMMON_ACCOUNT_CONTROL
		// Actualizamos la sección de reportes.
		pController.LayoutSections["Reports"] = COMMON_NAV
		
    	pController.LayoutSections["AccountContent"] = COLLECTIONS
    	
    	_username := sessionVal.(session.UserSession).Username
    	email := sessionVal.(session.UserSession).Email
    	gender := sessionVal.(session.UserSession).Gender
		// Actualizamos los datos del formulario de registro.
		pController.Data["Username"] = _username
		pController.Data["Email"] = email
		// Actualizamos el avatar.
		if (gender == "male") {
			pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_male.jpg"
		} else {
			pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_female.jpg"
		}
		// Inertamos el titulo a la pagina.
		pController.Data["Title"] = WATCHLIST_TITLE
		// Actualizamos el cuerpo.
    	pController.LayoutSections["AccountContent"]=COLLECTION
    	// Traemos de BD la infformación necesaría para el cuerpo.
    	
        watchList, err := domain.WatchListByUsername(connection.GetConn(), _username)
        if err != nil {
			log.Error("No funciona la vista")
			pController.Redirect("/error", 302)
		}
		movieShortsContent := ""
		for _, item := range *watchList {
			// log.Info(item)
			// Parseamos a un MovieShort personalizado con el fin de reutilizar código.
			myMovieShort := wsinterface.MyMovieShort{
				Username : item.Username,
				ID: item.ID,
				Title: item.Title,
				OriginalTitle: item.OriginalTitle,
				ReleaseDate: item.ReleaseDate,
				VoteAverage: item.VoteAverage,
				PosterPath: item.PosterPath}
			
		    movieShortsContent += ConcatenateMovieShortResult(myMovieShort)
		}
		// Ponemos los resultados en la pagina.
		pController.Data["MovieShorts"]=movieShortsContent
    	
    	// Servimos la pagina.
		pController.TplName = ACCOUNT
    } else{
    	pController.Redirect("/", 302)
    }
}

func (pController *CollectionController) WatchListAdd() {
	// Obtenemos la pagina solicitada.
	request := wsinterface.AddRequest{}
	pController.ParseForm(&request)
	// Primero se verifica la sesión, ya que de ahi se obtendran los datos del usuario.
	sessionVal := pController.GetSession(session.USERSESSION)
	if(sessionVal == nil) {
		pController.ServeMessage(error_.KO, error_.ERR_0050)
    	pController.StopRun()
	}
	pController.ValidateAuditor(sessionVal.(session.UserSession).Username)
	// Verificamos si se encuentra ya en BD la pelicula.
 	_, err := domain.MovieMappingByID(connection.GetConn(), request.ID)
 	if err != nil {
 		// Si no se encuentra debemos generar el registo en nuestra BD.
 		// Realizamos el mapping de la pelicula a nuestra BD.
		// Convertimos el ID a int
		idInt, _ := strconv.Atoi(request.ID)
		// Realizamos la busqueda en la API.
	 	res, err := tmdb.GetMovieInfo(idInt)
	 	// Si no hay resultados para la pelicula informamos el error.
	 	if err != nil {
	 		log.Error(error_.ERR_0040) 
	 		pController.ServeMessage(error_.KO, error_.ERR_0040)
	     	pController.StopRun()
	 	}
	 	// Creamos el objeto de dominio.
	 	movieMapping := &domain.MovieMapping{
	 		ID: strconv.Itoa(res.ID),
	 		Title: res.Title,
	 		OriginalTitle: res.OriginalTitle,
	 		ReleaseDate: res.ReleaseDate,
	 		VoteAverage: float64(res.VoteAverage),
	 		PosterPath: res.PosterPath }
	 	
 		err = movieMapping.Insert(connection.GetConn())
 		// Si no se pudo realizar el mapeo informamos el error.
 		if err != nil {
 			log.Error(error_.ERR_0052+err.Error()) 
	 		pController.ServeMessage(error_.KO, error_.ERR_0051)
	     	pController.StopRun()
 		}
 		// Añadimos a la lista de peliculas por ver del usuario.
 	} 
 	// Si se encuentra asociamos el usuario a esa pelicula.
 	userWatchList := domain.UserNextviewsMovie{
 		Username: sessionVal.(session.UserSession).Username,
 		ID: request.ID,
 		Dateandtime: "NOW()",
 		Erased: false}
 	err = userWatchList.Insert(connection.GetConn())
 	if err != nil {
 		// Si ocurre un error quiere decir que la pelicula ya se encuentra añadida a la lista de peliculas por ver.
	 	pController.ServeMessage(error_.KO, error_.ERR_0053)
	    pController.StopRun()
 	}
 	pController.ServeMessage(error_.OK, "Añadida a lista de Películas por ver.")
}

func (pController *CollectionController) ServeMessage(pErrorCode, pErrorMessage string) {
	collectionResponse := wsinterface.CollectionResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &collectionResponse
    pController.ServeJSON()
}

func (pController *CollectionController) ValidateAuditor(pUsername string) {
	_, err := domain.AuditorByUsername(connection.GetConn(), pUsername)
	if err == nil {
		pController.ServeMessage(error_.KO, error_.ERR_0054)
	    pController.StopRun()
	}
}


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
	voteAverage = fmt.Sprintf(voteAverage, pMyMovieShort.VoteAverage)
	
	caption = fmt.Sprintf(caption, title+originalTitle+releaseDate+voteAverage)
	    
	if strings.Compare(pMyMovieShort.PosterPath, "") == 0 {
		img = fmt.Sprintf(img, "https://personalcollectionmovies-alobaton.c9users.io/public/images/image_no_available.png")
	} else {
		img = fmt.Sprintf(img, "https://image.tmdb.org/t/p/w500"+pMyMovieShort.PosterPath)
	}
	 	
	thumbnail = fmt.Sprintf(thumbnail, img+caption)
	    
	item = fmt.Sprintf(item, "https://personalcollectionmovies-alobaton.c9users.io/movie?ID="+pMyMovieShort.ID, thumbnail)
	    
	return item
}

