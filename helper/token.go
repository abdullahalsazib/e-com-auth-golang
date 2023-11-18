package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/neerajbg/go-gin-auth/model"
)

type CustomClaims struct {
	Email  string
	UserId uint

	jwt.RegisteredClaims
}

var secret string = "secret"

func GenerateToken(user model.User) (string, error) {

	claims := CustomClaims{
		user.Email,
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in token singing.", err)
		return "", err
	}

	return t, nil

}

// Validate Token
func ValidateToken(clientToken string) (claims *CustomClaims, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		msg = err.Error()
		return
	}

	return claims, msg
}
