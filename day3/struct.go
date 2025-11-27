package day3

import "fmt"

func Struct() {
	type car struct {
		Brand string
		Year  int
		Price int
	}
	Nexon := car{Brand: "TATA", Year: 2025, Price: 1000000}
	Thar := car{Brand: "Mahindra", Year: 2025, Price: 2480000}
	Sonet := car{Brand: "Kia", Year: 2025, Price: 2635789}
	fmt.Println("showrooms car models :", Nexon)
	fmt.Println("showrooms car models :", Thar)
	fmt.Println("showrooms car models :", Sonet)

}
