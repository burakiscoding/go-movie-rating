package stores

// Calculates the new average using the current average and number of values
func CalculateAverage(numberOfValues int, currentAvg float64, newNumber float64) float64 {
	total := float64(numberOfValues) * currentAvg
	return (total + newNumber) / (float64(numberOfValues) + 1)
}
