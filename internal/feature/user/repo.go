package user

import "github.com/fajaralfa/askme/internal/db"

type repo struct{}

func (repo) All() ([]user, error) {
	var users []user
	rows, err := db.Connect().Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user
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

func (repo) Find(id string) (*user, error) {
	var user *user = &user{}

	rows, err := db.Connect().Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Photo); err != nil {
			return nil, err
		}
	}

	return user, nil
}
func (repo) Create(email, password, photo string) (*user, error) {
	return &user{}, nil
}
func (repo) Update(email, password, photo string) (*user, error) {
	return &user{}, nil
}
func (repo) Delete(id int) error {
	return nil
}
