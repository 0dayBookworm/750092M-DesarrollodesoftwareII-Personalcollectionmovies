package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/session"
	"ws-personalcollectionmovies/model/database/connection"
	"ws-personalcollectionmovies/model/database/domain"
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/util"
)

type ProfileController struct {
	beego.Controller
}

// Metodo encargado de servir la pagina de perfil de usuario.
func (pController *ProfileController) Get() {
	pController.Layout = ACCOUNT
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = LOGIN_CONTROL
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	// Actualizamos los layouts.
    	pController.LayoutSections["SessionControl"] = LOGOUT_CONTROL
    	// Inertamos el titulo a la pagina.
		pController.Data["Title"] = sessionVal.(session.UserSession).Username
    	_, err := domain.AuditorByUsername(connection.GetConn(), sessionVal.(session.UserSession).Username)
		if err != nil {
    		// Actualizamos el layout de control de cuenta.
			pController.LayoutSections["AccountControl"] = COMMON_ACCOUNT_CONTROL
			// Actualizamos la sección de reportes.
			pController.LayoutSections["Reports"] = COMMON_NAV
    		pController.LayoutSections["AccountContent"] = PROFILE
    		
    		// Obtenemos el usuario de base de datos.
			// Verificamos que el usuario exista y obtenemos la información de base de datos.
			useraccount, err := domain.UseraccountByUsername(connection.GetConn(), sessionVal.(session.UserSession).Username)
			if err != nil {
				log.Error(error_.ERR_0031)
				// Se debería redireccionar a una pagina de error.
				pController.Redirect("/error", 302)
			}
			// Actualizamos los datos del formulario de registro.
			pController.Data["Username"] = sessionVal.(session.UserSession).Username
			pController.Data["Email"] = sessionVal.(session.UserSession).Email
			pController.Data["FirstName"] = useraccount.FirstName
			pController.Data["SecondName"] = useraccount.SecondName
			pController.Data["LastName"] = useraccount.LastName
			// pController.Data["BirthDate"] = util.ParseDate(useraccount.BirthDate)
			pController.Data["BirthDate"] = util.ParseDate(useraccount.BirthDate[:10])
			// Actualizamos el avatar.
			if (useraccount.Gender == "male") {
				pController.Data["Male"] = "checked"
				pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_male.jpg"
			} else {
				pController.Data["Female"] = "checked"
				pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_female.jpg"
			}
    		
    	} else {
    		pController.Redirect("/account/security", 302)
    	}
    	
		// Servimos la pagina.
		pController.TplName = ACCOUNT
	} else {
		pController.Redirect("/", 302)
	}
}

func (pController *ProfileController) Security() {
	pController.Layout = ACCOUNT
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = LOGIN_CONTROL
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	// Actualizamos los layouts.
    	pController.LayoutSections["SessionControl"] = LOGOUT_CONTROL
    		// Inertamos el titulo a la pagina.
		pController.Data["Title"] = sessionVal.(session.UserSession).Username
		pController.LayoutSections["AccountContent"] = SECURITY
		pController.Data["Username"] = sessionVal.(session.UserSession).Username
		pController.Data["Email"] = sessionVal.(session.UserSession).Email
			
		_, err := domain.AuditorByUsername(connection.GetConn(), sessionVal.(session.UserSession).Username)
		if err != nil {
	    	// Actualizamos el layout de control de cuenta.
			pController.LayoutSections["AccountControl"] = COMMON_ACCOUNT_CONTROL
			// Actualizamos la sección de reportes.
			pController.LayoutSections["Reports"] = COMMON_NAV
	    	// Actualizamos los datos del formulario de registro.
			
			// Actualizamos el avatar.
			if (sessionVal.(session.UserSession).Gender == "male") {
				pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_male.jpg"
			} else {
				pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_female.jpg"
			}
			pController.Data["EnableSecutiry"]=true
		} else {
			pController.LayoutSections["AccountControl"] = AUDIT_ACCOUNT_CONTROL
    		// Actualizamos el contenido del perfil.
    		pController.LayoutSections["AccountContent"] = SECURITY
    		// Actualizamos la sección de reportes.
    		// Actualizamos los datos del formulario de registro.
			pController.Data["Avatar"]="https://personalcollectionmovies-alobaton.c9users.io/public/images/avatar_audit.jpg"
			
			pController.Data["EnableSecutiry"]=false
		}
		// Servimos la pagina.
		pController.TplName = ACCOUNT
	} else {
		pController.Redirect("/", 302)
	}
}

// Metodo encargado de actualizar la información de perfil de usuario.
func (pController *ProfileController) Update() {
	request := wsinterface.UpdateProfileRequest{}
	err := pController.ParseForm(&request)
	
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	pController.StopRun()
	}
	
	// Validamos los campos.
	valid := validation.Validation{}
	b, err := valid.Valid(&request)
    if err != nil {
        log.Error(error_.ERR_0015+err.Error()) 
		pController.ServeMessage(error_.KO, error_.ERR_0015+err.Error())
    	pController.StopRun()
    }
    if !b {
        // Vaidation does not pass.
        log.Error("Validation does not pass.") 
        var errorMessage string
        for _, err := range valid.Errors {
        	errorMessage += `
        	`+err.Key+":"+err.Message+". "
        }
        log.Error(errorMessage) 
        pController.ServeMessage(error_.KO, errorMessage)
    	pController.StopRun()
    }
	
	// Procedemos a obtener el username a partir de la sesion.
	sessionVal := pController.GetSession(session.USERSESSION)
	// Verificamos primero que se encuentre en una sesion activa, es decir que no haya caducado la sesión.
	if (sessionVal == nil) {
		log.Error(error_.ERR_0032) 
		pController.ServeMessage(error_.KO, error_.ERR_0032)
    	pController.StopRun()
	}
	_username := sessionVal.(session.UserSession).Username
	
	log.Info(request)
	
	useraccount := domain.Useraccount{
	 	Username: _username, 
	 	FirstName: request.FirstName, 
	 	SecondName: request.SecondName, 
	 	LastName: request.LastName, 
	 	BirthDate: request.BirthDate, 
	 	Gender: request.Gender,
	 	Erased: false}
	 	
	err = useraccount.Update(connection.GetConn())
	if err != nil {
	 	log.Error(error_.ERR_0033+err.Error())
    	pController.ServeMessage(error_.KO, error_.ERR_0033)
	 	pController.StopRun() 
	}
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    pController.ServeMessage(error_.OK, "Tus datos personales se han actualizado con exito.")
}

// Metodo encargado de cambiar la contraseña de un usuario.
func (pController *ProfileController) ChangePassword() {
	request := wsinterface.ChangePasswordRequest{}
	err := pController.ParseForm(&request)
	
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	pController.StopRun()
	}
	// Procedemos a obtener el username a partir de la sesion.
	sessionVal := pController.GetSession(session.USERSESSION)
	// Verificamos primero que se encuentre en una sesion activa, es decir que no haya caducado la sesión.
	if (sessionVal == nil) {
		log.Error(error_.ERR_0032) 
		pController.ServeMessage(error_.KO, error_.ERR_0032)
    	pController.StopRun()
	}
	_username := sessionVal.(session.UserSession).Username
	
	root, err := domain.RootByUsername(connection.GetConn(), _username)
	if err != nil {
		//modifique el numero del error antes tenia el 31 preguntar si esta bien
		log.Error(error_.ERR_0013+err.Error())
	    pController.ServeMessage(error_.KO, error_.ERR_0013)
		pController.StopRun() 
	}
	
	encryptedPass := util.EncryptMD5(request.Password)
	if (encryptedPass == root.Pass){
		// Actualizamos la contraseña.
		root.Pass = util.EncryptMD5(request.NewPassword)
		log.Info(root)
		err = root.Update(connection.GetConn())
		if err != nil {
		 	log.Error(error_.ERR_0034+err.Error())
	    	pController.ServeMessage(error_.KO, error_.ERR_0034)
		 	pController.StopRun() 
		}
		// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
	    pController.ServeMessage(error_.OK, "Tu contraseña se ha actualizado con exito.")
	}else{
		// Se devuelve mensaje de contraseña incorrecta.
		pController.ServeMessage(error_.KO, error_.ERR_0023)
	}
}

// Metodo encargado de borrar de base de datos una cuenta de usuario.
func (pController *ProfileController) Delete() {

}

func (pController *ProfileController) ServeMessage(pErrorCode, pErrorMessage string) {
	profileResponse := wsinterface.ProfileResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &profileResponse
    pController.ServeJSON()
}

