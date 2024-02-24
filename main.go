package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urls = make(map[string]string)

func main() {

	http.HandleFunc("/", handleForm)
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/short/", handleRedirect)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w,
		`<!DOCTYPE html>
		<html>
		<head>
			<title> Welcome to URL SHORTENER </title>
		</head>
		<body>
			<form method = "post" action = "/shorten" >
				<input type = "url" name = "url" placeholder = "Enter a URL" required>
				<input type = "submit" value = "shorten">
			</form>
		</body>
		</html>
		`)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid operation", http.StatusMethodNotAllowed)
		return
	}

	originalUrl := r.FormValue("url")

	// Check if original URL is empty
	if originalUrl == "" {
		http.Error(w, "no url present", http.StatusBadRequest)
		return
	}

	shortKey := generateShortKey()
	urls[shortKey] = originalUrl

	shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w,
		`<!DOCTYPE html>
		<html>
		<head>
			<title> Welcome to URL SHORTENER </title>
		</head>
		<body>
			 <h3> URL SHortener </h3>
			 <p> Original URL := `, originalUrl, `</p>
			 <p> Shortened URL := <a href= "`, shortenedURL, `"> `, shortenedURL, `</a></p>
		</body>
		</html>
		`)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := strings.TrimPrefix(r.URL.Path, "/short/")

	if shortKey == "" {
		http.Error(w, "shortened key is missing", http.StatusBadRequest)
		return
	}

	originalURL, found := urls[shortKey]
	if !found {
		http.Error(w, "short key is not present", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)

}

func generateShortKey() string {
	const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUV123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)

	for i := range shortKey {
		shortKey[i] = char[rand.Intn(len(char))]
	}

	return string(shortKey)
}

