package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Catalog is used to represent Catalog profile data
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Image       string             `bson:"image,omitempty"`
	Description string             `bson:"description,omitempty"`
	Price       int                `bson:"price,omitempty"`
	CreatedOn   time.Time          `bson:"createdon,omitempty"`
}

type InsertResponse struct {
	Message   string `json:"message"`
	Insert_id string `json:"insert_id"`
}
