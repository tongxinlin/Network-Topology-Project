package routers

import (
	//"fmt"
	"controllers"
	"net/http"
)

func Init() {
	fs := http.FileServer(http.Dir("src"))
	http.Handle("/", fs)
	
	http.HandleFunc("/home",controllers.HomeController)

} 