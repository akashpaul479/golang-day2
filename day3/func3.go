package day3

import "fmt"

func Sumandaverage(numbers []int) (int, float64) {
	if len(numbers) == 0 {
		return 0, 0
	}
	sum := 0
	for _, m := range numbers {
		sum += m
	}
	average := float64(sum) / float64(len(numbers))
	return sum, average
}
func Sumnaverages() {
	nums := []int{10, 20, 30, 40, 50}

	total, avg := Sumandaverage(nums)

	fmt.Println("sum:", total)
	fmt.Println("avg", avg)
}
