package routers

import (
	"controllers"
	"net/http"
)

func Init() {
  	//testing out Go's file server stuff 
	fs := http.FileServer(http.Dir("src/views/html"))
	http.Handle("/", fs)
	http.Handle("/css/", http.FileServer(http.Dir("src/views")))
    http.Handle("/js/", http.FileServer(http.Dir("src/views")))
	
	// routes the homepage to the browser
	http.HandleFunc("/home",controllers.HomeController)
	http.HandleFunc("/upload",controllers.Upload)	

} 