package util

import (
	"time"
	"crypto/sha256"
	"ws-personalcollectionmovies/log"
	)
// ParseDate [Método encargado de parsear una fecha de tipo string a tipo time.Time]
func ParseDate (pDate string) time.Time {
	const shortForm = "2006-01-02"
	t, err := time.Parse(shortForm, pDate)
	if err != nil {
		log.Error("ParseDate(): Error occurs while try parsing date.")
	}
	return t
}
// Encrypt [Encripta una cadena de caracteres y retorna su checksum]
func Encrypt (pData string) string {
	b := sha256.Sum256([]byte(pData))
	return string( b[:] )
}