package handlers

import (
	"html/template"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("static/admin.html")
		if err != nil {
			http.Error(w, "HTML template error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
