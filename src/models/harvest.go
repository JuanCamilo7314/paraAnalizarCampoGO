package models

import (
	"errors"
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

type HarvestDetails struct {
	ID                     primitive.ObjectID `bson:"_id" json:"id"`
	Type                   string             `bson:"type" json:"type"`
	IDFarmLot              primitive.ObjectID `bson:"idFarmLot" json:"idFarmLot"`
	EvaluationStartDate    time.Time          `bson:"evaluationStartDate" json:"evaluationStartDate"`
	EvaluationEndDate      time.Time          `bson:"evaluationEndDate" json:"evaluationEndDate"`
	Estimates              []EstimateModel    `bson:"estimates" json:"estimates"`
	SummaryFinalProduction FinalProduction    `bson:"summaryFinalProduction" json:"summaryFinalProduction"`
}

type HarvestGraph struct {
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Year            int    `bson:"year,omitempty" json:"year,omitempty"`
	TotalProduction int    `bson:"totalProduction,omitempty" json:"totalProduction,omitempty"`
}

type ProductionGraph struct {
	TotalProduction int       `bson:"totalProduction,omitempty" json:"totalProduction,omitempty"`
	Date            time.Time `bson:"date,omitempty" json:"date,omitempty"`
}

type CreateHarvest struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type                string             `bson:"type,omitempty" json:"type,omitempty"`
	IDFarmLot           primitive.ObjectID `bson:"idFarmLot,omitempty" json:"idFarmLot,omitempty"`
	EvaluationStartDate string             `bson:"evaluationStartDate,omitempty" json:"evaluationStartDate,omitempty"`
	EvaluationEndDate   string             `bson:"evaluationEndDate,omitempty" json:"evaluationEndDate,omitempty"`
}

func (req *CreateHarvest) ValidateHarvest() error {

	if req.Type == "" {
		return errors.New("Type is required")
	}

	if req.IDFarmLot.String() == "" {
		return errors.New("IDFarmLot is required")
	}

	if req.EvaluationStartDate == "" {
		return errors.New("EvaluationStartDate is required")
	}

	if req.EvaluationEndDate == "" {
		return errors.New("EvaluationEndDate is required")
	}

	return nil
}
