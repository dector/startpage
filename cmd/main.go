package main

// TODO
// 1. Display search bar
// 2. When user doing search - redirect to the Google page
// 3. Add tailwind
// 4. Make page nicer
// 5. Display version in the bottom

import (
	"embed"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//go:embed assets/*
var assets embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, _ := assets.ReadFile("assets/index.html")
		w.Write(file)
	})
	http.HandleFunc("/favicon.svg", func(w http.ResponseWriter, r *http.Request) {
		file, _ := assets.ReadFile("assets/favicon.svg")
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(file)
	})
	http.HandleFunc("POST /search", func(w http.ResponseWriter, r *http.Request) {
		query := strings.TrimSpace(r.FormValue("q"))

		redirectUrl := ""

		if IsUrl(query) {
			redirectUrl = query
		} else if query == "mail" {
			redirectUrl = "https://gmail.com"
		} else {
			redirectUrl = "https://www.google.com/search?q=" + query
		}

		w.Header().Set("Referrer-Policy", "no-referrer")
		http.Redirect(w, r, redirectUrl, http.StatusFound)
	})

	port := port()

	fmt.Printf("Started on port %s\n", port)
	http.ListenAndServe(port, nil)
}

func port() string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	if os.Getenv("DEV") != "" {
		return ":1110"
	}
	return ":1111"
}

func IsUrl(s string) bool {
	_, err := url.Parse(s)
	return err == nil
}
