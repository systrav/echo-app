package services

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Hello() string {
	return "Hello User"
	// fmt.Println("Hello User")
}
