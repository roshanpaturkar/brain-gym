package main

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	intResult := AddInt(3, 4)
	floatResult := AddFloat(3.5, 4.5)
	println("Integer addition result:", intResult)
	println("Float addition result:", floatResult)

	genericResultInt := Add(3, 4)
	genericResultFloat := Add(3.5, 4.5)
	println("Generic integer addition result:", genericResultInt)
	println("Generic float addition result:", genericResultFloat)
}

