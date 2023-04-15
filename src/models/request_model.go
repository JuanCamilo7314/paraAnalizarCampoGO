package models

type ReqEstimate struct {
	IdFarmLot     string          `json:"idFarm"`
	TreesAssessed []TreesAssessed `json:"treesAssessed"`
}
