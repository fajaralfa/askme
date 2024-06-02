package repo

import (
	"database/sql"

	"github.com/fajaralfa/askme/internal/model"
)

type UserInterface interface {
	FindAll() ([]model.User, error)
	Find(id string) (*model.User, error)
	Create(email, password, photo string) (*model.User, error)
	Update(email, password, photo string) (*model.User, error)
	Delete(id int) error
	FindByEmail(email string) (*model.User, error)
}

type User struct {
	Db *sql.DB
}

func (r *User) FindAll() ([]model.User, error) {
	var users []model.User
	rows, err := r.Db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Photo); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *User) Find(id string) (*model.User, error) {
	var user *model.User

	rows, err := r.Db.Query("SELECT * FROM Users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user = new(model.User)
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Photo); err != nil {
			return nil, err
		}
	}

	return user, nil
}
func (r *User) Create(email, password, photo string) (*model.User, error) {
	var user *model.User

	stmt, err := r.Db.Prepare("INSERT INTO users (email, password, photo) VALUES (?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(email, password, photo)
	if err != nil {
		return nil, err
	}
	lastId, _ := res.LastInsertId()

	user = &model.User{Id: lastId, Email: email, Password: password, Photo: photo}

	return user, nil
}
func (r *User) Update(email, password, photo string) (*model.User, error) {
	return &model.User{}, nil
}
func (r *User) Delete(id int) error {
	return nil
}

func (r *User) FindByEmail(email string) (*model.User, error) {
	var user *model.User

	rows, err := r.Db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		user = new(model.User)
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Photo); err != nil {
			return nil, err
		}
	}

	return user, nil
}
