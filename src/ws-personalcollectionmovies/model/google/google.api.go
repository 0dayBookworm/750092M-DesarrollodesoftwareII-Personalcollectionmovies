package google

import (
    "strings"
	"ws-personalcollectionmovies/model/myhttplib"
	"ws-personalcollectionmovies/log"
)


var (
	ApiKey = ""
	ApiUri =""
	)

func Init(pApiKey, pApiUri string) {
	ApiKey = pApiKey
	ApiUri = pApiUri
}
// Result for GetPlaces
type AutocompleteResult struct {
    Predtions []Prediction `json:"predictions"`
    Status string `json:"status"`
}
// Un resultado de busqueda.
type Prediction struct {
    Description string `json:"description"`
    ID string `json:"id"`
    MatchedSubstrings []MatchedSubstring `json:"matched_substrings"`
    PlaceId string `json:"place_id"`
    Reference string `json:"reference"`
    Terms []Term `json:"terms"`
    Types []string `json:"types"`
}
type MatchedSubstring struct {
    Length int `json:"length"`
    Offset int `json:"offset"`
}
type Term struct {
    Offset int `json:"offset"`
    Value string `json:"value"`
}

func GetPlaces(pPlace string) (AutocompleteResult){
    // Armamos los parametros.
    params := []myhttplib.Param{
        myhttplib.Param{"input", strings.Replace(pPlace, " ", "%20", -1)}, 
        myhttplib.Param{"types", "establishment"}, 
        myhttplib.Param{"key", ApiKey}}
    // Creamos el objeto donde almacenaremos la respuesta.
    res := AutocompleteResult{}
    err := myhttplib.GetJSONFrom(ApiUri, params, &res)
    if err != nil {
        log.Error(err.Error())
    }
    return res
}