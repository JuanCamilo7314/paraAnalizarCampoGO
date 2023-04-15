package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimateModel struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	Date                 time.Time          `json:"date" bson:"date"`
	NumTrees             int                `json:"numTrees" bson:"numTrees"`
	TotalFruitsEstimates int                `json:"totalFruitsEstimates" bson:"totalFruitsEstimates"`
	AverageFruits        float32            `json:"averageFruits" bson:"averageFruits"`
	EstimatedProduction  int                `json:"estimatedProduction" bson:"estimatedProduction"`
	TreesAssessed        []TreesAssessed    `json:"treesAssessed" bson:"treesAssessed"`
}

type TreesAssessed struct {
	NumFruits    int `json:"numFruits"`
	NumQuartiles int `json:"numQuartiles"`
}

func (estimate *EstimateModel) CreateEstimation(treesAssesed []TreesAssessed, infFarmLot FarmLot) {
	estimate.NumTrees = len(treesAssesed)
	estimate.Date = time.Now()
	estimate.setAverageFruitsPerTree(treesAssesed)
	estimate.setTotalFruitsEstimates(infFarmLot.NumberTrees)
	estimate.setEstimateProduction(infFarmLot.AverageFruitWeight)
	estimate.TreesAssessed = treesAssesed
}

func (estimate *EstimateModel) setAverageFruitsPerTree(treesAssesed []TreesAssessed) {
	var totalFruits float32
	for _, tree := range treesAssesed {
		totalFruits += float32(estimate.calculateNumFruitsPerTree(tree))
	}

	estimate.AverageFruits = totalFruits / float32(len(treesAssesed))
}

func (estimate *EstimateModel) setTotalFruitsEstimates(TreesFarmLot int) {
	totalFruitsEstimates := int(float64(estimate.AverageFruits) * float64(TreesFarmLot))
	estimate.TotalFruitsEstimates = totalFruitsEstimates
}

func (estimate *EstimateModel) setEstimateProduction(fruitWeight float32) {
	estimateProduction := int(float64(estimate.TotalFruitsEstimates) * float64(fruitWeight))
	estimate.EstimatedProduction = estimateProduction
}

func (estimate *EstimateModel) calculateNumFruitsPerTree(treesAssese TreesAssessed) float32 {
	return float32(treesAssese.NumFruits) * float32(4/treesAssese.NumQuartiles)
}
