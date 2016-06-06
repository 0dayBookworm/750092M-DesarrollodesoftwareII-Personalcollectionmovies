package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"time"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/model/database"
	"ws-personalcollectionmovies/model/domain"
	"ws-personalcollectionmovies/model/form"
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
	useraccount := form.Useraccount{}
	err := pController.ParseForm(&useraccount)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error("CreateUseraccount(): Error occurs while try parsing useraccount from form.")
	}
	log.Info("CreateUseraccount(): Record attempt [" + form.UseraccountToString(useraccount) + "]")
	
	// Realizamos las validaciones.
	valid := validation.Validation{}
	// customErr es un error personalizado despues de validar los tags de la estructura. Remitirse a 'model/form/useraccount.go'
    customErr, err := valid.Valid(&useraccount)
    if err != nil {
        // Handle error.
        log.Error("CreateUseraccount(): Error occurs while try validating useraccount from form.")
        return 
    }
	if !customErr {
        // Validation does not pass.
        log.Error("CreateUseraccount(): Validation does not pass.")
        // Obtenemos todos los errores capturados.
        for _, err := range valid.Errors {
        	log.Error("CreateUseraccount(): ["+err.Key+"] "+err.Message)
        }
        return 
    }
    // Si se pasan las validaciones entonces:
	// Creamos el objeto de dominio que sera almacenado en base de datos.
	domainUseraccount := domain.Useraccount{}
	domainUseraccount.Username = useraccount.Username
	domainUseraccount.FirstName = useraccount.FirstName
	domainUseraccount.SecondName = useraccount.SecondName
	domainUseraccount.LastName = useraccount.LastName
	domainUseraccount.BirthDate, _ = time.Parse("dd/mm/yyyy", useraccount.BirthDate)
	domainUseraccount.Gender = useraccount.Gender
	domainUseraccount.Country = ""
	domainUseraccount.Email = useraccount.Email
	domainUseraccount.Erased = false
	
	err = domainUseraccount.Insert(database.OpenDataBase())
	if err != nil {
		log.Error("CreateUseraccount(): Useraccount could not be created."+err.Error())
		return 
	}
	// Para redireccionar a otra pagina. En este caso a home.
	pController.Ctx.Redirect(302, "/")
}