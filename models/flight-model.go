package models

type FlightData struct {
	Data struct {
		FilteredFlightsCount int `json:"filteredFlightsCount"`
		Flights              []struct {
			ID             string `json:"id"`
			ShareableURL   string `json:"shareableUrl"`
			TravelerPrices []struct {
				Price struct {
					Price struct {
						Currency struct {
							Code string `json:"code"`
						} `json:"currency"`
						Value int `json:"value"`
					} `json:"price"`
					VAT struct {
						Value int `json:"value"`
					} `json:"vat"`
				} `json:"price"`
			} `json:"travelerPrices"`
		} `json:"flights"`
		Routes []struct {
			Destination struct {
				CityCode string `json:"cityCode"`
				CityName string `json:"cityName"`
			} `json:"destination"`
			Origin struct {
				CityCode string `json:"cityCode"`
				CityName string `json:"cityName"`
			} `json:"origin"`
		} `json:"routes"`
	} `json:"data"`
}

// package models

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
// 	JsonChan    chan string
// 	ErrorChan   chan error
// }

// type ErrorMessage struct {
// 	Message    string
// 	SubMessage string
// 	Code       int
// }

// type PaymentMethodPrices struct {
// 	Typename string `json:"__typename"`
// 	Name     string `json:"name"`
// 	Price    Price  `json:"price"`
// 	Type     string `json:"type"`
// }

// type Recommendation struct {
// 	Typename string `json:"__typename"`
// 	Value    int    `json:"value"`
// }
// type PriceRange struct {
// 	Typename string `json:"__typename"`
// 	Max      int    `json:"max"`
// 	Min      int    `json:"min"`
// }
// type TravelTimeRange struct {
// 	Typename string `json:"__typename"`
// 	Max      int    `json:"max"`
// 	Min      int    `json:"min"`
// }
// type Destination struct {
// 	Typename      string `json:"__typename"`
// 	CityCode      string `json:"cityCode"`
// 	CityName      string `json:"cityName"`
// 	Code          string `json:"code"`
// 	ContinentCode string `json:"continentCode"`
// 	ContinentName string `json:"continentName"`
// 	CountryCode   string `json:"countryCode"`
// 	CountryName   string `json:"countryName"`
// 	Name          string `json:"name"`
// }
// type Origin struct {
// 	Typename      string `json:"__typename"`
// 	CityCode      string `json:"cityCode"`
// 	CityName      string `json:"cityName"`
// 	Code          string `json:"code"`
// 	ContinentCode string `json:"continentCode"`
// 	ContinentName string `json:"continentName"`
// 	CountryCode   string `json:"countryCode"`
// 	CountryName   string `json:"countryName"`
// 	Name          string `json:"name"`
// }
// type Routes struct {
// 	Typename           string      `json:"__typename"`
// 	DepartureAt        string      `json:"departureAt"`
// 	DepartureDate      string      `json:"departureDate"`
// 	DepartureTimeOfDay any         `json:"departureTimeOfDay"`
// 	Destination        Destination `json:"destination"`
// 	Origin             Origin      `json:"origin"`
// }
