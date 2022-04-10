package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	noOfCarProduced := float64((successRate / 100) * float64(productionRate))
	return noOfCarProduced
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsProducedPerMin := int(((successRate / 100) * float64(productionRate)) / 60)
	return carsProducedPerMin
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	individual := carsCount % 10
	group := int(carsCount / 10)
	cost := uint((group * 95000) + (individual * 10000))
	return cost
}
