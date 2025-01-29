package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"online-shop-backend/internal/domain"
)

const UsersTable = "users"

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) SignUp(person domain.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(username, password, role) VALUES($1, $2, $3) RETURNING id", UsersTable)

	row := r.db.QueryRow(query, person.Username, person.Password, person.Role)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
