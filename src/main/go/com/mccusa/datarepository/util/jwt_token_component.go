package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"time"
)

// JwtTokenComponent handles JWT generation and validation using RSA keys.
type JwtTokenComponent struct {
	ClientID   string
	UserID     string
	Audition   string
	Expiration time.Duration
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewJwtTokenComponent creates a JwtTokenComponent and loads keys from files.
func NewJwtTokenComponent(
	clientID, userID, audition string,
	expirationSeconds int64,
	privateKeyPath, publicKeyPath string,
) (*JwtTokenComponent, error) {
	privBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read private key: %w", err)
	}
	pubBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read public key: %w", err)
	}
	privKey, err := parsePrivateKey(privBytes)
	if err != nil {
		return nil, err
	}
	pubKey, err := parsePublicKey(pubBytes)
	if err != nil {
		return nil, err
	}
	return &JwtTokenComponent{
		ClientID:   clientID,
		UserID:     userID,
		Audition:   audition,
		Expiration: time.Duration(expirationSeconds) * time.Millisecond,
		privateKey: privKey,
		publicKey:  pubKey,
	}, nil
}

// GenerateJwtToken returns a signed RS256 JWT.
func (j *JwtTokenComponent) GenerateJwtToken() (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"iss":   j.ClientID,
		"sub":   j.UserID,
		"aud":   j.Audition,
		"iat":   now.Unix(),
		"exp":   now.Add(j.Expiration).Unix(),
		"scope": "signature impersonation",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

// ValidateToken verifies the token signature and expiration.
func (j *JwtTokenComponent) ValidateToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.publicKey, nil
	})
	return err == nil && token.Valid
}

func parsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key PEM")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	privKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not RSA private key")
	}
	return privKey, nil
}

func parsePublicKey(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "PUBLIC KEY" && block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pubRsa, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return pubRsa, nil
	}
	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return pubKey, nil
}
