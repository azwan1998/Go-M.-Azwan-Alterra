package appModel

type UserModel interface {
	GetByEmailAndPassword(email string, password string) (User, error)
	GetAll() ([]User, error)
	Add(User) (User, error)
}
