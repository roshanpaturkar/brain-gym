package main

import "fmt"

type ParkingSystem struct {
    cars [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{cars: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
    this.cars[carType - 1] -= 1

    return this.cars[carType - 1] >= 0
}

func main() {
	parkingSystem := Constructor(1, 1, 0)
	fmt.Println("Parking System: ", parkingSystem.AddCar(1))
	fmt.Println("Parking System: ", parkingSystem.AddCar(2))
	fmt.Println("Parking System: ", parkingSystem.AddCar(3))
	fmt.Println("Parking System: ", parkingSystem.AddCar(1))
}