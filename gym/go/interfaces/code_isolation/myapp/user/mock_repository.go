package user

import "fmt"

type MockUserRepository struct{}

func (m *MockUserRepository) GetUserByID(id int) (*User, error) {
    return &User{ID: id, Name: "MockUser"}, nil
}

func (m *MockUserRepository) SaveUser(user *User) error {
    fmt.Println("Mock save:", user.Name)
    return nil
}
