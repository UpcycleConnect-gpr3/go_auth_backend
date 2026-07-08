package jwt

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	privateKeyPEM, err := os.ReadFile("private_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		log.Info("Failed to decode PEM block containing the private key")
	}

	privateKeyAny, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	var ok bool
	PrivateKey, ok = privateKeyAny.(*rsa.PrivateKey)
	if !ok {
		log.Info("Private key is not an RSA key")
	}

	publicKeyPEM, err := os.ReadFile("public_key.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ = pem.Decode(publicKeyPEM)
	if block == nil {
		log.Info("Failed to decode PEM block containing the public key")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	PublicKey, ok = publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		log.Info("Public key is not an RSA key")
	}
}

func GenerateJWT(userId string, role user_models.Role) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"role":   string(role),
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(PrivateKey)
}

func VerifyJWT(tokenString string) (string, user_models.Role, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("invalid token claims or not valid")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", "", fmt.Errorf("userId not found in token")
	}

	role, _ := claims["role"].(string)

	return userId, user_models.Role(role), nil
}

func Auth(w http.ResponseWriter, r *http.Request) (string, user_models.Role) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		response.NewErrorMessage(w, response.ErrAuthTokenRequired, http.StatusUnauthorized)
		return "", ""
	}

	userId, role, err := VerifyJWT(tokenString)
	if err != nil {
		response.NewErrorMessage(w, response.ErrInvalidAuthToken, http.StatusUnauthorized)
		return "", ""
	}

	return userId, role
}
