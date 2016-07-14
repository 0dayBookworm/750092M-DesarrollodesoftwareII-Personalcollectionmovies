package util

import (
	"time"
	"crypto/sha256"
	"crypto/md5"
	"io"
	"bytes"
	"fmt"
	"ws-personalcollectionmovies/log"
	)
	
var(
	MONTHS = [...]string{"ENE", "FEB", "MAR", "ABR", "MAY", "JUN", "JUL", "AGO", "SEP", "OCT", "NOV", "DIC"}
	)
	
// ParseDate [Método encargado de parsear una fecha de tipo string a tipo time.Time]
func ParseDate (pDate string) time.Time {
	const shortForm = "2006-01-02"
	t, err := time.Parse(shortForm, pDate)
	if err != nil {
		log.Error("ParseDate(): Error occurs while try parsing date:"+pDate)
	}
	return t
}
// EncryptSHA256 [Encripta una cadena de caracteres y retorna su checksum]
func EncryptSHA256 (pData string) string {
	b := sha256.Sum256([]byte(pData))
	return string( b[:] )
}
// EncryptMD5 [Encripta una cadena de caracteres y retorna su checksum]
func EncryptMD5 (pData string) string {
	md5Data := md5.New()
    io.WriteString(md5Data, pData)
    buffer := bytes.NewBuffer(nil)
    fmt.Fprintf(buffer, "%x", md5Data.Sum(nil))
    return buffer.String()
}

func GetMonth(pMonth int) string {
	// fmt.Println(pMonth)
	return MONTHS[pMonth-1]
}

func Round(v float64, decimals int) float64 {
     var pow float64 = 1
     for i:=0; i<decimals; i++ {
         pow *= 10
     }
     return float64(int((v * pow) + 0.5)) / pow
}