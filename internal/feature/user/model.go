package user

type user struct {
	Id       int
	Email    string
	Password string
	Photo    string
}

type userInterface interface {
	All() ([]user, error)
	Find(id string) (*user, error)
	Create(email, password, photo string) (*user, error)
	Update(email, password, photo string) (*user, error)
	Delete(id int) error
}
