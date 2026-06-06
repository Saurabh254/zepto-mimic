package repositories

import (
	"context"

	"github.com/saurabh254/zepto-mimic/auth/internal/core/database"
	"github.com/saurabh254/zepto-mimic/auth/internal/models"
)

type UserRepositoryImpl struct {
	db  database.Database
	ctx context.Context
}

func NewUserRepository(db database.Database, ctx context.Context) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db:  db,
		ctx: ctx,
	}
}

func (ar *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, hash_salt, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := ar.db.QueryRow(ar.ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ar *UserRepositoryImpl) UpdateUser(id int, email string) error {
	query := `
		UPDATE users
		SET email = $1
		WHERE id = $2
	`

	rows, err := ar.db.Query(ar.ctx, query, email, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (ar *UserRepositoryImpl) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	rows, err := ar.db.Query(ar.ctx, query, id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (ar *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	query := `
		SELECT id, email, password_hash, hash_salt, created_at, updated_at
		FROM users
	`

	rows, err := ar.db.Query(ar.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
