package models

//import (
//	 "fmt"
// 	"labix.org/v2/mgo"
// 	"bufio"
// 	"os"
// 	"strings"
// 	"labix.org/v2/mgo/bson"
//	"sort"
//)
//
//
//const DB_HOST = "localhost"
//
//
//func InititalizeDB() {
//	session, err := GetDB()
//	defer session.Close()
//	if err == nil {
//		log.Println("Database connection verified")
//	} else {
//		log.Fatalln("models.init failed", err)
//	}
//
//}
//
//func GetDB() (session *mgo.Session, err error) {
//	session, err = mgo.Dial(DB_HOST)
//	return
//}