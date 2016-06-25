package wsinterface

import (
	"github.com/astaxie/beego/validation"
)

type GetMovieRequest struct {
   ID int `form:"ID" valid:"Required`
}


type SearchMovieRequest struct {
	Title string `form:"Title" valid:"Required`
	Year string 
}

type SearchMovieResponse struct {
	Status     string
	Message    string
}

// Validadores
func (u *SearchMovieRequest) Valid(v *validation.Validation) {
}
