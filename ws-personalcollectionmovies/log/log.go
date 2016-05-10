package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	errorr *log.Logger
	warn   *log.Logger
	info   *log.Logger
	debug  *log.Logger

	NivelDebug = 1
	NivelInfo  = 2
	NivelWarn  = 3
	NivelError = 4
	NivelLog   = NivelDebug // Valor por defecto
)

func Warn(v ...interface{}) {
	if NivelWarn >= NivelLog {
		warn.Output(2, fmt.Sprintln(v...))
	}
}

func Warnf(format string, v ...interface{}) {
	if NivelWarn >= NivelLog {
		warn.Output(2, fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if NivelInfo >= NivelLog {
		info.Output(2, fmt.Sprintln(v...))
	}
}

func Infof(format string, v ...interface{}) {
	if NivelInfo >= NivelLog {
		info.Output(2, fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	if NivelError >= NivelLog {
		errorr.Output(2, fmt.Sprintln(v...))
	}
}

func Errorf(format string, v ...interface{}) {
	if NivelError >= NivelLog {
		errorr.Output(2, fmt.Sprintf(format, v...))
	}
}

func Debug(v ...interface{}) {
	if NivelDebug >= NivelLog {
		debug.Output(2, fmt.Sprintln(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if NivelDebug >= NivelLog {
		debug.Output(2, fmt.Sprintf(format, v...))
	}
}

/*
	Descripcion
		Funcion que se encarga de configurar la salida de logs en un fichero.
		- Inicializa las variables WARN, INFO, DEBUG.
	Parametros
		string :  pPath, identifica el nombre del fichero que contendra la salida de logs.
*/
func InitLog(pPath string, pLogLevel int) {
	NivelLog = pLogLevel
	file, err := os.OpenFile(pPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	multi := io.MultiWriter(file)

	info = log.New(multi, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(multi, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorr = log.New(multi, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	debug = log.New(multi, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

}
