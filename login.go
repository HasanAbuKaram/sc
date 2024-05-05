package main

import (
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(content, "login.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
