// Package weather provides the forecast for the city provided.
package weather

// CurrentCondition stores the condition that was passed in the function.
var CurrentCondition string

// CurrentLocation stores the location that was passed in the function.
var CurrentLocation string

// Forecast returns the weather condition of a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
