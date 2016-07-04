package connection

import (
    "database/sql"
    "github.com/knq/dburl"
    _ "github.com/lib/pq"
    "ws-personalcollectionmovies/log"
    // "ws-personalcollectionmovies/model/database/domain"
    )
    
var (
    Db = &sql.DB{}
    )

func Init (pDriver, pUser, pPass, pHost, pName string) {
    db, err := dburl.Open(pDriver+"://"+pUser+":"+pPass+"@"+pHost+"/"+pName)
	if err != nil {
		log.Error(err)
		panic(err)
	} else {
	    Db = db
	}
}

func GetConn() (*sql.DB) {
    return Db
}