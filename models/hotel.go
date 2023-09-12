package models

type HotelInfo struct {
	CurrentPage int `json:"currentPage"`
	Data        []struct {
		IDDetail          string `json:"idDetail"`
		BasicPropertyData struct {
			Photos struct {
				Main struct {
					Urls struct {
						ImgURL string `json:"relativeUrl"`
					} `json:"highResJpegUrl"`
				} `json:"main"`
			} `json:"photos"`
		} `json:"basicPropertyData"`
		DisplayName struct {
			Title string `json:"text"`
		} `json:"displayName"`
		PriceDisplayInfoIrene struct {
			DisplayPrice struct {
				AmountPerStay struct {
					AmountRounded string `json:"amountRounded"`
				} `json:"amountPerStay"`
			} `json:"displayPrice"`
		} `json:"priceDisplayInfoIrene"`
	} `json:"data"`
}
