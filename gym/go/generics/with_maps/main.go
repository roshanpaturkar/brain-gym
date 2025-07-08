package main

type CustomMap[T comparable, V int | string] map[T]V

func main() {
	// Example usage of CustomMap with int keys and string values
	intStringMap := CustomMap[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	// Example usage of CustomMap with string keys and int values
	stringIntMap := CustomMap[string, int]{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	println("intStringMap:", intStringMap[1], intStringMap[2], intStringMap[3])
	println("stringIntMap:", stringIntMap["one"], stringIntMap["two"], stringIntMap["three"])
}