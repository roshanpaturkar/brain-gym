package user

import "fmt"

type DBUserRepository struct {
    // db connection or mock
}

func (repo *DBUserRepository) GetUserByID(id int) (*User, error) {
    fmt.Println("Fetching user from DB...")
    // Simulate DB logic
    return &User{ID: id, Name: "Roshan"}, nil
}

func (repo *DBUserRepository) SaveUser(user *User) error {
    fmt.Printf("Saving user %v to DB...\n", user.Name)
    return nil
}
