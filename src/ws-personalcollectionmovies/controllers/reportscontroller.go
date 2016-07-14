package controllers

import (
 	"fmt"
// 	"strconv"
 	"strings"
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
// 	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/session"
// 	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/model/database/connection"
	"ws-personalcollectionmovies/model/database/domain"
// 	"ws-personalcollectionmovies/model/tmdb"
)

type ReportsController struct {
	beego.Controller
}


func (pController *ReportsController) ViewListInLast30Days() {
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
		pController.LayoutSections["AccountControl"] = AUDIT_ACCOUNT_CONTROL
		// Actualizamos la sección de reportes.
		pController.LayoutSections["Reports"] = AUDIT_NAV
		// Actualizamos el cuerpo.
    	pController.LayoutSections["AccountContent"]=COLLECTION
    	
    	_username := sessionVal.(session.UserSession).Username
    	email := sessionVal.(session.UserSession).Email
		// Actualizamos los datos del formulario de registro.
		pController.Data["Username"] = _username
		pController.Data["Email"] = email
		// Actualizamos el avatar.
		pController.Data["Avatar"]="http://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_audit.jpg"
			
		// Inertamos el titulo a la pagina.
		pController.Data["Title"] = "PELÍCULAS VISTAS EN LOS ÚLTIMOS 30 DÍAS"
	
    	// Traemos de BD la infformación necesaría para el cuerpo.
    	
        viewList, err := domain.ViewListInLast30Days(connection.GetConn())
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

func (pController *ReportsController) ViewListOrderByCinema() {
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
		pController.LayoutSections["AccountControl"] = AUDIT_ACCOUNT_CONTROL
		// Actualizamos la sección de reportes.
		pController.LayoutSections["Reports"] = AUDIT_NAV
		// Actualizamos el cuerpo.
    	pController.LayoutSections["AccountContent"]=LIST
    	
    	_username := sessionVal.(session.UserSession).Username
    	email := sessionVal.(session.UserSession).Email
		// Actualizamos los datos del formulario de registro.
		pController.Data["Username"] = _username
		pController.Data["Email"] = email
		// Actualizamos el avatar.
		pController.Data["Avatar"]="http://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_audit.jpg"
			
		// Inertamos el titulo a la pagina.
		pController.Data["Title"] = "PELÍCULAS VISTAS AGRUPADAS POR CINEMA"
	
    	// Traemos de BD la infformación necesaría para el cuerpo.
    	
        viewList, err := domain.ViewListOrderByCinema(connection.GetConn())
        if err != nil {
			log.Error("No funciona la vista")
			pController.Redirect("/error", 302)
		}
		// Variables
		group := `<li class="list-group-item"> 
					<a class="dropdown-toggle" data-toggle="collapse" data-target="#item%d">
						<i class="glyphicon glyphicon-chevron-down"></i>
						<i class="glyphicon glyphicon-map-marker"></i>
    					<strong>  %s </strong>
    					<small> (%d resultados) </small>
					</a>
					<ul id="item%d" class="collapse nav">
						<hr>
    					%s
					</ul> 
				</li>`
		firstItem := (*viewList)[0]
		tempPlace := firstItem.Place
		itemHtml := `<li class="list-item"> 
						<a class="thumbnail row" href="http://personalcollectionmovies-alobaton.c9users.io/movie?ID=%s" > 
							<img class="pull-left" src="https://image.tmdb.org/t/p/w500%s" alt="Generic placeholder thumbnail" />
							<div class"caption> 
								<h3> %s </h3>
								<h6> %s </h6>
								<p> %s <i class="glyphicon glyphicon-calendar"></i> </p>
							</div>
						</a> 
					</li>`
		tempContent := fmt.Sprintf(itemHtml, firstItem.ID, firstItem.PosterPath, firstItem.Title, firstItem.OriginalTitle, firstItem.ReleaseDate[:10])
		movieShortsContent := ""
		totalCinema := 1 
		totalItemsByCinema := 1
		for  i := 1; i < len(*viewList); i++ {
			item := (*viewList)[i]
			if strings.Compare(item.Place, tempPlace) != 0 {
				// Guardamos el resultado.
				movieShortsContent += fmt.Sprintf(group, totalCinema, tempPlace, totalItemsByCinema, totalCinema, tempContent) 
				totalCinema = totalCinema+1
				// Actualizamos las variables temporales.
				totalItemsByCinema = 1
				tempPlace = item.Place
				tempContent = fmt.Sprintf(itemHtml, item.ID, item.PosterPath, item.Title, item.OriginalTitle, item.ReleaseDate[:10])
			} else {
				totalItemsByCinema = totalItemsByCinema+1
				tempContent += fmt.Sprintf(itemHtml, item.ID, item.PosterPath, item.Title, item.OriginalTitle, item.ReleaseDate[:10])	
			}
		}
		movieShortsContent += fmt.Sprintf(group, totalCinema, tempPlace, totalItemsByCinema, totalCinema, tempContent) 
		// Ponemos los resultados en la pagina.
		pController.Data["ListItems"]=movieShortsContent
    	pController.Data["Total"]=totalCinema
    	// Servimos la pagina.
		pController.TplName = ACCOUNT
    } else{
    	pController.Redirect("/", 302)
    }
}