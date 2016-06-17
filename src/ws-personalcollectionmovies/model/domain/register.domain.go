package domain


import (
	"github.com/astaxie/beego/validation"
)
// RegistrationRequest  [Estructura correspondiente a una petición de registro de cuenta de usuario.]
type RegistrationRequest struct {
	Username   string `form:"Username" valid:"Required"`
	Email      string `form:"Email" valid:"Required"`
	Password   string `form:"Password" valid:"Required"`
	ConfirmPassword string `form:"ConfirmPassword" valid:"Required"`
	FirstName  string `form:"FirstName" valid:"Required"`
	SecondName string `form:"SecondName"`
	LastName   string `form:"LastName" valid:"Required"`
	BirthDate  string  `form:"BirthDate" valid:"Required"`
	Gender     string  `form:"Gender" valid:"Required"`
	TermsAndConditions string `form:"TermsAndConditions" valid:"Required"`
}

// Registrationresponse [Estructura correspondiente a una respuesta para una peticion de registro]
type RegistrationResponse struct {
	Status string
	Message string
}

// RegistrationRequestToString [Metodo encargado de convertir a string los datos de la estructura RegistrationRequest correspondiente a un formulario.]
func  (pRegistrationRequest *RegistrationRequest) ToString() string{
    user := "{Username:"+pRegistrationRequest .Username+", Email:"+pRegistrationRequest .Email+", Password:"+pRegistrationRequest .Password+", FirstName:"+pRegistrationRequest .FirstName+", SecondName:"+pRegistrationRequest .SecondName+", LastName:"+pRegistrationRequest .LastName+", BirthDate:"+pRegistrationRequest .BirthDate+", Gender:"+pRegistrationRequest .Gender+", TermsAndConditions:"+pRegistrationRequest .TermsAndConditions+"}"
	return user
}

func (u *RegistrationRequest) Valid(v *validation.Validation) {
}