package controllers

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/trongtb88/rateservice/api/constants"
	"github.com/trongtb88/rateservice/api/lib"
	"github.com/trongtb88/rateservice/api/responses"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}
type CurrencyRate struct {
	CurrencyCode string  `json:"currency_code"`
	Rate          float64 `json:"rate"`
}

type AnalyzedCurrencyRate struct {
	CurrencyCode     string      `json:"currency_code"`
	MinRate          float64 `json:"min_rate"`
	MaxRate          float64 `json:"max_rate"`
	AvgRate			 float64 `json:"avg_rate"`
}

type AnalyzedDetailCurrencyRate struct {
	MinRate          float64 `json:"min_rate"`
	MaxRate          float64 `json:"max_rate"`
	AvgRate			 float64 `json:"avg_rate"`
}

type RateResonse struct {
	Base string                 `json:"base"`
	Rates lib.MapSlice          `json:"rates"` // We  have to use MapSlice here to keep order json encoding,
	// we cam simple use map[string]float but the order json will not maintained
}

func (server * Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to rate system")
}

func (server * Server) GetLatestRate(w http.ResponseWriter, r *http.Request) {
	var currencyRates [] CurrencyRate
	err := server.DB.Raw("SELECT currency_code, rate FROM currency_rates WHERE date = ( SELECT max(date) FROM currency_rates) order by rate ").Scan(&currencyRates).Error
	if err != nil {
		log.Printf("Error on select latest rates %v", err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	ms := lib.MapSlice{
	}
	for _, currencyRate :=  range currencyRates {
		mapItem :=  lib.MapItem{
			Key: currencyRate.CurrencyCode,
			Value: currencyRate.Rate,
		}
		ms = append(ms, mapItem)
	}

	var rateResp = RateResonse{
		Base: "EUR",
		Rates: ms,
	}
	responses.JSON(w, http.StatusOK, rateResp)
}

func (server * Server) GetRateSpecificDate(w http.ResponseWriter, r *http.Request) {
	var currencyRates [] CurrencyRate

	vars := mux.Vars(r)
	specificDate := strings.TrimSpace(vars["date"])
	if len(specificDate) == 0 {
		responses.ERROR(w, http.StatusBadRequest, errors.New("Missing Required Field."))
	}
	exchangeDate, err := time.Parse(constants.LayoutISO, specificDate)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	err = server.DB.Raw("SELECT currency_code, rate from currency_rates where date = ? order by rate ", exchangeDate).Scan(&currencyRates).Error
	if err != nil {
		log.Printf("Error on select rates by date %v", err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	ms := lib.MapSlice{
	}
	for _, currencyRate :=  range currencyRates {
		mapItem :=  lib.MapItem{
			Key: currencyRate.CurrencyCode,
			Value: currencyRate.Rate,
		}
		ms = append(ms, mapItem)
	}

	var rateResp = RateResonse{
		Base: "EUR",
		Rates: ms,
	}
	responses.JSON(w, http.StatusOK, rateResp)
}

func (server * Server) AnalyzeRates(w http.ResponseWriter, r *http.Request) {
	var analyzedCurrencyRates [] AnalyzedCurrencyRate
	err := server.DB.Raw("select currency_code, max(rate) as 'max_rate', min(rate) as 'min_rate', avg(rate) as 'avg_rate' " +
		" from currency_rates group by currency_code ").Scan(&analyzedCurrencyRates).Error
	if err != nil {
		log.Printf("Error on analyze rates by date %v", err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	ms := lib.MapSlice{}
	for _, analyzedCurrencyRates :=  range analyzedCurrencyRates {
		mapItem :=  lib.MapItem{
			Key: analyzedCurrencyRates.CurrencyCode,
			Value:  AnalyzedDetailCurrencyRate{
				MinRate: analyzedCurrencyRates.MinRate,
				MaxRate: analyzedCurrencyRates.MaxRate,
				AvgRate: analyzedCurrencyRates.AvgRate,
			},
		}
		ms = append(ms, mapItem)
	}

	var rateResp = RateResonse{
		Base: "EUR",
		Rates: ms,
	}
	responses.JSON(w, http.StatusOK, rateResp)
}
