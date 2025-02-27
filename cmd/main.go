package main

// TODO
// 1. Display search bar
// 2. When user doing search - redirect to the Google page
// 3. Add tailwind
// 4. Make page nicer
// 5. Display version in the bottom

import (
	"embed"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, _ := assets.ReadFile("assets/index.html")
		w.Write(file)
	})
	http.ListenAndServe(":1010", nil)
}
