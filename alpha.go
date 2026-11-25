package main

import "fmt"

func main() {
	nilai := []int{80, 75, 90}
	nilai = append(nilai, 85, 95)
	nilai[0] = 100
	for i := 0; i < len(nilai); i++ {
		fmt.Println(1, nilai[i])
	}
	for i, v := range nilai {
		fmt.Println(i, v)
	}
	for _, v := range nilai {
		fmt.Println(v)
	}
}
