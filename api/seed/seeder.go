package seed

import (
	"github.com/trongtb88/rateservice/api/constants"
	"log"

	"github.com/trongtb88/rateservice/api/models"
	"github.com/trongtb88/rateservice/api/parser"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Load(db *gorm.DB) {
	log.Print("Reading xml file")
	xmlByte, err := parser.GetXML(constants.XML_URL)
	if err != nil {
		log.Fatalf("Can not read data from xml url")
	}
	currencyRates, err := parser.GetCurrencyRates(xmlByte)
	if err != nil {
		log.Fatalf("Parsing from byte data to target model failed  %v", err)
	}
	err = db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"rate"}),
	}).Model(&models.CurrencyRate{}).CreateInBatches(currencyRates, 50).Error
	if err != nil {
		log.Fatalf("cannot seed currency rate data into table: %v", err)
	}
	log.Println("Load data from url into database successfully.")
}


