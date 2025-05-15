package main

import "fmt"

type Rectangle struct {
	length int
	bridth int
}

type Circle struct {
	radius float32
}

func (rect Rectangle) Area() int {
	return rect.bridth * rect.length
}

func (rect *Rectangle) Scale(fact int) {
	rect.bridth *= fact
	rect.length *= fact
}

func (cir Circle) Area() float32 {
	return cir.radius * cir.radius
}

func main() {
	rect := Rectangle{
		4,
		6,
	}
	
	cir := Circle{
		6,
	}
	fmt.Println(rect.Area())
	fmt.Println(cir.Area())
	rect.Scale(2)
	fmt.Println("After scaling:", rect.Area())
}