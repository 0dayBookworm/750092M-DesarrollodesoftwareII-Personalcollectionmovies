package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/database"
	"ws-personalcollectionmovies/model/domain"
	"ws-personalcollectionmovies/model/session"
)

// Constantes.
const REGISTER = "register.tpl"
const REGISTER_TITLE = "Registrate"

// Controlador.
type RegisterController struct {
	beego.Controller
}

func (pController *RegisterController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = REGISTER_TITLE
    
   	pController.VerifySession()
	
	// Servimos la pagina.
	pController.TplName = REGISTER
}
	
// CreateUseraccount [Metodo encargado de procesar una petición de registro de una nueva cuenta de usuario.]
func (pController *RegisterController) CreateUseraccount() {
	// Extraemos los datos del usuario a partir del formulario.
	// Esto se realiza mediante un parseo, se debe crear una estructura refente al modulo en l paquete 'form'.
	// Además es importante conocer que el parseo se realiza a partir de un JSon, es decir el formulario es traducido a un JSon.
	request := domain.RegistrationRequest{}
	err := pController.ParseForm(&request)
	
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error("registercontroller.go: "+error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	return
	}
	
	log.Info("registercontroller.go: Record attempt [" + request.ToString() + "]")
	
	
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
	// Verificamos que el usuario no haya sido registrado.
    rootVerification, err := domain.RootByUsername(database.OpenDataBase(), request.Username)
	if err == nil && rootVerification.Username != "" {
		log.Error("registercontroller.go: "+error_.ERR_0014)
		pController.ServeMessage(error_.KO, error_.ERR_0014)
	 	return 
	}
	
	// Si el usuario no existe entonces:
    // Dado el diseño debemos primero insertar el usuario en la tabla ROOT.
    root := domain.Root{
     	Username: strings.ToLower(request.Username),
     	Pass: util.EncryptMD5(request.Password)}
     	
	err = root.Insert(database.OpenDataBase())
	if err != nil {
		log.Error("registercontroller.go: "+error_.ERR_0010+err.Error())
		pController.ServeMessage(error_.KO, error_.ERR_0013)
	 	return 
	}
    
	// Creamos el objeto de dominio que sera almacenado en la tabla USEACCOUNT.
	useraccount := domain.Useraccount{
	 	Username: strings.ToLower(request.Username), 
	 	FirstName: request.FirstName, 
	 	SecondName: request.SecondName, 
	 	LastName: request.LastName, 
	 	BirthDate: request.BirthDate, 
	 	Gender: request.Gender,
	 	Email: strings.ToLower(request.Email), 
	 	Erased: false}
	
	err = useraccount.Insert(database.OpenDataBase())
	if err != nil {
	 	log.Error("registercontroller.go: "+error_.ERR_0011+err.Error())
    	pController.ServeMessage(error_.KO, error_.ERR_0013)
	 	return 
	}
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    pController.ServeMessage(error_.OK, "Te has registrado con exito, te hemos enviado un mail de confirmación.")
}

func (pController *RegisterController) VerifySession() {
    pController.Layout = REGISTER
    pController.LayoutSections = make(map[string]string)
    pController.LayoutSections["SessionControl"] = "logincontrol.tpl"
    // Se realiza el inicio de sesión almacenando el username en memoria.
 	sessionVal := pController.GetSession(session.USERSESSION)
 	// Si es nulo significa que no esta en una sesión activa e iniciamos la sesión.
    if sessionVal != nil {
    	 pController.LayoutSections["SessionControl"] = "logoutcontrol.tpl"
    	 pController.Data["Username"] = sessionVal.(session.UserSession).Username
    }
}

func (pController *RegisterController) ServeMessage(pErrorCode, pErrorMessage string) {
	registrationResponse := domain.RegistrationResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &registrationResponse
    pController.ServeJSON()
}

