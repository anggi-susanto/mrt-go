package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WasteWaterData represents waste water data
type WasteWaterData struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Timestamp          time.Time          `json:"timestamp" bson:"timestamp"`
	BOD                float64            `json:"BOD" bson:"BOD"`
	COD                float64            `json:"COD" bson:"COD"`
	TOC                float64            `json:"TOC" bson:"TOC"`
	DOC                float64            `json:"DOC" bson:"DOC"`
	OpticalBrighteners float64            `json:"Optical_Brighteners" bson:"Optical_Brighteners"`
	Ammonium           float64            `json:"Ammonium"`
	DissolvedOxygen    float64            `json:"Dissolved_Oxygen"`
	Nitrate            float64            `json:"Nitrate"`
	ECSalinityTDS      float64            `json:"EC_Salinity_TDS"`
	Pressure           float64            `json:"Pressure"`
	ORPRedox           float64            `json:"ORP_REDOX"`
	Turbidity          float64            `json:"Turbidity"`
	Chloride           float64            `json:"Chloride"`
	Coliforms          ColiformsData      `json:"Coliforms" bson:"Coliforms"`
	CrudeOils          float64            `json:"Crude_Oils"`
	PH                 float64            `json:"pH"`
	Tryptophan         float64            `json:"Tryptophan"`
	CDOM               float64            `json:"CDOM"`
	Temperature        float64            `json:"Temperature"`
	RefinedOils        float64            `json:"Refined_Oils"`
}

type WastewaterDataRequest struct {
	Timestamp          time.Time     `json:"timestamp" bson:"timestamp"`
	BOD                float64       `json:"BOD" bson:"BOD"`
	COD                float64       `json:"COD" bson:"COD"`
	TOC                float64       `json:"TOC" bson:"TOC"`
	DOC                float64       `json:"DOC" bson:"DOC"`
	OpticalBrighteners float64       `json:"Optical_Brighteners" bson:"Optical_Brighteners"`
	Ammonium           float64       `json:"Ammonium"`
	DissolvedOxygen    float64       `json:"Dissolved_Oxygen"`
	Nitrate            float64       `json:"Nitrate"`
	ECSalinityTDS      float64       `json:"EC_Salinity_TDS"`
	Pressure           float64       `json:"Pressure"`
	ORPRedox           float64       `json:"ORP_REDOX"`
	Turbidity          float64       `json:"Turbidity"`
	Chloride           float64       `json:"Chloride"`
	Coliforms          ColiformsData `json:"Coliforms" bson:"Coliforms"`
	CrudeOils          float64       `json:"Crude_Oils"`
	PH                 float64       `json:"pH"`
	Tryptophan         float64       `json:"Tryptophan"`
	CDOM               float64       `json:"CDOM"`
	Temperature        float64       `json:"Temperature"`
	RefinedOils        float64       `json:"Refined_Oils"`
}

// ColiformsData represents coliform data
type ColiformsData struct {
	Fecal float64 `json:"fecal" bson:"fecal"`
	EColi float64 `json:"E_coli" bson:"E_coli"`
	Total float64 `json:"total" bson:"total"`
}
