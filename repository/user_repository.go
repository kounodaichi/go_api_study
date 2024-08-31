package repository

import (
	"database/sql"
	"myapp/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUserByID(id int) (*model.User, error) {
	user := &model.User{}
	err := r.DB.QueryRow("select id, name, email from users where id = ? ", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
