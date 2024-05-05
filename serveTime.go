package main

import (
	"html/template"
	"net/http"
	"time"
)

func serverTime(w http.ResponseWriter, _ *http.Request) {
	currentTime := time.Now().Format("15:04:05")

	tmpl, err := template.ParseFS(content, "index.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, currentTime)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
