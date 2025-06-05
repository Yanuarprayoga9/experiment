package main

import (
	"fmt"
	"log"
	"net/http"


)

func main() {
	http.HandleFunc("/login", loginHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !authenticateLDAP(username, password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Check permission with Casbin
	if !checkPermission(username, "/dashboard", "read") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Welcome %s! Access granted to /dashboard", username)
}
