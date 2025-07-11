package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type LogEntry struct {
	Time       string              `json:"time"`
	Method     string              `json:"method"`
	Path       string              `json:"path"`
	RemoteAddr string              `json:"remote_addr"`
	UserAgent  string              `json:"user_agent"`
	Query      map[string][]string `json:"query_params"`
	Form       map[string][]string `json:"body,omitempty"`
}

var logFile *os.File

func InitLogger() {
	log_file_path := os.Getenv("APP_LOG_FILE")

	if log_file_path == "" {
		log_file_path = "/var/log/honeypot/iventory_pot.json"
	}

	var err error
	logFile, err = os.OpenFile(log_file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var formData map[string][]string
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err == nil {
				formData = r.PostForm
			}
		}
		entry := LogEntry{
			Time:       time.Now().Format(time.RFC3339),
			Method:     r.Method,
			Path:       r.URL.Path,
			RemoteAddr: r.RemoteAddr,
			UserAgent:  r.UserAgent(),
			Query:      r.URL.Query(),
			Form:       formData,
		}
		jsonEntry, err := json.Marshal(entry)
		if err == nil {
			logFile.Write(jsonEntry)
			logFile.Write([]byte("\n"))
		}
		next.ServeHTTP(w, r)
	})
}
