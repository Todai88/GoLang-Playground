package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}
