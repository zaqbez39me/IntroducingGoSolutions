package main

func Min(nums []float64) (m float64) {
	for i, val := range nums {
		if i == 0 || val < m {
			m = val
		}
	}
	return
}

func Max(nums []float64) (m float64) {
	for i, val := range nums {
		if i == 0 || val > m {
			m = val
		}
	}
	return
}
