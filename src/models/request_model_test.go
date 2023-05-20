package models

import (
	"testing"
)

func TestValidateEstimateOk(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot: "1234",
		IdHarvest: "1234",
		TreesAssessed: []TreesAssessed{
			{
				NumFruits:    10,
				NumQuartiles: 4,
			},
		},
	}

	err := estimate.ValidateEstimate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
}

func TestValidateEstimateIdFarmLotEmpty(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot: "",
		IdHarvest: "1234",
		TreesAssessed: []TreesAssessed{
			{
				NumFruits:    10,
				NumQuartiles: 4,
			},
		},
	}

	err := estimate.ValidateEstimate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestValidateEstimateIdHarvestEmpty(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot: "1234",
		IdHarvest: "",
		TreesAssessed: []TreesAssessed{
			{
				NumFruits:    10,
				NumQuartiles: 4,
			},
		},
	}

	err := estimate.ValidateEstimate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestValidateEstimateTreesAssessedEmpty(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot:     "1234",
		IdHarvest:     "1234",
		TreesAssessed: nil,
	}

	err := estimate.ValidateEstimate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestValidateEstimateTreesAssessedNumFruitsNegative(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot: "1234",
		IdHarvest: "1234",
		TreesAssessed: []TreesAssessed{
			{
				NumFruits:    -1,
				NumQuartiles: 4,
			},
		},
	}

	err := estimate.ValidateEstimate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestValidateEstimateTreesAssessedNumQuartilesInvalid(t *testing.T) {
	estimate := ReqEstimate{
		IdFarmLot: "1234",
		IdHarvest: "1234",
		TreesAssessed: []TreesAssessed{
			{
				NumFruits:    10,
				NumQuartiles: 5,
			},
		},
	}

	err := estimate.ValidateEstimate()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
