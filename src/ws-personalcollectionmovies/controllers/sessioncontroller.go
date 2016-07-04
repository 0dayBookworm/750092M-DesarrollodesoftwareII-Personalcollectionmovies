package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/database/connection"
	"ws-personalcollectionmovies/model/database/domain"
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/model/session"
)
// Controlador.
type SessionController struct {
	beego.Controller
}

func (pController *SessionController) Login() {
    
	request := wsinterface.LoginRequest{}
	err := pController.ParseForm(&request)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error("SessionController.go: "+error_.ERR_0020) 
		pController.ServeMessage(error_.KO, error_.ERR_0020)
    	pController.StopRun()
	}
	// Validamos los campos.
	valid := validation.Validation{}
	b, err := valid.Valid(&request)
    if err != nil {
        log.Error("registercontroller.go: "+error_.ERR_0015+err.Error()) 
		pController.ServeMessage(error_.KO, error_.ERR_0015+err.Error())
    	pController.StopRun()
    }
    if !b {
        // Vaidation does not pass.
        log.Error("registercontroller.go: Validation does not pass.") 
        var errorMessage string
        for _, err := range valid.Errors {
        	errorMessage += `
        	`+err.Key+":"+err.Message+". "
        }
        log.Error(errorMessage) 
        pController.ServeMessage(error_.KO, errorMessage)
    	pController.StopRun()
    }
	
    rootVerification, err := domain.RootByUsername(connection.GetConn(), request.Username)
	if err != nil || rootVerification.Username == "" {
		log.Error(error_.ERR_0022+err.Error())
	    pController.ServeMessage(error_.KO, error_.ERR_0022)
	 	pController.StopRun() 
	}
	
	// Encriptamos la contraseña para verificarla con la almacenada.
	encryptedPass := util.EncryptMD5(request.Password)
	// Verificamos que se trate de la misma contraseña.
	if rootVerification.Pass != encryptedPass {
	    log.Error(error_.ERR_0023+rootVerification.Pass+"|"+encryptedPass+"|"+request.Password)
		pController.ServeMessage(error_.KO, error_.ERR_0023)
	 	pController.StopRun()
	}
	// Creamos la variable de sessión.
	newSessionVal := session.UserSession{}
	newSessionVal.Username = rootVerification.Username
	
	// Obtenemos otros datos necesarios para el manejo de sesión.
	useraccount, err := domain.UseraccountByUsername(connection.GetConn(), request.Username)
	if err != nil {
		// Verificamos que sea entonces un auditor.
		// Obtenemos los datos del auditor.
		_, err := domain.AuditorByUsername(connection.GetConn(), request.Username)
		if err != nil {
			log.Error(error_.ERR_0023)
			pController.ServeMessage(error_.KO, error_.ERR_0023)
			pController.StopRun()
		} 
	} else {
		newSessionVal.Email = useraccount.Email
		newSessionVal.Gender = useraccount.Gender
	}
	
	// Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal == nil {
        pController.SetSession(session.USERSESSION, newSessionVal)
    }
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    pController.ServeMessage(error_.OK, "El usuario "+rootVerification.Username+" ha iniciado sesión.")
}

func (pController *SessionController) Logout() {
	if pController.GetSession(session.USERSESSION) != nil {
		pController.DelSession(session.USERSESSION)
	}
	pController.Redirect("/", 302)
}

func (pController *SessionController) ServeMessage(pErrorCode, pErrorMessage string) {
    loginResponse := wsinterface.LoginResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &loginResponse
    pController.ServeJSON()
}
