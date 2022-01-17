package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	CreatedAt    time.Time          `json:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt"`
	UserId       string             `json:"userId"`
	FirstName    *string            `json:"firstName" validate:"required,min=2,max=100"`
	LastName     *string            `json:"lastName" validate:"required,min=2,max=100"`
	Email        *string            `json:"email" validate:"required,email"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Avatar       *string            `json:"avatar" validate:"required"`
	Phone        *string            `json:"phone"`
	AccessToken  *string            `json:"accessToken"`
	RefreshToken *string            `json:"refreshToken"`
}
