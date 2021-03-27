package parser

import (
	"encoding/xml"
	"github.com/trongtb88/rateservice/api/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type xmlStruct struct {
	XMLName xml.Name `xml:"Envelope"`
	SoapEnv string   `xml:"xmlns:gesmes,attr"`
	Body  Cube       `xml:"Cube"`
}

type Cube struct {
	XMLName xml.Name                  `xml:"Cube"`
	CubeTypes []CubeType              `xml:"Cube"`
}

type CubeType struct {
	XMLName xml.Name                  `xml:"Cube"`
	Time  string                      `xml:"time,attr"`
	DetailCubeType []DetailCubeType     `xml:"Cube"`
}

type DetailCubeType struct {
	XMLName xml.Name                 `xml:"Cube"`
	Currency  string                 `xml:"currency,attr"`
	Rate      string                 `xml:"rate,attr"`
}

func GetXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func GetCurrencyRates(xmlBytes [] byte) ([]models.CurrencyRate, error){
	var resultXML xmlStruct
	err := xml.Unmarshal(xmlBytes, &resultXML)
	if err != nil {
		log.Println("Error on unmarshal data ", err)
		return nil, err
	}
	log.Print("Parsing done")
	var result []models.CurrencyRate
	const (
		layoutISO = "2006-01-02"
	)
	log.Print(len(resultXML.Body.CubeTypes))

	for _, cube := range resultXML.Body.CubeTypes {
		for _, cubeDetail := range cube.DetailCubeType {
			exchangeDate, err := time.Parse( layoutISO, cube.Time)
			if err != nil {
				log.Println("Can not convert from string to date, invalid datetime format ", err)
				return nil, err
			}
			rate, err := strconv.ParseFloat(cubeDetail.Rate, 64)
			if err != nil {
				log.Println("Can not convert from string to float int32, invalid string number format ", err)
				return nil, err
			}
			currencyRate := models.CurrencyRate{
				Date:       exchangeDate,
				CurrencyCode: cubeDetail.Currency,
				Rate:         rate,
			}
			result = append(result, currencyRate)
		}
	}

	log.Printf("Found %v curreny rates ", len(result))
	return result, nil
}
