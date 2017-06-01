package main

import (
	"log"
	"net/http"
)

func fn(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/status")
}
func fn2(w http.ResponseWriter, r *http.Request) {
	  http.ServeFile(w, r, "files/status.asc")

}
func main() {
	http.HandleFunc("/server", fn)
	http.HandleFunc("/server.asc", fn2)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

