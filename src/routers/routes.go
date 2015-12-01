package routers

import (
	"controllers"
	"net/http"
)

func InititalizeApp() {
  	// set up the html
	fs := http.FileServer(http.Dir("src/views/html"))
	http.Handle("/", fs)
	http.Handle("/css/", http.FileServer(http.Dir("src/views")))
    http.Handle("/js/", http.FileServer(http.Dir("src/views")))
	
	// routes the pages of the apps to the browser
	http.HandleFunc("/home",controllers.RenderHomepage)
	http.HandleFunc("/upload",controllers.UploadFile)
	http.HandleFunc("/process",controllers.ProcessQuery)	
	http.HandleFunc("/query",controllers.RenderQueryPage)
	http.HandleFunc("/getReachability",controllers.GetNeighbors)    
} 