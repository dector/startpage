package main

import (
	"embed"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dector/startpage/cmd/ui"
	"github.com/dector/startpage/internal"
)

//go:embed assets/*
var assets embed.FS

func main() {
	config, _ := internal.LoadConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ui.IndexPage(*config).Render(r.Context(), w)
	})
	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		ui.HelpPage(*config).Render(r.Context(), w)
	})
	http.HandleFunc("/favicon.svg", func(w http.ResponseWriter, r *http.Request) {
		file, _ := assets.ReadFile("assets/favicon.svg")
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(file)
	})
	http.HandleFunc("POST /search", func(w http.ResponseWriter, r *http.Request) {
		query := strings.TrimSpace(r.FormValue("q"))

		redirectUrl := ""

		if query == "/reload" {
			config, _ = internal.LoadConfig()
			redirectUrl = "/"
		} else if query == "/help" {
			redirectUrl = "/help"
		} else if IsUrl(query) {
			redirectUrl = query
		} else if url, exists := config.Redirects[query]; exists {
			redirectUrl = url
		} else {
			redirectUrl = strings.ReplaceAll(config.SearchEngine, "%%query%%", query)
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
	parsedUrl, _ := url.Parse(s)

	scheme := strings.ToLower(parsedUrl.Scheme)
	return scheme == "http" || scheme == "https"
}
