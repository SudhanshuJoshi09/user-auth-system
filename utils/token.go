package utils

import (
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// GenerateLoginToken generates a JWT for the given email and jwtSecret.
func GenerateLoginToken(email string, jwtSecret string) (string, error) {
	secretKey := []byte(jwtSecret)

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// DecodeLoginToken decodes a JWT token and returns the email claim.
func DecodeLoginToken(loginToken, jwtSecret string) (string, error) {
	secretKey := []byte(jwtSecret)

	// Parse the token
	token, err := jwt.Parse(loginToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method matches the expected method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// Return the secret key to verify the token's signature
		return secretKey, nil
	})

	// If there's an error or if the token is not valid, return an error
	if err != nil {
		return "", err
	}

	// Check if the token is valid and has claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the email from the claims
		email, ok := claims["email"].(string)
		if !ok {
			return "", errors.New("email claim missing or invalid")
		}

		// Check if the token has expired
		exp := claims["exp"].(float64)
		if time.Now().After(time.Unix(int64(exp), 0)) {
			return "", errors.New("token has expired")
		}

		return email, nil
	}

	return "", errors.New("invalid token")
}

