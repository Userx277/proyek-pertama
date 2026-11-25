package main

import "fmt"

func main() {
	angka := make([]int, 0, 3)
	angka = append(angka, 10, 20, 30)
	fmt.Println(angka, len(angka), cap(angka))

	angka = append(angka, 40, 70, 80)
	fmt.Println(angka, len(angka), cap(angka))

	angka = append(angka, 50, 60, 90)
	fmt.Println(angka, len(angka), cap(angka))
}
