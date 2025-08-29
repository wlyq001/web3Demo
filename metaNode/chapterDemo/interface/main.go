package main

import (
	"fmt"
	"strconv"
)

type Man interface {
	pee(int) string
	person
}
type person interface {
	eat() bool
}

type Dad struct {
	age int
}

func (d *Dad) eat() bool {
	if d.age > 20 {
		return true
	}
	return false
}

func (d *Dad) pee(time int) string {
	time = d.age
	return strconv.Itoa(time + 1)
}
func main() {
	dd := Dad{age: 25}
	pd := person(&dd)
	md := Man(&dd)

	fmt.Println(pd.eat())
	fmt.Println(md.pee(1))
	fmt.Println(md.eat())

}
