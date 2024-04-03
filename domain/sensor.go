package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sensor struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	DeviceID    primitive.ObjectID `bson:"device_id" json:"device_id"`
	CreatedAt   MyTime             `bson:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05" time_utc:"true"`
	UpdatedAt   MyTime             `bson:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05" time_utc:"true"`
}

type SensorRequest struct {
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	DeviceID    primitive.ObjectID `bson:"device_id" json:"device_id"`
	CreatedAt   MyTime             `bson:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05" time_utc:"true"`
	UpdatedAt   MyTime             `bson:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05" time_utc:"true"`
}
