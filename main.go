package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
    "net/url"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/status")
}
func status_ascHandler(w http.ResponseWriter, r *http.Request) {
	  http.ServeFile(w, r, "static/status.asc")

}
func home(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST"{
        body, _ := ioutil.ReadAll(r.Body)
        values, _ := url.ParseQuery(string(body))
        for v , k := range values{
            fmt.Println(v , strings.Join(k, ""))
        }
    }
    if r.Method == "GET"{
        r.ParseForm() 
        for k, v := range r.Form {
            fmt.Println(v , strings.Join(k, ""))
        }

    }
}

func main() {
	http.HandleFunc("/showoff", home) 
	http.HandleFunc("/vlc/status", statusHandler)
	http.HandleFunc("/vlc/status.asc", status_ascHandler)
	log.Fatal(http.ListenAndServe(":80", nil))

}

