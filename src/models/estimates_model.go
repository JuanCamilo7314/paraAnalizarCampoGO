package models

import (
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimateModel struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	Date                 time.Time          `json:"date" bson:"date"`
	NumTrees             int                `json:"numTrees" bson:"numTrees"`
	TotalFruitsEstimates int                `json:"totalFruitsEstimates" bson:"totalFruitsEstimates"`
	AverageFruits        int                `json:"averageFruits" bson:"averageFruits"`
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
	var totalFruits int
	for _, tree := range treesAssesed {
		totalFruits += estimate.calculateNumFruitsPerTree(tree)
	}

	average := float64(totalFruits) / float64(len(treesAssesed))
	estimate.AverageFruits = int(math.Round(average))
}

func (estimate *EstimateModel) setTotalFruitsEstimates(TreesFarmLot int) {
	totalFruitsEstimates := float64(estimate.AverageFruits) * float64(TreesFarmLot)
	estimate.TotalFruitsEstimates = int(math.Round(totalFruitsEstimates))
}

func (estimate *EstimateModel) setEstimateProduction(fruitWeight float32) {
	estimateProduction := float64(estimate.TotalFruitsEstimates) * float64(fruitWeight) / float64(1000)
	estimate.EstimatedProduction = int(math.Round(estimateProduction))
}

func (estimate *EstimateModel) calculateNumFruitsPerTree(treesAssese TreesAssessed) int {
	fruitPerTree := float64(treesAssese.NumFruits) * float64(4.0/treesAssese.NumQuartiles)
	return int(math.Round(fruitPerTree))
}
