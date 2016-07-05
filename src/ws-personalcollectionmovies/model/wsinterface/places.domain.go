package wsinterface

import (
	"github.com/astaxie/beego/validation"
)

type PlacesRequest struct {
	Place string `form:"Place" valid:"Required`
}

// Estructura respuesta.
type PlacesResponse struct {
    Status string
    Message string
}
func (u *PlacesRequest) Valid(v *validation.Validation) {
}
