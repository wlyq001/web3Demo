package main

import "fmt"

var (
	str   string = "Hello World"
	x     int    = 77
	mp    map[int]string
	bytes []byte
	p     *int
)

func main() {
	fmt.Println(str, x)
	fmt.Println(p, mp)
	fmt.Println(bytes)
	k := 1
	fmt.Println(k)

	var a, b, c rune = 1, 3, '往'
	fmt.Println(a, b, c)
}
