package cars

var car_work_rate float64
var calc_rate int
var calc_rate_f float64
var calc_cost uint

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    car_work_rate = successRate * float64(productionRate) / float64(100)
    return car_work_rate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var car_work_rate_f float64 = successRate * float64(productionRate) / float64(100)
    calc_rate = int(car_work_rate_f) / 60
    return calc_rate
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    var car_amount uint = uint(carsCount) / 10 
    var car_remainder uint = uint(carsCount) % 10
    var calc_cost uint = (car_amount * uint(95000)) + (car_remainder * uint(10000))
    return calc_cost
}
