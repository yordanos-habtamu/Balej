package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var relaySecret = []byte("balej-secure-relay-sector-key")

// SecureRelayService handles the generation and validation of cryogenic sharing tokens.
type SecureRelayService interface {
	GenerateToken(jobID string, issuerID uint) (string, error)
	VerifyToken(tokenStr string) (string, error)
}

type secureRelayService struct{}

func NewSecureRelayService() SecureRelayService {
	return &secureRelayService{}
}

// GenerateToken creates an HS256 signed relay token with a 24-hour expiration node.
func (s *secureRelayService) GenerateToken(jobID string, issuerID uint) (string, error) {
	claims := jwt.MapClaims{
		"job_id":    jobID,
		"issuer_id": issuerID,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(relaySecret)
}

// VerifyToken decrypts the relay token and returns the JobID if the node is still active.
func (s *secureRelayService) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return relaySecret, nil
	})

	if err != nil {
		return "", fmt.Errorf("relay_token_invalid: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jobID := claims["job_id"].(string)
		return jobID, nil
	}

	return "", fmt.Errorf("relay_token_expired_or_corrupt")
}
