package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	OrderItemId string             `json:"orderItemId"`
	OrderId     string             `json:"orderId" validate:"required"`
	Size        *string            `json:"size" validate:"eq=S|eq=M|eq=L"`
	UnitPrice   *float64           `json:"unitPrice" validate:"required"`
	FoodId      *string            `json:"foodId" validate:"required"`
}
