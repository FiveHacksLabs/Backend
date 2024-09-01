package models

type Route struct {
	ID        int     `json:"id"`
	RouteName string  `json:"routeName"`
	TotalTime string  `json:"totalTime"`
	TotalPrice float64 `json:"totalPrice"`
	Distance  float64 `json:"distance"`
	FromLat   string  `json:"fromLat"`
	FromLon   string  `json:"fromLon"`
	ToLat     string  `json:"toLat"`
	ToLon     string  `json:"toLon"`
    Price     float64 `json:"price"`
}


type Step struct {
	ID                   int     `json:"id"`
	RouteID              int     `json:"routeId"`
	FromStationName      string  `json:"fromStationName"`
	TransportationID     string  `json:"transportationId"`
	TransportationName   string  `json:"transportationName"`
	NextDestinationTime  string  `json:"nextDestinationTime"`
	NextDestinationDistance float64 `json:"nextDestinationDistance"`
	ToStationName        string  `json:"toStationName"`
	Price                float64 `json:"price"`
	StatusItem           string  `json:"statusItem"`
}
