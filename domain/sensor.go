package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sensor struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	DeviceID    primitive.ObjectID `bson:"device_id" json:"device_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type SensorRequest struct {
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	DeviceID    primitive.ObjectID `bson:"device_id" json:"device_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}