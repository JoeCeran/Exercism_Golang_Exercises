//Package weather provides tools for 
//current weather conditions. 
package weather

//CurrentCondition represents current weather condition.
var CurrentCondition string
//CurrentLocation represents current location of user.
var CurrentLocation string

//Forecast returns the weather conditions of the city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
