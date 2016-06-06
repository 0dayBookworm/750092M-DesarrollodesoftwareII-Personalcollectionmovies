package form

import (
    "github.com/astaxie/beego/validation"
    )

// Useraccount 
type Useraccount struct {
	Username   string `form:"Username" valid:"Required"`
	Email      string `form:"Email" valid:"Required;Email;MaxSize(100)"`
	Password   string `form:"Password" valid:"Required"`
	ConfirmPassword string `form:"ConfirmPassword" valid:"Required"`
	FirstName  string `form:"FirstName" valid:"Required"`
	SecondName string `form:"SecondName" valid:"Required"`
	LastName   string `form:"LastName" valid:"Required"`
	BirthDate  string  `form:"BirthDate" valid:"Required"`
	Gender     string  `form:"Gender" valid:"Required"`
	TermsAndConditions string `form:"TermsAndConditions" valid:"Required"`
}

// Funcion para validaciones personalizadas de la estructura 'useraccount' del paquete 'model/form'.
// Se ejecuta siempre despues de validar los tags de la estructura.
func (u *Useraccount) Valid(v *validation.Validation) {
}

func UseraccountToString(pUseraccount Useraccount) string{
    user := "Username:"+pUseraccount.Username+", Email:"+pUseraccount.Email+", Password:"+pUseraccount.Password+", FirstName:"+pUseraccount.FirstName+", SecondName:"+pUseraccount.SecondName+", LastName:"+pUseraccount.LastName+", BirthDate:"+pUseraccount.BirthDate+", Gender:"+pUseraccount.Gender+", TermsAndConditions:"+pUseraccount.TermsAndConditions
	return user
}

