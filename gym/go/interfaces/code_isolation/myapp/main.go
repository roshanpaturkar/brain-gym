package main

import (
	"myapp/user/user"
)

func main() {
    dbRepo := &user.MockUserRepository{}
    service := user.NewUserService(dbRepo)

    err := service.RegisterUser("Roshan")
    if err != nil {
        panic(err)
    }
}
