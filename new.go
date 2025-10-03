package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// User represents a user in the system.
type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

// In-memory user store
var users = []User{
	{ID: 1, Username: "alice", Email: "alice@example.com", CreatedAt: time.Now().Add(-48 * time.Hour)},
	{ID: 2, Username: "bob", Email: "bob@example.com", CreatedAt: time.Now().Add(-24 * time.Hour)},
}

// getUsersHandler returns all users as a JSON response.
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "[")
	for i, user := range users {
		fmt.Fprintf(w, `{"id":%d,"username":"%s","email":"%s","created_at":"%s"}`,
			user.ID, user.Username, user.Email, user.CreatedAt.Format(time.RFC3339))
		if i < len(users)-1 {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprint(w, "]")
}

// createUserHandler creates a new user from query parameters.
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")
	if username == "" || email == "" {
		http.Error(w, "Missing username or email", http.StatusBadRequest)
		return
	}
	id := len(users) + 1
	user := User{
		ID:        id,
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id":%d,"username":"%s","email":"%s","created_at":"%s"}`,
		user.ID, user.Username, user.Email, user.CreatedAt.Format(time.RFC3339))
}

// homeHandler displays a welcome message.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go User API!")
}

// logRequest is a middleware that logs incoming requests.
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", getUsersHandler)
	mux.HandleFunc("/create", createUserHandler)

	log.Printf("Starting server on :%s...", port)
	err := http.ListenAndServe(":"+port, logRequest(mux))
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}