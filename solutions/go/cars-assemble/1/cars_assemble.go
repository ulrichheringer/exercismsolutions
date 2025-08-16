package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * float64((successRate / 100))
}
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}
func CalculateCost(carsCount int) uint {
	remainder := carsCount % 10
	main := (carsCount - remainder) / 10
	return uint(main*95000 + remainder*10000)
}
