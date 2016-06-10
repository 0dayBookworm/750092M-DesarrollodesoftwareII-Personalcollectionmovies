package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/util"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/database"
	"ws-personalcollectionmovies/model/domain"
	"ws-personalcollectionmovies/model/services"
)

type RegisterController struct {
	beego.Controller
}

func (pController *RegisterController) Get() {
	// Para insertar datos en los templates usamos: pController.Data["Identificador en el template"] = "Valor que queremos que tome"
	// Inertamos el titulo a la pagina.
	pController.Data["Title"] = "Registrate"
	
	
	
	
	// Servimos la pagina.
	pController.TplName = "register.html"
}
// CreateUseraccount [Metodo encargado de procesar una petición de registro de una nueva cuenta de usuario.]
func (pController *RegisterController) CreateUseraccount() {
	// Extraemos los datos del usuario a partir del formulario.
	// Esto se realiza mediante un parseo, se debe crear una estructura refente al modulo en l paquete 'form'.
	// Además es importante conocer que el parseo se realiza a partir de un JSon, es decir el formulario es traducido a un JSon.
	request := services.RegistrationRequest{}
	err := pController.ParseForm(&request)
	
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error("CreateUseraccount: "+error_.ERR_0012) 
		RegistrationResponse := services.RegistrationResponse{error_.KO, error_.ERR_0012}
        pController.Data["json"] = &RegistrationResponse
    	pController.ServeJSON()
    	return
	}
	
	log.Info("CreateUseraccount: Record attempt [" + request.ToString() + "]")
	
    // Dado el diseño debemos primero insertar el usuario en la tabla ROOT.
    // Además esto nos ahorrara procesamiento de no poderse llevar acabo.
    domainRoot := domain.Root{
     	Username: request.Username,
     	// Pass: util.Encrypt(request.Password)}
     	Pass: request.Password}
     	
	err = domainRoot.Insert(database.OpenDataBase())
	if err != nil {
		log.Error(error_.ERR_0010+err.Error())
		RegistrationResponse := services.RegistrationResponse{error_.KO, error_.ERR_0013}
        pController.Data["json"] = &RegistrationResponse
    	pController.ServeJSON()
	 	return 
	}
    
	// Creamos el objeto de dominio que sera almacenado en base de datos.
	useraccount := domain.Useraccount{
	 	Username: request.Username, 
	 	FirstName: request.FirstName, 
	 	SecondName: request.SecondName, 
	 	LastName: request.LastName, 
	 	BirthDate: util.ParseDate(request.BirthDate), 
	 	Gender: request.Gender,
	 	Email: request.Email, 
	 	Erased: false}
	
	err = useraccount.Insert(database.OpenDataBase())
	if err != nil {
	 	log.Error(error_.ERR_0011+err.Error())
	 	RegistrationResponse := services.RegistrationResponse{error_.KO, error_.ERR_0013}
        pController.Data["json"] = &RegistrationResponse
    	pController.ServeJSON()
	 	return 
	}
	// Si el proceso se llevo a cabo respondemos con el mensaje de exito asociado
    RegistrationResponse := services.RegistrationResponse{error_.OK, "Te has registrado con exito, te hemos enviado un mail de confirmación."}
    pController.Data["json"] = &RegistrationResponse
    pController.ServeJSON()
}

