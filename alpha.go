package main

import "fmt"

func main() {
	a := 25
	b := 4
	c := 2.5
	p := a / b
	q := a % b
	r := float64(a / b)
	s := float64(a) / float64(b)
	u := int(c*float64(b)) + a%3
	t := (a%2 == 1) && (float64(b) > c)
	fmt.Println("Nama : Made Kevin Arta Wijaya NIM : 42630015")
	fmt.Println(p, q, r, s, u, t)
}
