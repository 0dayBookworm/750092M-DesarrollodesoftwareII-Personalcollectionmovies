package domain

import (
	"github.com/astaxie/beego/validation"
)
// Estructura peticion.
type LoginRequest struct {
    Username string `form:"Username" valid:"Required"`
    Password string `form:"Password" valid:"Required"`
}
// Estructura respuesta.
type LoginResponse struct {
    Status string
    Message string
}
func (u *LoginRequest) Valid(v *validation.Validation) {
}