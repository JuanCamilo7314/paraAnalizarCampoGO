package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqEstimate struct {
	IdFarmLot     string          `json:"idFarm"`
	TreesAssessed []TreesAssessed `json:"treesAssessed"`
}

func (req *ReqEstimate) ValidateEstimate() error {

	if req.IdFarmLot == "" {
		return errors.New("IdFarmLot is required")
	}

	if len(req.TreesAssessed) == 0 || req.TreesAssessed == nil {
		return errors.New("TreesAssessed is required")
	}

	return nil
}

type ReqIdsEstimates struct {
	Ids []primitive.ObjectID `json:"ids"`
}
