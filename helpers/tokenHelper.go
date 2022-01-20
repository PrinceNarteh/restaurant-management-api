package helpers

import "github.com/PrinceNarteh/restaurant-management-api/models"

func GenerateAccessToken(user *models.User) string {
	return "Access Token"
}
func GenerateRefreshToken(user *models.User) string {
	return "Refresh Token"
}
func ValidateAccessToken()  {}
func ValidateRefreshToken() {}
