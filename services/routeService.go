package services

import (
	"encoding/json"
	"fmt"
	"users/models"
	"users/utils"

	"github.com/go-resty/resty/v2"
)

type NominatimResponse []struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

func GetCoordinates(locationName string) (string, string, error) {
	client := resty.New()
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", locationName)

	resp, err := client.R().
		SetHeader("User-Agent", "YourAppName").
		SetHeader("Accept", "application/json").
		Get(url)

	if err != nil {
		return "", "", err
	}

	var nominatimResp NominatimResponse
	err = json.Unmarshal(resp.Body(), &nominatimResp)
	if err != nil {
		return "", "", err
	}

	if len(nominatimResp) == 0 {
		return "", "", fmt.Errorf("no results found for location: %s", locationName)
	}

	return nominatimResp[0].Lat, nominatimResp[0].Lon, nil
}

func GetRoutes(email string, fromLocation string, toLocation string, price string) ([]models.Route, error) {
	var routes []models.Route

	// Fetch latitude and longitude for fromLocation
	fromLat, fromLon, err := GetCoordinates(fromLocation)
	if err != nil {
		return routes, err
	}

	// Fetch latitude and longitude for toLocation
	toLat, toLon, err := GetCoordinates(toLocation)
	if err != nil {
		return routes, err
	}

	// Now that you have coordinates, you can calculate distance or just store these for routes
	query := "SELECT id, route_name, total_time, total_price, distance FROM routes WHERE total_price = ?"
	rows, err := utils.GetDB().Query(query, price)
	if err != nil {
		return routes, err
	}
	defer rows.Close()

	for rows.Next() {
		var route models.Route
		if err := rows.Scan(&route.ID, &route.RouteName, &route.TotalTime, &route.TotalPrice, &route.Distance); err != nil {
			return routes, err
		}
		route.FromLat = fromLat
		route.FromLon = fromLon
		route.ToLat = toLat
		route.ToLon = toLon
		routes = append(routes, route)
	}

	return routes, nil
}

func GetRouteDetails(routeID string) ([]models.Step, error) {
	var steps []models.Step
	query := "SELECT id, route_id, from_station_name, transportation_id, transportation_name, next_destination_time, next_destination_distance, to_station_name, price, status_item FROM steps WHERE route_id = ?"
	rows, err := utils.GetDB().Query(query, routeID)
	if err != nil {
		return steps, err
	}
	defer rows.Close()

	for rows.Next() {
		var step models.Step
		if err := rows.Scan(&step.ID, &step.RouteID, &step.FromStationName, &step.TransportationID, &step.TransportationName, &step.NextDestinationTime, &step.NextDestinationDistance, &step.ToStationName, &step.Price, &step.StatusItem); err != nil {
			return steps, err
		}
		steps = append(steps, step)
	}

	return steps, nil
}
