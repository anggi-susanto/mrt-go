package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   MyTime             `bson:"created_at" json:"created_at" time_format:"2006-01-02T15:04:05"`
	UpdatedAt   MyTime             `bson:"updated_at" json:"updated_at" time_format:"2006-01-02T15:04:05"`
}

type DeviceRequest struct {
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	CreatedAt   MyTime `bson:"created_at" json:"created_at" time_format:"2006-01-02T15:04:05"`
	UpdatedAt   MyTime `bson:"updated_at" json:"updated_at" time_format:"2006-01-02T15:04:05"`
}

type MyTime struct {
	time.Time
}

func (m *MyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	*m = MyTime{tt}
	return err
}
