package main

type CustomData interface {
	any | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func main () {
	user1 := User[int]{ID: 1, Name: "Alice", Data: 42}
	user2 := User[string]{ID: 2, Name: "Bob", Data: "Hello"}

	println("User 1:", user1.ID, user1.Name, user1.Data)
	println("User 2:", user2.ID, user2.Name, user2.Data)
}