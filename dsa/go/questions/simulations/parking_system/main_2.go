package main

import "fmt"

type ParkingSystem struct {
    big, small, medium int
    bigAvailable, smallAvailable, mediumAvailable int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        big, small, medium,
        big, small, medium,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    switch carType {
        case 1:
            if this.bigAvailable > 0 {
                this.bigAvailable--
                return true
            }
        case 2:
            if this.mediumAvailable > 0 {
                this.mediumAvailable--
                return true
            }
        case 3:
            if this.smallAvailable > 0 {
                this.smallAvailable--
                return true
            }
    }
    return false
}

func main() {
	parkingSystem := Constructor(1, 1, 0)
	fmt.Println("Parking System: ", parkingSystem.AddCar(1))
	fmt.Println("Parking System: ", parkingSystem.AddCar(2))
	fmt.Println("Parking System: ", parkingSystem.AddCar(3))
	fmt.Println("Parking System: ", parkingSystem.AddCar(1))
}