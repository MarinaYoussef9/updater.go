package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/status")
}
func status_ascHandler(w http.ResponseWriter, r *http.Request) {
	  http.ServeFile(w, r, "static/status.asc")

}
func home(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() 
    for k, v := range r.Form {
        fmt.Print(k , ": ")
        fmt.Println(strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello VLC!")
}

func main() {
	http.HandleFunc("/showoff", home) 
	http.HandleFunc("/vlc/status", statusHandler)
	http.HandleFunc("/vlc/status.asc", status_ascHandler)
	log.Fatal(http.ListenAndServe(":80", nil))

}

