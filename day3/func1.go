package day3

import "fmt"

func Largestnumbers(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, m := range numbers {
		if m > max {
			max = m
		}
	}
	return max
}
func Largestnumber() {
	nums := []int{20, 30, 40, 50, 60, 70, 100}
	fmt.Println("Largest number is:", Largestnumbers(nums))
}
