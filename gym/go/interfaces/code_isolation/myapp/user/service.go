package user

type UserService struct {
    repo UserRepository // Use the interface, not the concrete type
}

func NewUserService(r UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) RegisterUser(name string) error {
    user := &User{Name: name}
    return s.repo.SaveUser(user)
}
