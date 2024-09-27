package main

import (
	"net/http"
	"routers"
	"log" 
)

func main() {

	//inititalize app 
	routers.InititalizeApp()
	//set the static hosting server, local host for testing right now
	host := ":8080"
	
	//error log
	if err := http.ListenAndServe(host,nil); err != nil {
		log.Fatal("Server Error", err)
	}
}