package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

// just practicing constants
const port string = ":8000"

var db *sql.DB

func main() {
	// SSL Mode disable
	pgUrl, err := pq.ParseURL("postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protetectEndpoint)).Methods("GET")

	// Log the listening port
	log.Println("Listen on port", port, "...")
	log.Fatal(http.ListenAndServe(port, router))
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error

	json.NewDecoder(r.Body).Decode(&user)

	// Uncomment to debug
	//spew.Dump(user)

	// Send an error if the Email is missing
	if user.Email == "" {
		error.Message = "Email is missing."
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	// Send an error if the Password is missing
	if user.Password == "" {
		error.Message = "Password is missing."
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	// Convert hash to string
	user.Password = string(hash)

	// Insert into DB
	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	//stmt := "select * from users RETURNING id;"

	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		error.Message = "Server Error."
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successfully called signup"))
}

func protetectEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked.")
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleWare invoked.")
	return nil
}
