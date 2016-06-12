package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/database"
	"ws-personalcollectionmovies/model/domain"
	"ws-personalcollectionmovies/model/session"
	"fmt"
)
// Controlador.
type SessionController struct {
	beego.Controller
}

func (pController *SessionController) Login() {
    
	request := domain.LoginRequest{}
	err := pController.ParseForm(&request)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error("SessionController.go: "+error_.ERR_0020) 
		pController.ServeMessage(error_.KO, error_.ERR_0020)
    	return
	}
	fmt.Println("aqui estoy")
    // Verificamos que el usuario no haya sido registrado.
    
    rootVerification, err := domain.RootByUsername(database.OpenDataBase(), request.Username)
    // En caso de que el usuario ya exista.
	if err != nil || rootVerification.Username == "" {
		log.Error("SessionController.go: "+error_.ERR_0022+err.Error())
	    pController.ServeMessage(error_.KO, error_.ERR_0022)
	 	return 
	}
	
	// Encriptamos la contraseña para verificarla con la almacenada.
	encryptedPass := util.EncryptMD5(request.Password)
	// Verificamos que se trate de la misma contraseña.
	if rootVerification.Pass != encryptedPass {
	    log.Error("registercontroller.go: "+error_.ERR_0023)
		pController.ServeMessage(error_.KO, error_.ERR_0023)
	 	return
	}
	
	// Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	fmt.Println(sessionVal)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal == nil {
    	newSessionVal := session.UserSession{rootVerification.Username}
        pController.SetSession(session.USERSESSION, newSessionVal)
        fmt.Println(session.USERSESSION)
    }
    fmt.Println(pController.GetSession(session.USERSESSION))
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    pController.ServeMessage(error_.OK, "El usuario "+rootVerification.Username+" ha iniciado sesión.")
}

func (pController *SessionController) Logout() {
	sessionVal := pController.GetSession(session.USERSESSION)
	if sessionVal != nil {
		pController.DelSession(sessionVal)
		//sessionVal:=nil
	}
}

func (pController *SessionController) ServeMessage(pErrorCode, pErrorMessage string) {
    loginResponse := domain.LoginResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &loginResponse
    pController.ServeJSON()
}
