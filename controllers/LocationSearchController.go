package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type LocationSearchController struct {
	web.Controller
}

type LocationSearchResponse struct {
	Status    bool                   `json:"status"`
	Message   string                 `json:"message"`
	Timestamp int64                  `json:"timestamp"`
	Data      []LocationSearchResult `json:"data"`
}

type LocationSearchResult struct {
	Title         string `json:"title"`
	GeoID         string `json:"geoId"`
	DocumentID    string `json:"documentId"`
	TrackingItems string `json:"trackingItems"`
	SecondaryText string `json:"secondaryText"`
}

func (c *LocationSearchController) Get() {
	// Get the query parameter "query" from the URL
	query := c.GetString("query")

	// Define the TripAdvisor API URL for location search
	url := "https://tripadvisor16.p.rapidapi.com/api/v1/hotels/searchLocation?query=" + query

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
	var response LocationSearchResponse

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

	// Pass the location search results to the view template
	c.Data["Results"] = response.Data

	// Render the location search results view
	c.TplName = "location_search.tpl"
}
