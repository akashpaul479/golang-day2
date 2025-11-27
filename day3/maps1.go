package day3

import "fmt"

func Map2() {
	countrys := make(map[string]string)
	countrys["India"] = "Newdelhi"
	countrys["Germany"] = "Berlin"
	countrys["Japan"] = "Tokyo"
	countrys["France"] = "Paris"
	fmt.Println("India's capital is:", countrys["India"])
	fmt.Println("Germany's capital is:", countrys["Germany"])
	fmt.Println("Japan's capital is:", countrys["Japan"])
	fmt.Println("France's capital is:", countrys["France"])
}
