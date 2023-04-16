package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Harvest struct {
	ID                     primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Type                   string               `bson:"type,omitempty" json:"type,omitempty"`
	IDFarmLot              primitive.ObjectID   `bson:"idFarmLot,omitempty" json:"idFarmLot,omitempty"`
	EvaluationStartDate    time.Time            `bson:"evaluationStartDate,omitempty" json:"evaluationStartDate,omitempty"`
	EvaluationEndDate      time.Time            `bson:"evaluationEndDate,omitempty" json:"evaluationEndDate,omitempty"`
	Estimates              []primitive.ObjectID `bson:"estimates ,omitempty" json:"estimates ,omitempty"`
	SummaryFinalProduction primitive.ObjectID   `bson:"summaryFinalProduction,omitempty" json:"summaryFinalProduction,omitempty"`
}
