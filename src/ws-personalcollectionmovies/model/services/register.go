package services

import (
    )

// RegistrationRequest  [Estructura correspondiente a una petición de registro de cuenta de usuario.]
type RegistrationRequest struct {
	Username   string `form:"Username"`
	Email      string `form:"Email"`
	Password   string `form:"Password"`
	ConfirmPassword string `form:"ConfirmPassword"`
	FirstName  string `form:"FirstName"`
	SecondName string `form:"SecondName"`
	LastName   string `form:"LastName"`
	BirthDate  string  `form:"BirthDate"`
	Gender     string  `form:"Gender"`
	TermsAndConditions string `form:"TermsAndConditions"`
}

// RegistrationRequestToString [Metodo encargado de convertir a string los datos de la estructura RegistrationRequest correspondiente a un formulario.]
func  (pRegistrationRequest *RegistrationRequest) ToString() string{
    user := "{Username:"+pRegistrationRequest .Username+", Email:"+pRegistrationRequest .Email+", Password:"+pRegistrationRequest .Password+", FirstName:"+pRegistrationRequest .FirstName+", SecondName:"+pRegistrationRequest .SecondName+", LastName:"+pRegistrationRequest .LastName+", BirthDate:"+pRegistrationRequest .BirthDate+", Gender:"+pRegistrationRequest .Gender+", TermsAndConditions:"+pRegistrationRequest .TermsAndConditions+"}"
	return user
}

// Registrationresponse {Estructura correspondiente a una respuesta para una peticion de registro]
type RegistrationResponse struct {
	Status string
	Message string
}



