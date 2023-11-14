// auth.go
package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("OwM4ODSl0yaPuDg6vWiZPm1Mj2lw456j") // Replace with a strong secret key

// Claims represents the data to be encoded in the JWT.
type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

// GenerateToken creates a new JWT token for a user.
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyAndDecodeToken verifies and decodes a JWT token.
func VerifyAndDecodeToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
