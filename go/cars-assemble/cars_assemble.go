// Package cars implements a solution to the cars-assemble exercise
package cars

import "errors"

// verifySuccessRate calculates if the successRate is within the
// acceptable range.
func verifySuccessRate(successRate float64) error {
	if successRate < 0 || successRate > 100 {
		return errors.New("successRate must be between 0 and 100")
	}
	return nil
}


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if err := verifySuccessRate(successRate); err != nil {
		// Handle the error. For example, return 0 if successRate is invalid.
		return 0
	}
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	group := carsCount / 10
	individual := carsCount % 10

	return uint(group * 95000 + individual * 10000)
}