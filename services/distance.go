package services

import (
	"errors"
	"math"

	"github.com/besanh/soa/models"
)

type (
	IDistance interface {
		CalculateDistance(ip, city string) (models.DistanceResult, error)
	}

	Distance struct{}
)

var DistanceService IDistance

func NewDistance() IDistance {
	return &Distance{}
}

func (s *Distance) CalculateDistance(ip, city string) (models.DistanceResult, error) {
	ipCoord, err := getCoordinatesForIP(ip)
	if err != nil {
		return models.DistanceResult{}, err
	}

	cityCoord, err := getCoordinatesForCity(city)
	if err != nil {
		return models.DistanceResult{}, err
	}

	distance := haversineDistance(ipCoord, cityCoord)
	return models.DistanceResult{
		IP:        ip,
		City:      city,
		Distance:  distance,
		Unit:      "km",
		IPCoord:   ipCoord,
		CityCoord: cityCoord,
	}, nil
}

// getCoordinatesForIP simulates an IP geolocation lookup.
// In production, integrate with a real geolocation service.
func getCoordinatesForIP(ip string) (models.Coordinates, error) {
	// Dummy data: if ip equals "8.8.8.8", return Mountain View, CA.
	if ip == "8.8.8.8" {
		return models.Coordinates{Lat: 37.3861, Lon: -122.0839}, nil
	}
	// Otherwise, default to Ho Chi Minh City.
	return models.Coordinates{Lat: 10.8231, Lon: 106.6297}, nil
}

// getCoordinatesForCity simulates a city geocoding lookup.
// In production, use a geocoding API or database.
func getCoordinatesForCity(city string) (models.Coordinates, error) {
	switch city {
	case "Ho Chi Minh City":
		return models.Coordinates{Lat: 10.8231, Lon: 106.6297}, nil
	case "Paris":
		return models.Coordinates{Lat: 48.8566, Lon: 2.3522}, nil
	case "London":
		return models.Coordinates{Lat: 51.5074, Lon: -0.1278}, nil
	case "New York":
		return models.Coordinates{Lat: 40.7128, Lon: -74.0060}, nil
	default:
		return models.Coordinates{}, errors.New("unknown city")
	}
}

// haversineDistance calculates the distance in kilometers between two points.
func haversineDistance(c1, c2 models.Coordinates) float64 {
	const earthRadiusKm = 6371.0

	lat1 := c1.Lat * math.Pi / 180
	lon1 := c1.Lon * math.Pi / 180
	lat2 := c2.Lat * math.Pi / 180
	lon2 := c2.Lon * math.Pi / 180

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(dlon/2)*math.Sin(dlon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusKm * c
}
