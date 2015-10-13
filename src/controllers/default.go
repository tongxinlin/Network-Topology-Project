package controllers

import (
	"html/template"
	"log"
	//"models"
	//"strings"
	"net/http"
)

func HomeController(rw http.ResponseWriter, req *http.Request) {

	
	// grab the homepage from views 
	homepage, err := template.ParseFiles("src/views/homepage.html")
	
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	//send it to the browser
	homepage.Execute(rw, homepage)
}