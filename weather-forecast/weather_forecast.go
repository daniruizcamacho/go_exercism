// Package weather provides methods to get the current status.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast returns a string with the current weahter condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
