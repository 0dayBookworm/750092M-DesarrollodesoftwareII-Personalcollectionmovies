package myhttplib

import (
	"github.com/astaxie/beego/httplib"
	"bytes"
 	"fmt"
	
// 	"net/http"
// 	"io/ioutil"
    )

// Estructura que nos permita definir los parametros de una busqueda de tipo Get principalmente.
type Param struct {
    Key string
    Value string
}

func get(pUrl string, pParams []Param) *httplib.BeegoHTTPRequest {
    buffer := bytes.NewBufferString(pUrl)
    buffer.WriteString("?")
    for i := range pParams {
        param := pParams[i]
        buffer.WriteString(param.Key)
        buffer.WriteString("=")
        buffer.WriteString(param.Value)
        buffer.WriteString("&")
    }
    getUrl := buffer.String()
    fmt.Println(getUrl)
    req := httplib.Get(getUrl)
    return req
}

func GetStringFrom(pUrl string, pParams []Param) (string, error) {
    req := get(pUrl, pParams)
    return req.String()
}

func GetJSONFrom (pUrl string, pParams []Param, pStore interface{}) error {
    req := get(pUrl, pParams)
    return req.ToJSON(pStore)
}
