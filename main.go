package main

import(
	"fmt"
	"net/http"
	"strings"
	"time"
)

var	url = make(map[string]string)
func main() {

	http.HandleFunc("/", handleForm)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/short", handleRedirect)

	http.ListenAndServe(":8080", nil)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	
}

func handleShorten(w http.ResponseWriter, r *http.Request) {

}

func handleRedirect(w http.ResponseWriter, r *http.Request) {

}