package main

import (
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/model/tmdb"
	_ "ws-personalcollectionmovies/routers"
	"github.com/astaxie/beego"
)

// Variable utilizada para guardar el nombre del archivo de configuracion
// del servicio que se le pasa como parametro al iniciar la instancia
//var fileConfig = flag.String("fileconfig", "", "The configuration file")

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
	
	// Inicializamos la TMDb Api.
	ApiKey := beego.AppConfig.String("TMDb::ApiKey")
	Language := beego.AppConfig.String("TMDb::Language")
	tmdb.Init(ApiKey, Language)
	
	// Necesario para levantar el servidor de arvhivos en la ruta public.
	beego.SetStaticPath("/public", "/home/ubuntu/workspace/src/ws-personalcollectionmovies/resources/public")
	// Beego run.
	log.Info("Listen and serve at " + beego.AppConfig.String("httpport"))
	beego.Run()
}
