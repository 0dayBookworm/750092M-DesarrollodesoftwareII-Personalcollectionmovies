package controllers

import (
	"time"
	"strconv"
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/util"
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
		// Actualizamos el cuerpo.
    	pController.LayoutSections["AccountContent"]=COLLECTION
    	
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
    	pController.Data["Total"]=len(*watchList)
    	// Servimos la pagina.
		pController.TplName = ACCOUNT
    } else{
    	pController.Redirect("/", 302)
    }
}

func (pController *CollectionController) Viewlist() {
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
		// Actualizamos el cuerpo.
    	pController.LayoutSections["AccountContent"]=COLLECTION
    	
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
		pController.Data["Title"] = VIEWLIST_TITLE
		// Traemos de BD la infformación necesaría para el cuerpo.
    	
        viewList, err := domain.ViewListByUsername(connection.GetConn(), _username)
        if err != nil {
			log.Error("No funciona la vista")
			pController.Redirect("/error", 302)
		}
		movieShortsContent := ""
		for _, item := range *viewList {
			// log.Info(item)
			// Parseamos a un MovieShort personalizado con el fin de reutilizar código.
			myMovieShort := wsinterface.MyMovieShort{
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
    	pController.Data["Total"]=len(*viewList)
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
	// REALIZAMOS MAPEO.
	_, err := domain.MovieMappingByID(connection.GetConn(), request.ID)
 	if err != nil {
		pController.DoMovieMapping(request.ID)
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



func (pController *CollectionController) ViewListAdd() {
	// Obtenemos la pagina solicitada.
	request := wsinterface.AddRequest{}
	pController.ParseForm(&request)
	// Validaciones
	// Primero se verifica la sesión, ya que de ahi se obtendran los datos del usuario.
	sessionVal := pController.GetSession(session.USERSESSION)
	if(sessionVal == nil) {
		pController.ServeMessage(error_.KO, error_.ERR_0050)
    	pController.StopRun()
	}
	pController.ValidateAuditor(sessionVal.(session.UserSession).Username)
	
	// Verificamos si se encuentra ya en BD la pelicula.
	movieMapping, err := domain.MovieMappingByID(connection.GetConn(), request.ID)
 	if err != nil {
		pController.DoMovieMapping(request.ID)
 	}
 	// Validamos que la pelicula ya haya sido estrenada.
 	t := time.Now()
 	if t.Before(util.ParseDate(movieMapping.ReleaseDate[:10])) {
 		pController.ServeMessage(error_.KO, error_.ERR_0055)
    	pController.StopRun()
 	}
 	
 	// Si se encuentra asociamos el usuario a esa pelicula.
 	userViewList := domain.UserViewMovie {
 		Username: sessionVal.(session.UserSession).Username,
 		ID: request.ID,
 		Place: request.Place,
 		Dateandtime: "NOW()",
 		Erased: false}
 	err = userViewList.Insert(connection.GetConn())
 	if err != nil {
 		log.Error(err.Error())
 		// Si ocurre un error quiere decir que la pelicula ya se encuentra añadida a la lista de peliculas por ver.
	 	pController.ServeMessage(error_.KO, error_.ERR_0053)
	    pController.StopRun()
 	}
	pController.ServeMessage(error_.OK, "Añadida a lista de Películas vistas.")
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

func (pController *CollectionController) DoMovieMapping(pID string) {
	
 	// Si no se encuentra debemos generar el registo en nuestra BD.
 	// Realizamos el mapping de la pelicula a nuestra BD.
	// Convertimos el ID a int
	idInt, _ := strconv.Atoi(pID)
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




