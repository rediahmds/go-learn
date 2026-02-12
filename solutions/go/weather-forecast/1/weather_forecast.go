// Package weather provides tools to forecast current weather in Goblinocus country.
package weather


var (
    // CurrentCondition represents current weather condition.
	CurrentCondition string
    // CurrentLocation represents the location which the weather will be forecasted. 
	CurrentLocation  string
)

// Forecast returns a string that gives you forecast result in specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
