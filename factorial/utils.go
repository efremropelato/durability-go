package factorial

import (
	"math"
	"sort"
)

func Max(nums []float64) float64 {
	if len(nums) == 0 {
		return math.NaN() // Restituisce NaN se la slice è vuota
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func Min(values []float64) float64 {
	if len(values) == 0 {
		return math.Inf(1) // Ritorna un valore positivo infinito se la slice è vuota
	}

	minValue := values[0]
	for _, v := range values[1:] {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
func Skewness(data []float64) float64 {
	n, sum, sumSq, sumCube := float64(len(data)), 0.0, 0.0, 0.0
	for _, x := range data {
		sum += x
	}
	mean := sum / n
	for _, x := range data {
		dev := x - mean
		sumSq += dev * dev
		sumCube += dev * dev * dev
	}
	variance := sumSq / n
	stdDev := math.Sqrt(variance)
	return (n / ((n - 1) * (n - 2))) * (sumCube / (stdDev * stdDev * stdDev))
}
func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0 // gestire la slice vuota
	}

	// Ordina la slice
	sort.Float64s(data)

	// Controlla se n è pari o dispari e calcola la mediana
	if n%2 == 1 {
		return data[n/2]
	}
	return (data[n/2-1] + data[n/2]) / 2
}
func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	var sum float64
	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}
func Variance(data []float64) float64 {
	n := float64(len(data))
	if n == 0 {
		return 0 // Evita la divisione per zero
	}

	// Calcola la media
	var sum float64
	for _, value := range data {
		sum += value
	}
	mean := sum / n

	// Calcola la somma dei quadrati delle differenze dalla media
	var sumOfSquares float64
	for _, value := range data {
		sumOfSquares += (value - mean) * (value - mean)
	}

	// Calcola la varianza
	return sumOfSquares / n
}
func StdDev(data []float64) float64 {
	if len(data) == 0 {
		return 0 // Evita la divisione per zero
	}

	// Calcola la media
	var sum float64
	for _, value := range data {
		sum += value
	}
	mean := sum / float64(len(data))

	// Calcola la varianza
	var varianceSum float64
	for _, value := range data {
		varianceSum += (value - mean) * (value - mean)
	}
	variance := varianceSum / float64(len(data))

	// Restituisce la radice quadrata della varianza come deviazione standard
	return math.Sqrt(variance)
}
