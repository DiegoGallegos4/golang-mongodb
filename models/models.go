package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Gym gym type
type Gym struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
}

// Get DAO get gyms
func Get(ctx context.Context, db *mongo.Database) ([]Gym, error) {
	cursor, err := db.Collection("gyms").Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("Error fetching records: %v", err)
	}
	defer cursor.Close(ctx)

	records := make([]Gym, 0)
	for cursor.Next(ctx) {
		var record Gym
		cursor.Decode(&record)
		records = append(records, record)
	}

	return records, cursor.Err()
}
