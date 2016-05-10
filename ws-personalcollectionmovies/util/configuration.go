package util

// Config estructura utilizada para guardar los datos del archivo
// config.ini para la configuración del servicio
type Config struct {
	Host
	Logs      Logs
	Public    Public
	Templates Templates
}

// Host estructura utilizada para guardar los datos del archivo
// config.ini para configuración de la aplicación.
type Host struct {
	Ip          string
	Port        string
	ServicePath string
}

// Logs estructura utilizada para guardar los datos del archivo
// config.ini para la configuración de los logs
type Logs struct {
	LogPath  string
	LogName  string
	LogLevel int // LevelDebug = 1,	LevelInfo  = 2,	LevelWarn  = 3,	LevelError = 4 , LevelDebug = 5 o superior(desactiva el log)
}

// Public estructura utilizada para guardar los datos del archivo config.ini
// para la configuración de la ruta public/
type Public struct {
	PublicName string
	PublicPath string
}

// Templates estructura utilizada para guardar los datos del archivo config.ini
// para la configuración de la carga de html.
type Templates struct {
	TemplatesPath string
	Home          string
}
