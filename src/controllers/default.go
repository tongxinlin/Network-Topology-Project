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

// Global input and output directories
const(
	upload_dir = "./src/tmp/input/"
	output_dir = "./src/tmp/output/"
)

// Global variables
var uploadedFileName, outputFileName, dest, src, kPaths, node, neighborFileName string

// Handles the queries for shortest paths 
func ProcessQuery(rw http.ResponseWriter, req *http.Request){
	GetQueryData(rw,req)
	outputContent, _ := ioutil.ReadFile(outputFileName)
	fmt.Fprintln(rw,string(outputContent))
}

// Handles queries for neighbor searches
func GetNeighbors(rw http.ResponseWriter, req *http.Request){
	node = req.FormValue("reach")
    neighborFileName = dbhandler.NeighborsOf(node)
	outputContent, _ := ioutil.ReadFile(neighborFileName)
	fmt.Fprintln(rw,string(outputContent))
}

// If the input and output folders are not there create them. 
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

// Handles the file upload
func UploadFile(w http.ResponseWriter, r *http.Request){
	PrepareDirs()

	inputFile, header, _ := r.FormFile("upload-file")
	defer inputFile.Close()
	
	// create a new file
	uploadedFile, _ := os.OpenFile(upload_dir + header.Filename, os.O_CREATE|os.O_WRONLY, 0660)	
	defer uploadedFile.Close()
	
	// writes to the serverFile from the POST
	io.Copy(uploadedFile, inputFile)

	//save current inputfile name (in global)
	uploadedFileName = uploadedFile.Name()
    
    // Write the data and the shortest paths to db
    dbhandler.WriteToDB(uploadedFileName)

    ExecuteYensAlgorithm()
    dbhandler.WriteResultsToDB()
}

// Parses the values from query and calls QueryShortestPaths
func GetQueryData(w http.ResponseWriter, r *http.Request){
	// save all current query value (in global)
	dest = r.FormValue("dest")
	src = r.FormValue("src")
	kPaths = r.FormValue("kpaths")
    
    // Pass the resulting output in global variable
    outputFileName = dbhandler.QueryShortestPaths(src, dest, kPaths)
}


func ExecuteYensAlgorithm(){
	executablePath := "./src/executable/algorithm"
	// Process the uploaded file (uploadedFileName stored in global)
	cmd := exec.Command(executablePath, uploadedFileName)
    err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
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

