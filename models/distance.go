package models

// Coordinates holds latitude and longitude.
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// DistanceResult represents the response structure.
type DistanceResult struct {
	IP        string      `json:"ip"`
	City      string      `json:"city"`
	Distance  float64     `json:"distance"` // in km
	Unit      string      `json:"unit"`
	IPCoord   Coordinates `json:"ipCoord"`
	CityCoord Coordinates `json:"cityCoord"`
}
