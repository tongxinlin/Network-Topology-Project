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
    "dbhandler"
)

const(
	upload_dir = "./src/tmp/input/"
	output_dir = "./src/tmp/output/"
)


var uploadedFileName, outputFileName, dest, src, kPaths string

func ProcessQuery(rw http.ResponseWriter, req *http.Request){
	GetQueryData(rw,req)
	ExecuteAlgorithm()
	outputContent, _ := ioutil.ReadFile(outputFileName)
	fmt.Fprintln(rw,string(outputContent))
}

func PrepareDirs(){
	//create ./src/tmp/input & ./src/tmp/output
	_, err1 := os.Stat(upload_dir)
	if err1 != nil {
		os.MkdirAll(upload_dir, 0711)
	}

	_, err2 := os.Stat(output_dir)
	if err2 != nil {
		os.MkdirAll(output_dir, 0711)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request){
	PrepareDirs()

	// "upload-file" is from the POST method of the form on the web page
	inputFile, header, _ := r.FormFile("upload-file")
	defer inputFile.Close()
	
	// tells OS to create a file with appicable permissions
	uploadedFile, _ := os.OpenFile(upload_dir + header.Filename, os.O_CREATE|os.O_WRONLY, 0660)	
	defer uploadedFile.Close()
	
	// writes to the serverFile from the POST
	io.Copy(uploadedFile, inputFile)
	//save current inputfile name (in global)
	uploadedFileName = uploadedFile.Name()
}


func GetQueryData(w http.ResponseWriter, r *http.Request){
	//save all current query value (in global)
	dest = r.FormValue("dest")
	src = r.FormValue("src")
	kPaths = r.FormValue("kpaths")
}


func ExecuteAlgorithm(){
	executablePath := "./src/executable/algorithm"
	//command line arguments that will be passed to the algorithm
	argv := []string{uploadedFileName, dest, src, kPaths}
	cmd := exec.Command(executablePath, argv...)

	//get the output file name from stdout
	output, _ := cmd.Output()
	outputFileName = string(output)
}


func RenderHomepage(rw http.ResponseWriter, req *http.Request) {
	
	// grab index.html from views 
	homepage, err := template.ParseFiles("src/views/html/index.html")
	
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// send it to the browser
	homepage.Execute(rw, homepage)
}


func RenderQueryPage(rw http.ResponseWriter, req *http.Request) {
	
	// grab query.html from views 
	query, err := template.ParseFiles("src/views/html/query.html")
	
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// send it to the browser
	query.Execute(rw, query)
}

