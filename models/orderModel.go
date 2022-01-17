package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	OrderId   string             `json:"orderId"`
	OrderDate time.Time          `json:"orderDate" validate:"required"`
	TableId   *string            `json:"tableId" validate:"required"`
}
