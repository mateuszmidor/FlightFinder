package geo

import "math"

const earthRadiusKm = 6371

// GreatCircleDistance calculates flight distance between 2 locations, in kilometers.
func GreatCircleDistance(lat1, lat2 Latitude, lon1, lon2 Longitude) float32 {
	phi1 := float64(lat1) * math.Pi / 180
	phi2 := float64(lat2) * math.Pi / 180
	deltaPhi := float64(lat2-lat1) * math.Pi / 180
	deltaLambda := float64(lon2-lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) + math.Cos(phi1)*math.Cos(phi2)*math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return float32(earthRadiusKm * c)
}
