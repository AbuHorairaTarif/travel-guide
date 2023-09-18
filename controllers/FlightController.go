// package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"travelguide/models"

// 	"github.com/beego/beego/v2/core/logs"
// 	beego "github.com/beego/beego/v2/server/web"
// )

// // FlightController handles flight search and API integration.
// type FlightController struct {
// 	beego.Controller
// }

// type Price struct {
// 	Currency struct {
// 		Code string `json:"code"`
// 	} `json:"currency"`
// 	Value float64 `json:"value"`
// }

// type SellPriceBaggage struct {
// 	MaxWeight     int    `json:"maxWeight"`
// 	NumberOfUnits int    `json:"numberOfUnits"`
// 	WeightUnit    string `json:"weightUnit"`
// }

// type SellSpecification struct {
// 	SellPriceBaggage []SellPriceBaggage `json:"sellPriceBaggage"`
// }

// type TravelerPrice struct {
// 	Price struct {
// 		Price Price `json:"price"`
// 		VAT   struct {
// 			Value float64 `json:"value"`
// 		} `json:"vat"`
// 	} `json:"price"`
// }

// type Flight struct {
// 	AvailableExtraProducts []struct {
// 		SellSpecification SellSpecification `json:"sellSpecification"`
// 	} `json:"availableExtraProducts"`
// 	ID             string          `json:"id"`
// 	ShareableURL   string          `json:"shareableUrl"`
// 	TravelerPrices []TravelerPrice `json:"travelerPricesWithoutPaymentDiscounts"`
// 	Type           string          `json:"type"`
// }

// type FlightStruct struct {
// 	CurrentPage  int `json:"currentPage"`
// 	FlightStruct struct {
// 		Flights []Flight `json:"flights"`
// 	} `json:"data"`
// }

// type FlightInfo struct {
// 	Id            string
// 	Url           string
// 	Price         float64
// 	Vat           float64
// 	Currency      string
// 	MaxWeight     int
// 	NumberOfUnits int
// }

// type FlightData struct {
// 	Origin      string
// 	Destination string
// 	Departure   string
// 	Return      string
// 	Flights     []FlightInfo
// }

// type ErrorMessage struct {
// 	Message    string
// 	SubMessage string
// 	Code       int
// }

// // Search handles flight search and API integration.

// func fetchFlight(apiKey, apiHost, locationFrom, locationTo, departureDate, returnDate string, data models.FlightData) {
// 	apiURL := fmt.Sprintf("https://%s/flights/return?location_from=%s&location_to=%s&departure_date=%s&return_date=%s&page=1&country_flag=us", apiHost, locationFrom, locationTo, departureDate, returnDate)

// 	req, err := http.NewRequest("GET", apiURL, nil)
// 	if err != nil {
// 		logs.Error("Failed to create request: %v", err)
// 		data.JsonChan <- "" // Send an empty response on error
// 		return
// 	}

// 	req.Header.Add("X-RapidAPI-Key", apiKey)
// 	req.Header.Add("X-RapidAPI-Host", apiHost)

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		logs.Error("Failed to make request: %v", err)
// 		data.JsonChan <- "Failed to make request " // Send an empty response on error
// 		return
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		logs.Error("Received non-OK status code: %d", res.StatusCode)
// 		data.JsonChan <- "Received non-OK status code" // Send an empty response on error
// 		return
// 	}

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		logs.Error("Error reading response body: %v", err)
// 		data.JsonChan <- "Error reading response body" // Send an empty response on error
// 		return
// 	}
// 	data.JsonChan <- string(body) // Send the JSON response to the channel
// }

// func (c *FlightController) Get() {
// 	apiKey, _ := beego.AppConfig.String("X-RapidAPI-Key")
// 	apiHost, _ := beego.AppConfig.String("X-RapidAPI-Host")

// 	// Parse request parameters, e.g., locationFrom, locationTo, departureDate, returnDate
// 	locationFrom := c.GetString("locationFrom")
// 	locationTo := c.GetString("locationTo")
// 	departureDate := c.GetString("departureDate")
// 	returnDate := c.GetString("returnDate")

// 	if locationFrom != "" && locationTo != "" && departureDate != "" && returnDate != "" {
// 		data := models.FlightStruct{}
// 		// Call fetchFlight with the required arguments
// 		jsonChan := make(chan string)
// 		flightData := models.FlightData{}
// 		flightData.Origin = locationFrom
// 		flightData.Destination = locationTo
// 		flightData.Departure = departureDate
// 		flightData.Return = returnDate
// 		flightData.Flights = []models.FlightInfo{}
// 		// flightData.JsonChan = jsonChan

// 		// Pass the individual string parameters to fetchFlight using goroutine
// 		go fetchFlight(apiKey, apiHost, locationFrom, locationTo, departureDate, returnDate, flightData)

// 		jsonString := <-jsonChan

// 		err := json.Unmarshal([]byte(jsonString), &data)
// 		if err != nil {
// 			fmt.Println("Error unmarshaling JSON:", err)
// 		}
// 		c.Data["json"] = &data

// 	}
// 	c.ServeJSON()
// }

package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"travelguide/models"

	beego "github.com/beego/beego/v2/server/web"
)

type FlightController struct {
	beego.Controller
}

func (c *FlightController) Get() {
	source := c.GetString("location_from")
	c.Data["location_from"] = source
	destination := c.GetString("location_to")
	c.Data["location_to"] = destination
	departure_date := c.GetString("departure_date")
	c.Data["departure_date"] = departure_date
	return_date := c.GetString("return_date")
	c.Data["return_date"] = return_date

	if source == "" || destination == "" || departure_date == "" || return_date == "" {
		c.Data["Error"] = "Please Fill the all Required Field"
		c.Redirect("/", 302)
		return
	}

	url := "https://booking-com13.p.rapidapi.com/flights/return" +
		"?location_from=" + source +
		"&location_to=" + destination +
		"&departure_date=" + departure_date +
		"&return_date=" + return_date +
		"&page=1&country_flag=us"

	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["Error"] = "Error creating request"
		return
	}

	req.Header.Add("X-RapidAPI-Key", "16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df")
	req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

	flightDataChan := make(chan models.FlightData)

	go func() {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			c.Data["Error"] = "Error making the request"
			flightDataChan <- models.FlightData{}
			return
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.Data["Error"] = "Error reading the response"
			flightDataChan <- models.FlightData{}
			return
		}

		var allFlights models.FlightData
		if err = json.Unmarshal(body, &allFlights); err != nil {
			c.Data["Error"] = "Error parsing JSON response"
			flightDataChan <- models.FlightData{}
			return
		}

		flightDataChan <- allFlights
	}()

	extractedData := <-flightDataChan
	c.Data["Flights"] = extractedData.Data

	fmt.Println("Total flights", extractedData.Data.FilteredFlightsCount)
	c.Data["Length"] = len(extractedData.Data.Flights)

	c.TplName = "flights.tpl"
}
