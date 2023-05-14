package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmLot struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	NameLot        string             `bson:"nameLot" json:"nameLot"`
	NumberTrees    int                `bson:"numberTrees" json:"numberTrees"`
	TreesAge       float32            `bson:"treesAge" json:"treesAge"`
	ProductionDate struct {
		Primary struct {
			Initial time.Time `bson:"initial" json:"initial"`
			Final   time.Time `bson:"final" json:"final"`
		} `bson:"primary" json:"primary"`
		Secondary struct {
			Initial time.Time `bson:"initial" json:"initial"`
			Final   time.Time `bson:"final" json:"final"`
		} `bson:"secondary" json:"secondary"`
	} `bson:"productionDate" json:"productionDate"`

	AverageFruitWeight float32 `bson:"averageFruitWeight" json:"averageFruitWeight"`
}

type FarmLotReq struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	NameLot        string             `bson:"nameLot" json:"nameLot"`
	NumberTrees    int                `bson:"numberTrees" json:"numberTrees"`
	TreesAge       float32            `bson:"treesAge" json:"treesAge"`
	ProductionDate struct {
		Primary struct {
			Initial string `bson:"initial" json:"initial"`
			Final   string `bson:"final" json:"final"`
		} `bson:"primary" json:"primary"`
		Secondary struct {
			Initial string `bson:"initial" json:"initial"`
			Final   string `bson:"final" json:"final"`
		} `bson:"secondary" json:"secondary"`
	} `bson:"productionDate" json:"productionDate"`

	AverageFruitWeight float32 `bson:"averageFruitWeight" json:"averageFruitWeight"`
}
