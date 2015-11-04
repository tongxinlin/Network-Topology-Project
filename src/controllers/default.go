package controllers

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
    "os"
	"io"
	"io/ioutil"
	"os/exec"
)

const(
	upload_dir = "./src/tmp/input/"
	output_dir = "./src/tmp/output/"
)


////////////////////////////////////////////////////////////////////////////////////////////////
//just want to make it looks easier, ignore all error checking temporarily


func ProcessRequest(rw http.ResponseWriter, req *http.Request){
	PrepareDirs()
	outputFileName := ProcessedFile(UploadedFile(rw,req))
	
	fmt.Println(outputFileName)
	
	outputContent, _ := ioutil.ReadFile(outputFileName)
	fmt.Fprintln(rw,string(outputContent))
}

func PrepareDirs(){
	_, err1 := os.Stat(upload_dir)
	if err1 != nil {
		os.MkdirAll(upload_dir, 0711)
	}

	_, err2 := os.Stat(output_dir)
	if err2 != nil {
		os.MkdirAll(output_dir, 0711)
	}
}

func UploadedFile(w http.ResponseWriter, r *http.Request) (string, string) {
	// "upload-file" is from the POST method of the form on the web page
	inputFile, header, _ := r.FormFile("upload-file")
	kValue := r.FormValue("k")
	defer inputFile.Close()
	
	// tells OS to create a file with appicable permissions
	uploadedFile, _ := os.OpenFile(upload_dir + header.Filename, os.O_CREATE|os.O_WRONLY, 0660)
	
	defer uploadedFile.Close()
	
	// writes to the serverFile from the POST
	io.Copy(uploadedFile, inputFile)
	uploadedFileName := uploadedFile.Name()
	return uploadedFileName, kValue
}

func ProcessedFile(uploadedFile string, kValue string)string{
	executablePath := "./src/executable/test.exe"
	argv := []string{uploadedFile, kValue}
	cmd := exec.Command(executablePath, argv...)
	output, _ := cmd.Output()
	outputFileName := string(output)
	
	return outputFileName
}
////////////////////////////////////////////////////////////////////////////////////////////////

func HomeController(rw http.ResponseWriter, req *http.Request) {

	renderHomepage(rw, req)
}

func renderHomepage(rw http.ResponseWriter, req *http.Request) {
	
	// grab the homepage from views 
	homepage, err := template.ParseFiles("src/views/html/homepage.html")
	
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// send it to the browser
	homepage.Execute(rw, homepage)
}


