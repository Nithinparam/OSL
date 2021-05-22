package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func GetFile(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	m, _ := filepath.Glob("upload/" + name)
	if m == nil {
		fmt.Fprintln(w, "No file")
	} else {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprintln(w, `<a href="/download/`+name+`">Click Here</a>`)

	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, _ := os.Create("upload/" + handler.Filename)
	io.Copy(f, file)
	f.Close()
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
func Download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Header().Set("Content-Type", "multipart/form-data")
	fmt.Fprintln(w, "upload/"+name)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/get", GetFile)
	router.HandleFunc("/uploads", uploadFile)
	router.HandleFunc("/download/{name}", Download)
	http.Handle("/upload/", http.StripPrefix("/upload/", http.FileServer(http.Dir("./upload/"))))
	http.Handle("/", router)
	http.ListenAndServe(":9090", nil)

}
