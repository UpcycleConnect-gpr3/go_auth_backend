package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	privateKeyPEM, err := os.ReadFile("private_key.pem")
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing the private key")
	}

	privateKeyAny, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse private key (PKCS#8): %v", err)
	}

	var ok bool
	privateKey, ok = privateKeyAny.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("Private key is not an RSA key")
	}

	publicKeyPEM, err := os.ReadFile("public_key.pem")
	if err != nil {
		log.Fatalf("Failed to read public key: %v", err)
	}
	block, _ = pem.Decode(publicKeyPEM)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing the public key")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("Failed to parse public key (PKIX): %v", err)
	}

	publicKey, ok = publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Public key is not an RSA key")
	}
}

func GenerateJWT(userId string) (string, error) {

	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims or not valid")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return "", fmt.Errorf("userId not found in token")
	}
	return userId, nil
}

func Auth(w http.ResponseWriter, r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return false
	}
	_, err := VerifyJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return false
	}

	return true
}
