package helpers

import (
	"log"
	"os"
	"time"

	"github.com/PrinceNarteh/restaurant-management-api/models"
	"github.com/dgrijalva/jwt-go/v4"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserId    string
	jwt.StandardClaims
}

func GenerateAccessToken(user *models.User) string {
	claims := &SignedDetails{
		Email:     *user.Email,
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		UserId:    user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Minute * time.Duration(30)).Unix())),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
	}

	return accessToken
}
func GenerateRefreshToken(user *models.User) string {
	claims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * time.Duration(198)).Unix())),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
	}
	return refreshToken
}

func ValidateAccessToken(accessToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(accessToken, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = err.Error()
	}

	return claims, msg
}

func ValidateRefreshToken(refreshToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(refreshToken, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = err.Error()
	}

	return claims, msg
}
