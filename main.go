package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Text   string
	Banner string
	Art    string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		showError(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		showError(w, http.StatusMethodNotAllowed)
		return
	}

	render(w, PageData{Banner: "standard"})
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		showError(w, http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		showError(w, http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || !isValidBanner(banner) {
		showError(w, http.StatusBadRequest)
		return
	}

	art, err := makeArt(text, banner)
	if errors.Is(err, errBadChar) {
		showError(w, http.StatusBadRequest)
		return
	}
	if errors.Is(err, os.ErrNotExist) {
		showError(w, http.StatusNotFound)
		return
	}
	if err != nil {
		showError(w, http.StatusInternalServerError)
		return
	}

	render(w, PageData{Text: text, Banner: banner, Art: art})
}

func render(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if errors.Is(err, os.ErrNotExist) {
		showError(w, http.StatusNotFound)
		return
	}
	if err != nil {
		showError(w, http.StatusInternalServerError)
		return
	}

	var page bytes.Buffer
	if err := tmpl.Execute(&page, data); err != nil {
		showError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(page.Bytes())
}

type ErrorData struct {
	Code  int
	Title string
	Hint  string
}

var errorTexts = map[int]ErrorData{
	400: {Title: "Bad request", Hint: "Use only English letters, numbers and symbols, and pick a font."},
	404: {Title: "Page not found", Hint: "Check the address, or go back to the form."},
	405: {Title: "Method not allowed", Hint: "Use the form on the home page to make art."},
	500: {Title: "Something went wrong", Hint: "The server had a problem. Try again in a moment."},
}

func showError(w http.ResponseWriter, code int) {
	data := errorTexts[code]
	data.Code = code

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(code), code)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	tmpl.Execute(w, data)
}