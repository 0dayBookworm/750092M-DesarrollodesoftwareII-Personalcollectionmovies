package main

import (
	"750092M-DesarrollodesoftwareII-Personalcollectionmovies/ws-personalcollectionmovies/handler"
	"750092M-DesarrollodesoftwareII-Personalcollectionmovies/ws-personalcollectionmovies/log"
	"750092M-DesarrollodesoftwareII-Personalcollectionmovies/ws-personalcollectionmovies/util"
	"flag"

	"gopkg.in/gcfg.v1"
)

// Variable utilizada para guardar el nombre del archivo de configuracion
// del servicio que se le pasa como parametro al iniciar la instancia
var fileConfig = flag.String("fileconfig", "", "The configuration file")

func main() {
	// Se utiliza para poder obtener el parametro del nombre del archivo
	flag.Parse()
	var cfg util.Config
	// Se lee el archivo de configuración del servicio y se parsea en la variable cfgConfig
	err := gcfg.ReadFileInto(&cfg, *fileConfig)
	if err != nil {
		panic(err)
	}
	// Inicializamos la informacion de la aplicacion.
	util.InitServiceInformation(cfg.Host.Ip, cfg.Host.Port, cfg.Host.ServicePath)
	// Inicializamos el log.
	log.InitLog(cfg.Logs.LogPath+cfg.Logs.LogName, cfg.Logs.LogLevel)
	// Inicializamos la configuración de la ruta publica.
	util.InitPublic(cfg.Public.PublicName, cfg.Public.PublicPath)
	// Inicializamos la configuración para la carga de templates.
	util.InitTemplates(cfg.Templates.TemplatesPath, cfg.Templates.Home)

	// Arrancamos el servicio.
	handler.HandleSetup()
}
