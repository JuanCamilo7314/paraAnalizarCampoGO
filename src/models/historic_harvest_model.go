package models

type HistoricHarvest struct {
	Harvest         Harvest         `json:"harvest"`
	Estimates       []EstimateModel `json:"estimates"`
	FinalProduction FinalProduction `json:"finalProduction"`
}

type HistoricHarvestPipeline struct {
	Harvest []HarvestDetails `json:"harvest"`
}
