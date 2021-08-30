package math

func Average(nums ...float64) float64 {
	var result float64 = 0
	for num := range nums {
		result += float64(num)
	}
	result = result / float64(len(nums))
	return float64(result)
}
