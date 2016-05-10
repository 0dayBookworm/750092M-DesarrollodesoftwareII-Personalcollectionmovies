package handler

import (
	"750092M-DesarrollodesoftwareII-Personalcollectionmovies/ws-personalcollectionmovies/log"
	"750092M-DesarrollodesoftwareII-Personalcollectionmovies/ws-personalcollectionmovies/util"
	"html/template"
	"net/http"
)

var tpl *template.Template

func HandleSetup() {
	// Inicio del servicio.
	log.Info("Service start.")

	// Definimos los diferentes endpoints.
	// Endpoint principal de la aplicacion. Ruta en la cual se desplegara la aplicacion
	http.HandleFunc(util.ServicePath, Home)

	// Endpoint publico encargado de proporcionar recursos publicos que necesitemos (Como lo son hojas de estilo).
	http.Handle("/public/", http.StripPrefix(util.PublicName, http.FileServer(http.Dir(util.PublicPath))))

	// Cargamos los archivos html.
	tpl = template.Must(template.ParseGlob(util.TemplatesPath + "*.html"))

	// Levantamos el servidor.
	http.ListenAndServe(util.Ip+":"+util.Port, nil)
}

func Home(pRes http.ResponseWriter, pReq *http.Request) {
	//serveTemplate(pRes, pReq, "home.html")
	tpl.ExecuteTemplate(pRes, util.Home, nil)
}
