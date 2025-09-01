package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {
	fmt.Println("Rectangle area is :", r)
}

func (r Rectangle) Perimeter() {
	fmt.Println("Rectangle perimeter area is :", r)
}

type Circle struct {
}

func (r Circle) Area() {
	fmt.Println("Circle area area is :", r)
}

func (r Circle) Perimeter() {
	fmt.Println("Circle perimeter area area is :", r)
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (ob *Employee) PrintInfo() {
	fmt.Println("Employee Print Info :", ob)
}

func main() {
	// task1:
	//rectangle := Rectangle{}
	//circle := Circle{}
	//recShape := Shape(rectangle)
	//cirShape := Shape(circle)
	//
	//recShape.Area()
	//recShape.Perimeter()
	//cirShape.Area()
	//cirShape.Perimeter()

	// task2:
	emp := Employee{
		Person: Person{
			Name: "wlyq",
			Age:  27,
		},
		EmployeeID: 7,
	}
	emp.PrintInfo()
}
