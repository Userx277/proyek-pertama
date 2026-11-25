package main

import "fmt"

func main() {
	var nilai [4]int

	nilai[0] = 88
	nilai[1] = 76

	suhu := [3]float64{24.5, 25, 23.8}
	hari := [...]string{"SEN", "SEL", "RAB"}

	fmt.Println(nilai)
	fmt.Println(suhu[1])
	fmt.Println((hari))
}
