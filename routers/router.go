package routers

import (
	"travelguide/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/searchLocation", &controllers.LocationSearchController{})
	// beego.Router("/hotel_search", &controllers.HotelSearchController{})
	beego.Router("/searchHotels", &controllers.HotelSearchController{})
	beego.Router("/searchRestaurants", &controllers.RestaurantSearchController{})

}
