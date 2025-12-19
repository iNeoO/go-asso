package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/config"
	"golang.org/x/crypto/bcrypt"
)

type TokenGenerated struct {
	Token     string
	ExpiresAt time.Time
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateAuthToken(email string, userID uuid.UUID) (*TokenGenerated, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	exp := time.Now().Add(time.Minute * 10)
	claims["email"] = email
	claims["user_id"] = userID
	claims["exp"] = exp.Unix()

	t, err := token.SignedString([]byte(config.Config("AUTH_SECRET")))
	if err != nil {
		return nil, err
	}
	return &TokenGenerated{Token: t, ExpiresAt: exp}, nil
}

func GenerateRefreshToken(email string, userID uuid.UUID) (*TokenGenerated, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	exp := time.Now().Add(time.Hour * 72)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["user_id"] = userID
	claims["exp"] = exp.Unix()

	t, err := token.SignedString([]byte(config.Config("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return &TokenGenerated{Token: t, ExpiresAt: exp}, nil
}
