package models

import (
    "fmt"
 	//"bufio"
 	//"os"
 	//"strings"
	//"sort"
    "log"
    "gopkg.in/mgo.v2"
    //"gopkg.in/mgo.v2/bson"
)


const DB_HOST = "localhost:27017"


func InititalizeDB() {
	session, err := GetDB()
	defer session.Close()
	if err == nil {
        fmt.Println("Database connection verified")
		log.Println("Database connection verified")
	} else {
		log.Fatalln("models.init failed", err)
		fmt.Println("models.init failed", err)
    }

}

func GetDB() (session *mgo.Session, err error) {
	session, err = mgo.Dial(DB_HOST)
	return
}

