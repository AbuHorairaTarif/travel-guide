package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type RestaurantSearchController struct {
	web.Controller
}

type RestaurantSearchResponse struct {
	Status    bool                     `json:"status"`
	Message   string                   `json:"message"`
	Timestamp int64                    `json:"timestamp"`
	Data      []RestaurantSearchResult `json:"data"`
}

type RestaurantSearchResult struct {
	Title         string `json:"title"`
	LocationID    string `json:"locationId"`
	DocumentID    string `json:"documentId"`
	TrackingItems string `json:"trackingItems"`
	SecondaryText string `json:"secondaryText"`
}

func (c *RestaurantSearchController) Get() {
	// Get the query parameter "query" from the URL
	query := c.GetString("query")

	// Define the TripAdvisor API URL for restaurant location search
	url := "https://tripadvisor16.p.rapidapi.com/api/v1/restaurant/searchLocation?query=" + query

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
		c.Ctx.WriteString("Error: Unexpected status code")
		return
	}

	// Initialize the response structure
	var response RestaurantSearchResponse

	// Decode the JSON response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		c.Ctx.WriteString("Error parsing JSON response")
		return
	}

	// Check if the response status is true
	if !response.Status || len(response.Data) == 0 {
		c.Ctx.WriteString("Error: Location not found")
		return
	}

	// Retrieve the locationId for the first restaurant location result
	locationID := response.Data[0].LocationID

	// Use the locationId to perform the actual restaurant search
	performRestaurantSearch(locationID, c)
}

func performRestaurantSearch(locationID string, c *RestaurantSearchController) {
	// Define the TripAdvisor API URL for restaurant search using the provided locationId
	url := fmt.Sprintf("https://tripadvisor16.p.rapidapi.com/api/v1/restaurant/searchRestaurants?locationId=%s", locationID)

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
	var response RestaurantSearchResponse

	// Decode the JSON response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		c.Ctx.WriteString("Error parsing JSON response")
		return
	}

	// Check if the response status is true
	if !response.Status {
		c.Ctx.WriteString("Error: API response status is false")
		return
	}

	// Pass the restaurant search results to the view template
	c.Data["Results"] = response.Data

	// Render the restaurant search results view (restaurant_search.tpl)
	c.TplName = "restaurant_search.tpl"
}
