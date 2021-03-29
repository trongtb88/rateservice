package tests

import (
	"fmt"
	"github.com/trongtb88/rateservice/api/constants"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/trongtb88/rateservice/api/controllers"
	"github.com/trongtb88/rateservice/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var server = controllers.Server{}
var rateModel = models.CurrencyRate{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()
	os.Exit(m.Run())

}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")
	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
		server.DB, err = gorm.Open(mysql.Open(DBURL))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", TestDbDriver)
		}
	}
}

func refreshCurrencyRateTable() error {
	err := server.DB.Debug().Migrator().DropTable(&models.CurrencyRate{})
	if err != nil {
		return err
	}
	err = server.DB.Debug().AutoMigrate(&models.CurrencyRate{})
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedCurrencyRates() ([]models.CurrencyRate, error) {

	err := refreshCurrencyRateTable()
	if err != nil {
		log.Fatal(err)
	}

	date1, err := time.Parse(constants.LayoutISO, "2021-03-26")
	date2, err := time.Parse(constants.LayoutISO, "2021-03-25")

	currencyRate1 := models.CurrencyRate{
		Date: date1,
		CurrencyCode: "AUD",
		Rate:    1.2,
	}

	currencyRate2 := models.CurrencyRate{
		Date: date1,
		CurrencyCode: "BGN",
		Rate:    1.3,
	}

	currencyRate3 := models.CurrencyRate{
		Date: date1,
		CurrencyCode: "CAD",
		Rate:    1.6,
	}

	currencyRate4 := models.CurrencyRate{
		Date: date2,
		CurrencyCode: "AUD",
		Rate:    1.16,
	}

	currencyRate5 := models.CurrencyRate{
		Date: date2,
		CurrencyCode: "BGN",
		Rate:    1.7,
	}

	currencyRate6 := models.CurrencyRate{
		Date: date2,
		CurrencyCode: "CAD",
		Rate:    1.4,
	}

	var rates [] models.CurrencyRate
	rates = append(rates, currencyRate1)
	rates = append(rates, currencyRate2)
	rates = append(rates, currencyRate3)
	rates = append(rates, currencyRate4)
	rates = append(rates, currencyRate5)
	rates = append(rates, currencyRate6)


	err = server.DB.Model(&models.CurrencyRate{}).Create(&rates).Error
	if err != nil {
		return nil, err
	}
	return rates, nil
}
