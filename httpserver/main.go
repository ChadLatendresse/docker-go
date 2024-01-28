package main

// docker build -t project -f ./Docker/Dockerfile .
// docker run -it --rm -p 8080:8080 project
// or
// docker-compose up --build

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	filePath := filepath.Join("public", "home.html")
	http.ServeFile(w, r, filePath)
}

func main() {

	fs := http.FileServer(http.Dir("./public/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", home)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
