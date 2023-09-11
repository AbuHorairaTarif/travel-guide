package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// FlightController handles flight search and API integration.
type FlightController struct {
	web.Controller
}

// Define structs to match the new JSON response structure
type TripSegment struct {
	AircraftType         string `json:"aircraftType"`
	ArrivedAt            string `json:"arrivedAt"`
	CabinClassName       string `json:"cabinClassName"`
	DeparturedAt         string `json:"departuredAt"`
	Destination          Region `json:"destination"`
	Duration             int    `json:"duration"`
	EquipmentCode        string `json:"equipmentCode"`
	FlightNumber         string `json:"flightNumber"`
	IncludedCabinBaggage struct {
		Pieces     int    `json:"pieces"`
		Weight     int    `json:"weight"`
		WeightUnit string `json:"weightUnit"`
	} `json:"includedCabinBaggage"`
	IncludedCheckedBaggage interface{} `json:"includedCheckedBaggage"`
	MarketingCarrier       Carrier     `json:"marketingCarrier"`
	NumberOfTechnicalStops int         `json:"numberOfTechnicalStops"`
	OperatingCarrier       Carrier     `json:"operatingCarrier"`
	OperatingInformation   string      `json:"operatingInformation"`
	Origin                 Region      `json:"origin"`
	SegmentDetails         []struct {
		NumberOfSeatsLeft int    `json:"numberOfSeatsLeft"`
		PaxType           string `json:"paxType"`
	} `json:"segmentDetails"`
	SegmentId string `json:"segmentId"`
}

type Region struct {
	CityCode string `json:"cityCode"`
	CityName string `json:"cityName"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type FlightResponse struct {
	CurrentPage          int           `json:"currentPage"`
	Data                 Data          `json:"data"`
	FilteredFlightsCount int           `json:"filteredFlightsCount"`
	Flights              []TripSegment `json:"flights"`
	ID                   string        `json:"id"`
	IncludedCabinBaggage struct {
		Pieces     int    `json:"pieces"`
		Weight     int    `json:"weight"`
		WeightUnit string `json:"weightUnit"`
	} `json:"includedCabinBaggage"`
	IncludedCheckedBaggage interface{}   `json:"includedCheckedBaggage"`
	IncludedExtraProducts  []interface{} `json:"includedExtraProducts"`
	IsVI                   bool          `json:"isVI"`
	PaymentMethodPrices    []struct {
		Name  string `json:"name"`
		Price struct {
			Value int `json:"value"`
		} `json:"price"`
		Type string `json:"type"`
	} `json:"paymentMethodPrices"`
}

type Data struct {
	Type               string        `json:"__typename"`
	AvailableFilters   []interface{} `json:"availableFilters"`
	AvailableSortTypes []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"availableSortTypes"`
	CarrierCodes []interface{} `json:"carrierCodes"`
	CarrierNames []interface{} `json:"carrierNames"`
}

// Search handles flight search and API integration.
func (c *FlightController) Search() {
	// Parse request parameters, e.g., locationFrom, locationTo, departureDate, returnDate
	locationFrom := c.GetString("locationFrom")
	locationTo := c.GetString("locationTo")
	departureDate := c.GetString("departureDate")
	returnDate := c.GetString("returnDate")

	// Define the flight API URL
	apiURL := fmt.Sprintf("https://booking-com13.p.rapidapi.com/flights/return?location_from=%s&location_to=%s&departure_date=%s&return_date=%s&page=1&country_flag=us", locationFrom, locationTo, departureDate, returnDate)

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		logs.Error("Failed to create request: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	// Add headers to the request
	req.Header.Set("X-RapidAPI-Key", "16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df")
	req.Header.Set("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("Failed to make request: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		logs.Error("Received non-OK status code: %d", resp.StatusCode)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	// Parse the JSON response from the flight API
	var flightResults FlightResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&flightResults); err != nil {
		logs.Error("Failed to parse JSON response: %v", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	// Pass the flight search results to the view
	c.Data["FlightResults"] = flightResults

	// Render the flight search results view
	c.TplName = "flights.tpl"
}
