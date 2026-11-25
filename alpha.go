package main

import "fmt"

func main() {
	nilai := []float64{78.5, 82, 91.5, 64, 88}
	var total float64
	lulus := 0

	for _, v := range nilai {
		total += v
		if v >= 75 {
			lulus++
		}
	}
	rata := total / float64(len(nilai))
	fmt.Println(total, rata, lulus)
}
