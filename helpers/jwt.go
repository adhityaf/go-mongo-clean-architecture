package helpers

import "github.com/dgrijalva/jwt-go"

var secretKey = "rahasia"

func GenerateToken(id, email, username string) string {
	claims := jwt.MapClaims{
		"id": id,
		"email": email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}