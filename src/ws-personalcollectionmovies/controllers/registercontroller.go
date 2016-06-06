package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
	"ws-personalcollectionmovies/log"
	//"ws-personalcollectionmovies/model/domain"
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

func (pController *RegisterController) CreateUseraccount() {
	// Extraemos los datos del usuario a partir del formulario.
	// Esto se realiza mediante un parseo, se debe crear una estructura refente al modulo en l paquete 'form'.
	useraccount := form.Useraccount{}
	err := pController.ParseForm(&useraccount)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Info("CreateUseraccount(): Error occurs while try parsing useraccount from form.")
	}
	log.Info("CreateUseraccount(): Se obtuvieron los datos del usuario correctamente [" + form.UseraccountToString(useraccount) + "]")
	// Creamos el objeto de dominio que sera almacenado en base de datos.
	
	
	
	// Para redireccionar a otra pagina. En este caso a home.
	pController.Ctx.Redirect(302, "/")
}