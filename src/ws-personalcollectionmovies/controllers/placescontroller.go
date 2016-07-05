package controllers

import (
	"github.com/astaxie/beego"
	"ws-personalcollectionmovies/log"
	"ws-personalcollectionmovies/error_"
	"ws-personalcollectionmovies/model/wsinterface"
	"ws-personalcollectionmovies/model/google"
)

type PlacesController struct {
	beego.Controller
}
func (pController *PlacesController) GetPlaces() {
    request := wsinterface.PlacesRequest{}
	err := pController.ParseForm(&request)
	// Si ha ocurrido un error al parsear.
	if err != nil {
		log.Error(error_.ERR_0012) 
		pController.ServeMessage(error_.KO, error_.ERR_0012)
    	pController.StopRun()
	}
	
    res := google.GetPlaces(request.Place)
    pController.Data["json"] =  res
    pController.ServeJSON()
}
func (pController *PlacesController) ServeMessage(pErrorCode, pErrorMessage string) {
    placesResponse := wsinterface.PlacesResponse{pErrorCode, pErrorMessage}
    pController.Data["json"] = &placesResponse
    pController.ServeJSON()
}
