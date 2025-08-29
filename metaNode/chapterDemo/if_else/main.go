package main

import "fmt"

func main() {
	var a = 0
	if b := 2; a > 1 {
		b++
		fmt.Println(a)
	} else if c := 5; b > 1 {
		c++
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

	var d interface{}
	d = 1
	switch t := d.(type) {
	case int:
		fmt.Println("int", t)
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}

	var e = 1
	switch {
	case e == 1:
		fmt.Println("1")
	case e > 0:
		fmt.Println("2")

	}

}
