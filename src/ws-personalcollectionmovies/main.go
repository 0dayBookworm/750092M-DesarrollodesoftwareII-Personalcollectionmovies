package main

import (
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/model/database/connection"
	"ws-personalcollectionmovies/model/tmdb"
	"ws-personalcollectionmovies/model/google"
	_ "ws-personalcollectionmovies/routers"
	"github.com/astaxie/beego"
)

func main() {
	// A este nivel debemos cargar el contexto de la aplicación
	// Inicializamos el log.
	logPath := beego.AppConfig.String("Logs::LogPath")
	logName  := beego.AppConfig.String("Logs::LogName")
	logLevel, err:= beego.AppConfig.Int("Logs::LogLevel")
	if (err != nil) {
		panic(err)
	}
	log.InitLog(logPath + logName, logLevel)
	// Informacion de los logs cargada correctamente.
	log.Info("Log information load correctly.")
	
	// Aviso de que se ha configurado el contexto correctamente.
	log.Info("The system is set up correctly.")
	
	// Iniciamos las conexiones con la base de datos.
	Driver := beego.AppConfig.String("DataBase::Driver")
	User := beego.AppConfig.String("DataBase::User")
	Pass := beego.AppConfig.String("DataBase::Pass")
	Host := beego.AppConfig.String("DataBase::Host")
	Name := beego.AppConfig.String("DataBase::Name")
	connection.Init(Driver, User, Pass, Host, Name)
	
	// Inicializamos la TMDb Api.
	ApiKey := beego.AppConfig.String("TMDb::ApiKey")
	Language := beego.AppConfig.String("TMDb::Language")
	tmdb.Init(ApiKey, Language)
	
	// Inicializamos la API de google
	GoogleKey := beego.AppConfig.String("Google::ApiKey")
	GoogleUri := beego.AppConfig.String("Google::PlacesUri")
	MapsUri := beego.AppConfig.String("Google::MapsUri")
	google.Init(GoogleKey, GoogleUri, MapsUri)
	
	// Necesario para levantar el servidor de arvhivos en la ruta public.
	beego.SetStaticPath("/public", "/home/ubuntu/workspace/src/ws-personalcollectionmovies/resources/public")
	// Beego run.
	log.Info("Listen and serve at " + beego.AppConfig.String("httpport"))
	beego.Run()
}
