package config

import (
	"app/config"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(email string, password string) error {
	hashedPassword := HashPassword(password)
	query := "INSERT INTO users (email, password_hash) VALUES ($1, $2)"
	_, err := config.DB.Exec(query, email, hashedPassword)
	return err
}

func AuthenticateUser(email string, password string) (string, error) {
	var storedHash string

	err := config.DB.QueryRow(
		"SELECT password_hash FROM users WHERE email = $1",
		email,
	).Scan(&storedHash)

	if err != nil || !ComparePassword(storedHash, password) {
		return "", err // Invalid credentials or DB error.
    }

    return GenerateJWT(email), nil // Generate token if valid.
}
