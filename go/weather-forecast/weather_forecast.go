// Package weather provides a weather forecast for a given city.
package weather

// CurrentCondition is the current weather condition for the city.
var CurrentCondition string
// CurrentLocation is the city for which the weather is being forecasted.
var CurrentLocation string

// Forecast returns the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}