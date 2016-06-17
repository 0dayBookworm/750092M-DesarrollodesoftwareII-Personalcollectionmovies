package wsinterface

import (
	"github.com/astaxie/beego/validation"
)
type UpdateProfileRequest struct {
	Email string `form:"Email" valid:"Required`
	FirstName string `form:"FirstName" valid:"Required`
	SecondName string `form:"SecondName" valid:"Required`
	LastName string `form:"LastName" valid:"Required`
	BirthDate string `form:"BirthDate" valid:"Required`
	Gender string `form:"Gender" valid:"Required`
}

type ChangePasswordRequest struct {
	Password string `form:"Password"`
	NewPassword string `form:"NewPassword"`
	ConfirmPassword string `form:"CofirmPassword"`
}

type DeleteAccountRequest struct {
	Password string
}

// RegistrationRequest  [Estructura correspondiente a una petición de registro de cuenta de usuario.]
type ProfileResponse struct {
	Status     string
	Message    string
}

// Validadores
func (u *UpdateProfileRequest) Valid(v *validation.Validation) {
}

func (u *ChangePasswordRequest) Valid(v *validation.Validation) {
}
func (u *DeleteAccountRequest) Valid(v *validation.Validation) {
}