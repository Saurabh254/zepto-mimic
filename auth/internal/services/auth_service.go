package services

import (
	"context"
	"errors"

	"github.com/saurabh254/zepto-mimic/auth/internal/core/auth"
	"github.com/saurabh254/zepto-mimic/auth/internal/core/database"
	"github.com/saurabh254/zepto-mimic/auth/internal/repositories"
)

func LoginUser(ctx context.Context, email, password string, db database.Database) (string, error) {
	user_repo := repositories.NewUserRepository(db, ctx)
	user, err := user_repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Verify password (this is a placeholder, implement proper password hashing and verification)
	if user.PasswordHash != password {
		return "", errors.New("invalid credentials")
	}
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
