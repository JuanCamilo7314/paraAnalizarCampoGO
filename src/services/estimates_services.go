package services

import (
	"AgroXpert-Backend/src/models"
	"time"
)

func CreateEstimate(estimateReq models.ReqEstimate) (models.EstimateModel, error) {
	newEstimate := models.EstimateModel{}

	err := calculateEstimate(estimateReq, &newEstimate)
	if err != nil {
		return newEstimate, err
	}

	newEstimate.Date = time.Now()
	newEstimate.TreesAssessed = estimateReq.TreesAssessed
	return newEstimate, nil
}

func calculateEstimate(estimateReq models.ReqEstimate, estimateResult *models.EstimateModel) error {
	infFarmLot, err := GetOneFarmLot(estimateReq.IdFarmLot)
	if err != nil {
		return err
	}

	estimateResult.NumTrees = len(estimateReq.TreesAssessed)
	averageFruits := averageFruitsPerTree(estimateReq.TreesAssessed)
	estimateResult.AverageFruits = averageFruits
	totalFruitsEstimates := totalFruitsEstimates(averageFruits, infFarmLot.NumberTrees)
	estimateResult.TotalFruitsEstimates = totalFruitsEstimates
	estimateResult.EstimatedProduction = estimateProduction(totalFruitsEstimates, infFarmLot.AverageFruitWeight)

	return nil
}

func averageFruitsPerTree(treesAssesed []models.TreesAssessed) float64 {
	var totalFruits float64
	for _, tree := range treesAssesed {
		totalFruits += float64(calculateNumFruitsPerTree(tree))
	}

	return totalFruits / float64(len(treesAssesed))
}

func totalFruitsEstimates(averageFruits float64, TreesFarmLot int) int {
	return int(averageFruits * float64(TreesFarmLot))
}

func estimateProduction(totalFruitsEstimates int, fruitWeight float32) int {
	return int(float32(totalFruitsEstimates)*fruitWeight) / 1000
}

func calculateNumFruitsPerTree(treesAssese models.TreesAssessed) int {
	return treesAssese.NumFruits * (4 / treesAssese.NumQuartiles)
}
