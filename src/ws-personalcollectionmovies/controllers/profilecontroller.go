package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/session"
	"ws-personalcollectionmovies/model/database"
	"ws-personalcollectionmovies/model/domain"
	"ws-personalcollectionmovies/util"
)


const PROFILE = "profile.tpl"
const PROFILE_TITLE ="Perfil"

type ProfileController struct {
	beego.Controller
}

// Metodo encargado de servir la pagina de perfil de usuario.
func (pController *ProfileController) Get() {
		// Inertamos el titulo a la pagina.
	pController.Data["Title"] = PROFILE_TITLE
  
	pController.Layout = PROFILE
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = "logincontrol.tpl"
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = "logoutcontrol.tpl"
    	 
    	 _username := sessionVal.(session.UserSession).Username
    
		// Obtenemos el usuario de base de datos.
		// Verificamos que el usuario exista y obtenemos la información de base de datos.
		Useraccount, err := domain.UseraccountByUsername(database.OpenDataBase(), _username)
		if err != nil {
			log.Error("profilecontroller.go: "+error_.ERR_0031)
			// Se debería redireccionar a una pagina de error.
			pController.Redirect("/error", 302)
		}
		// Actualizamos los datos del formulario de registro.
		pController.Data["Username"] = Useraccount.Username
		pController.Data["Email"] = Useraccount.Email
		pController.Data["FirstName"] = Useraccount.FirstName
		pController.Data["SecondName"] = Useraccount.SecondName
		pController.Data["LastName"] = Useraccount.LastName
		pController.Data["BirthDate"] = Useraccount.BirthDate
		if (Useraccount.Gender == "male") {
			pController.Data["Male"] = "checked"
		} else {
			pController.Data["Female"] = "checked"
		}
		// Servimos la pagina.
		pController.TplName = PROFILE
	} else {
		pController.Redirect("/", 302)
	}
}

// Metodo encargado de actualizar la información de perfil de usuario.
func (pController *ProfileController) Update() {
	request := domain.UpdateProfileRequest{}
	err := pController.ParseForm(&request)
	
	if err != nil {
		log.Error("registercontroller.go: "+error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	return
	}
	
	// Validamos los campos.
	valid := validation.Validation{}
	b, err := valid.Valid(&request)
    if err != nil {
        log.Error("registercontroller.go: "+error_.ERR_0015+err.Error()) 
		pController.ServeMessage(error_.KO, error_.ERR_0015+err.Error())
    	return
    }
    if !b {
        // Vaidation does not pass.
        log.Error("registercontroller.go: Validation does not pass.") 
        var errorMessage string
        for _, err := range valid.Errors {
        	errorMessage += `
        	`+err.Key+":"+err.Message+". "
        }
        log.Error("registercontroller.go: "+errorMessage) 
        pController.ServeMessage(error_.KO, errorMessage)
    	return
    }
	
	// Procedemos a obtener el username a partir de la sesion.
	sessionVal := pController.GetSession(session.USERSESSION)
	// Verificamos primero que se encuentre en una sesion activa, es decir que no haya caducado la sesión.
	if (sessionVal == nil) {
		log.Error("registercontroller.go: "+error_.ERR_0032) 
		pController.ServeMessage(error_.KO, error_.ERR_0032)
    	return
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
	 	Email: request.Email, 
	 	Erased: false}
	 	
	err = useraccount.Update(database.OpenDataBase())
	if err != nil {
	 	log.Error("profilecontroller.go: "+error_.ERR_0033+err.Error())
    	pController.ServeMessage(error_.KO, error_.ERR_0033)
	 	return 
	}
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    pController.ServeMessage(error_.OK, "Tus datos personales se han actualizado con exito.")
}

// Metodo encargado de cambiar la contraseña de un usuario.
func (pController *ProfileController) ChangePassword() {
	request := domain.ChangePasswordRequest{}
	err := pController.ParseForm(&request)
	
	if err != nil {
		log.Error("registercontroller.go: "+error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	return
	}
	// Procedemos a obtener el username a partir de la sesion.
	sessionVal := pController.GetSession(session.USERSESSION)
	// Verificamos primero que se encuentre en una sesion activa, es decir que no haya caducado la sesión.
	if (sessionVal == nil) {
		log.Error("registercontroller.go: "+error_.ERR_0032) 
		pController.ServeMessage(error_.KO, error_.ERR_0032)
    	return
	}
	_username := sessionVal.(session.UserSession).Username
	
	root, err := domain.RootByUsername(database.OpenDataBase(), _username)
	if err != nil {
		//modifique el numero del error antes tenia el 31 preguntar si esta bien
		log.Error("profilecontroller.go: "+error_.ERR_0013+err.Error())
	    pController.ServeMessage(error_.KO, error_.ERR_0013)
		return 
	}
	
	encryptedPass := util.EncryptMD5(request.Password)
	if (encryptedPass == root.Pass){
		// Actualizamos la contraseña.
		root.Pass = util.EncryptMD5(request.NewPassword)
		log.Info(root)
		err = root.Update(database.OpenDataBase())
		if err != nil {
		 	log.Error("profilecontroller.go: "+error_.ERR_0034+err.Error())
	    	pController.ServeMessage(error_.KO, error_.ERR_0034)
		 	return 
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
	profileResponse := domain.ProfileResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &profileResponse
    pController.ServeJSON()
}

