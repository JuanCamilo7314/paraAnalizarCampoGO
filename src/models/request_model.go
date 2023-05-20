package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReqEstimate struct {
	IdFarmLot     string          `json:"idFarm"`
	IdHarvest     string          `json:"idHarvest"`
	TreesAssessed []TreesAssessed `json:"treesAssessed"`
}

func (req *ReqEstimate) ValidateEstimate() error {

	if req.IdFarmLot == "" {
		return errors.New("IdFarmLot is required")
	}

	if req.IdHarvest == "" {
		return errors.New("IdHarvest is required")
	}

	if len(req.TreesAssessed) == 0 || req.TreesAssessed == nil {
		return errors.New("TreesAssessed is required")
	}

	if err := req.validateTressInfo(); err != nil {
		return err
	}

	return nil
}

func (req *ReqEstimate) validateTressInfo() error {
	for _, tree := range req.TreesAssessed {
		if tree.NumFruits < 0 {
			return errors.New("Number Fruits must be greater than 0")
		}

		if tree.NumQuartiles <= 0 && tree.NumQuartiles > 4 {
			return errors.New("Number of Quartiles must be greater than 0 and less than 4")
		}
	}

	return nil
}

type ReqIdsEstimates struct {
	Ids []primitive.ObjectID `json:"ids"`
}
