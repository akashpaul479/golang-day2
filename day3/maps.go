package day3

import "fmt"

func Maps() {
	grades := make(map[string]int)
	grades["akash"] = 57
	grades["abhi"] = 58
	grades["muzz"] = 60
	grades["vishal"] = 52
	fmt.Println("marks of akash:", grades["akash"])
	fmt.Println("marks of abhi:", grades["abhi"])
	fmt.Println("marks of muzz:", grades["muzz"])
	fmt.Println("marks of vishal:", grades["vishal"])
	grades["bunty"] = 48
	grades["biswas"] = 60
	grades["shiva"] = 51
	fmt.Println("marks of bunty:", grades["bunty"])
	fmt.Println("marks of biswas:", grades["biswas"])
	fmt.Println("marks of shiva:", grades["shiva"])
	total := 0
	count := 0
	for _, m := range grades {
		total += m
		count++
	}
	average := total / count

	fmt.Println("total marks:", total)
	fmt.Println("average marks:", average)
}
