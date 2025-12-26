package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	properCarCount := successRate / 100 * float64(productionRate)
    return properCarCount
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	properCarCount := int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
    return properCarCount
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groupedCarsCount := carsCount / 10
    individualCars := carsCount % 10
    
    productionCost := (groupedCarsCount * 95000) + (individualCars * 10000)
    return uint(productionCost)
}
