// Package weather can forecast the current weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition represents the weather condition given a location.
	CurrentCondition string	
    // CurrentLocation represents the location.
	CurrentLocation  string 
)

// Forecast formats the weather report into a readable string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
