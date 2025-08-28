package main

import "fmt"

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func main() {
	fmt.Println(January)
	fmt.Println(February)
	fmt.Println(March)
	fmt.Println(April)
	fmt.Println(May)
	fmt.Println(June)
	fmt.Println(July)
	fmt.Println(August)
	fmt.Println(September)
	fmt.Println(October)
	fmt.Println(November)
	fmt.Println(December)
}
