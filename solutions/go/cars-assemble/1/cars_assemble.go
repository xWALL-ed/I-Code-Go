package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 { return (successRate/100.00)*float64(productionRate) }

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int { return int((successRate/100.00)*float64(productionRate/60)) }

func CalculateCost(carsCount int) uint {
	if carsCount < 10 { return uint(carsCount*10000) }
    return uint(((carsCount/10)*95000) + ((carsCount%10)*10000))
}
