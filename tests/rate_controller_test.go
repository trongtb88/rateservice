package tests

import (
	_ "bytes"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"net/http/httptest"
	_ "strconv"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/trongtb88/rateservice/api/controllers"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetLatestRate(t *testing.T) {
	err := refreshCurrencyRateTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedCurrencyRates()
	if err != nil {
		log.Fatalf("Cannot seed currency rate in db test %v\n", err)
	}
	req, err := http.NewRequest("GET", "/rates/latest", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetLatestRate)
	handler.ServeHTTP(rr, req)

	var response controllers.RateResonse

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, response.Rates[0].Key, "AUD")
	assert.Equal(t, response.Rates[0].Value, 1.2)
}

func TestGetRateBySpecificDate(t *testing.T) {
	err := refreshCurrencyRateTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedCurrencyRates()
	if err != nil {
		log.Fatalf("Cannot seed currency rate in db test %v\n", err)
	}
	req, err := http.NewRequest("GET", "/rates/", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"date": "2021-03-25"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetRateSpecificDate)
	handler.ServeHTTP(rr, req)

	var response controllers.RateResonse

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, response.Rates[0].Key, "AUD")
	assert.Equal(t, response.Rates[0].Value, 1.16)
	assert.Equal(t, response.Rates[1].Key, "CAD")
	assert.Equal(t, response.Rates[1].Value, 1.4)
}

func TestGetAnalyzeRate(t *testing.T) {
	err := refreshCurrencyRateTable()
	if err != nil {
		log.Fatal(err)
	}
	_, err = seedCurrencyRates()
	if err != nil {
		log.Fatalf("Cannot seed currency rate in db test %v\n", err)
	}
	req, err := http.NewRequest("GET", "/rates/analyze", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.AnalyzeRates)
	handler.ServeHTTP(rr, req)

	var response controllers.RateResonse

	err = json.Unmarshal([]byte(rr.Body.String()), &response)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}
	var actualResultMap = response.Rates[0].String()
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, response.Rates[0].Key, "AUD")
	assert.Equal(t, actualResultMap, "{AUD map[avg_rate:1.18 max_rate:1.2 min_rate:1.16]}")
}
