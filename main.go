package main

import (
	"log"
	"net/http"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/status")
}
func status_ascHandler(w http.ResponseWriter, r *http.Request) {
	  http.ServeFile(w, r, "static/status.asc")

}
func main() {
	http.HandleFunc("/vlc/status", statusHandler)
	http.HandleFunc("/vlc/status.asc", status_ascHandler)
	log.Fatal(http.ListenAndServe(":80", nil))

}

