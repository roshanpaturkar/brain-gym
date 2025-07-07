package main

type Num interface {
	int | int16 | int32 | float64 | float32
}

func Add[T Num](a, b T) T {
	return a + b
}

func main() {
	intResult := Add(3, 4)
	floatResult := Add(3.5, 4.5)
	println("Integer addition result:", intResult)
	println("Float addition result:", floatResult)

	genericResultInt := Add(3, 4)
	genericResultFloat := Add(3.5, 4.5)
	println("Generic integer addition result:", genericResultInt)
	println("Generic float addition result:", genericResultFloat)
}