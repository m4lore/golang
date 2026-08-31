package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate/100) * float64(productionRate) 
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    const min int8 = 60
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / int(min)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const (
        group int = 10
        carsGroupPrice int = 95000
        carPrice int = 10000
    )

    var  (
        carsGroups int = carsCount / group
        carsLeft int = carsCount % group
    )

    return uint((carsGroups * carsGroupPrice) + (carsLeft * carPrice))
}
