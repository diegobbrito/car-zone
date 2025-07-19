package login

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/diegobbrito/car-zone/models"
	"github.com/golang-jwt/jwt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var credentials models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	valid := (credentials.Username == "admin" && credentials.Password == "password")

	if !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokenString, err := GenerateToken(credentials.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		log.Println("Error generating token:", err)
		return
	}

	response := map[string]string{
		"token": tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
