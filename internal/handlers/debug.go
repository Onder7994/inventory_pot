package handlers

import (
	"fmt"
	"net/http"
)

func DebugHandler(w http.ResponseWriter, r *http.Request) {
	fakeEnv := map[string]string{
		"SSH_HOST": "test-hostname",
		"SSH_USER": "test_user",
		"SSH_PASS": "test_pass",
	}
	fmt.Fprintln(w, "Debug info:")
	for k, v := range fakeEnv {
		fmt.Fprintf(w, "%s=%s; ", k, v)
	}
}
