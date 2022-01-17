package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	TableId        string             `json:"tableId"`
	TableNumber    *int               `json:"tableNumber" validate:"required"`
	NumberOfGuests *int               `json:"numberOfGuests"`
}
