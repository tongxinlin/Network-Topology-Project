package main

import (
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
)

http.HandleFunc("/upload", upload)

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(32 << 20) // Has to be called to parse the form
        uploadedFile, handler, err := r.FormFile("upload-file")
        if err != nil {
            fmt.Println(err)
            return
        }
        defer uplodedFile.Close()
        fmt.Fprintf(w, "%v", handler.Header)
        // Open local file with flags and permission (TODO: right permission?)
        localFile, err := os.OpenFile("/tmp/uploadfile", os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer localFile.Close()
        // If gets here no errors possible anymore
        io.Copy(localFile, uploadedFile)
}