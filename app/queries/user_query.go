package queries

import (
	"github.com/google/uuid"
	"github.com/ineoo/go-planigramme/app/models"
	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	users := []models.User{}
	query := `SELECT * FROM users`

	if err := q.Select(&users, query); err != nil {
		return users, err
	}
	return users, nil
}

func (q *UserQueries) GetUserByID(id uuid.UUID) (models.User, error) {
	user := models.User{}
	query := `SELECT * FROM users WHERE id=$1`

	err := q.Get(&user, query, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (q *UserQueries) CreateUser(u models.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING *`

	return q.Get(&u, query, u.FirstName, u.LastName, u.Email, u.Password)
}

func (q *UserQueries) UpdateUser(u models.User) error {
	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, password_hash=$4, updated_at=NOW() WHERE id=$5 RETURNING *`

	return q.Get(&u, query, u.FirstName, u.LastName, u.Email, u.Password, u.ID)
}

func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
