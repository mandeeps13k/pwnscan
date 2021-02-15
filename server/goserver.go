package server

import (
	"fmt"
	"log"
	"net/http"
)

func Goserver() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/serv", Serve)
	http.HandleFunc("/form", FormHandler)
	fmt.Printf("Hello")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Sending request
func Serve(w http.ResponseWriter, r *http.Request) {
	log.Printf("whatever")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "SUCCESS")
	name := r.FormValue("name")

	fmt.Fprintf(w, "Name = %s\n", name)
}
