package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	MenuId    string             `json:"menuId"`
	Name      string             `json:"name" validate:"required"`
	Category  string             `json:"category" validate:"required"`
	StartDate *time.Time         `json:"startDate"`
	EndDate   *time.Time         `json:"endDate"`
}
