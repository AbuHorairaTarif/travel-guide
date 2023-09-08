package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type HotelSearchController struct {
	web.Controller
}

type HotelSearchResponse struct {
	Status    bool                `json:"status"`
	Message   interface{}         `json:"message"` // Change the type to interface{}
	Timestamp int64               `json:"timestamp"`
	Data      []HotelSearchResult `json:"data"`
}

type HotelSearchResult struct {
	Title         string `json:"title"`
	GeoID         string `json:"geoId"`
	DocumentID    string `json:"documentId"`
	TrackingItems string `json:"trackingItems"`
	SecondaryText string `json:"secondaryText"`
	Image         struct {
		URLTemplate string `json:"urlTemplate"`
	} `json:"image"`
}

func (c *HotelSearchController) Get() {
	// Get query parameters for hotel search
	location := c.GetString("location")
	checkInDate := c.GetString("checkInDate")
	checkOutDate := c.GetString("checkOutDate")
	guests := c.GetString("guests")

	// Define the TripAdvisor API URL for hotel search
	url := fmt.Sprintf("https://tripadvisor16.p.rapidapi.com/api/v1/hotels/searchHotels?geoId=%s&checkIn=%s&checkOut=%s&guests=%s&pageNumber=1&currencyCode=USD",
		location, checkInDate, checkOutDate, guests)

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.WriteString("Error creating request")
		return
	}

	// Add the required headers for authentication
	req.Header.Add("X-RapidAPI-Key", "16e4aeb170mshafe993eedee5cdep154eb3jsn8d818903b1df")
	req.Header.Add("X-RapidAPI-Host", "tripadvisor16.p.rapidapi.com")

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Ctx.WriteString("Error making request")
		return
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		c.Ctx.WriteString(fmt.Sprintf("Error: Unexpected status code %d", res.StatusCode))
		return
	}

	// Initialize the response structure
	var response HotelSearchResponse

	// Decode the JSON response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error parsing JSON response: %v", err))
		return
	}

	// Check if the response status is true
	// Check if the response status is true
	if !response.Status {
		messages, ok := response.Message.([]interface{})
		if ok {
			for _, msg := range messages {
				if s, ok := msg.(string); ok {
					// Handle each message
					fmt.Println("Message:", s)
				}
			}
			c.Ctx.WriteString("API response status is false with multiple messages")
		} else if msg, ok := response.Message.(string); ok {
			// Handle the single string message
			c.Ctx.WriteString(fmt.Sprintf("API response status is false: %s", msg))
		} else {
			c.Ctx.WriteString("API response status is false with an unexpected message type")
		}
		return
	}

	// Pass the hotel search results to the view template
	c.Data["Results"] = response.Data

	// Render the hotel search results view (hotel_search.tpl)
	c.TplName = "hotel_search.tpl"
}
