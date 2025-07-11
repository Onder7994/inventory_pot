package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"top_ten/internal/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("static/login.html")
		if err != nil {
			http.Error(w, "HTML template error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		var foundUser string
		user := r.FormValue("user")
		pass := r.FormValue("pass")
		query := fmt.Sprintf("SELECT username FROM users WHERE username='%s' AND password='%s'", user, pass)
		row := utils.DB.QueryRow(query)

		err := row.Scan(&foundUser)

		if err != nil {
			http.Error(w, "Invalid Credintials", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
