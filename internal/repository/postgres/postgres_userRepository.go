package postgres

import (
	"database/sql"

	"github.com/FerrySDN/auth-service/internal/core/domain"
	"github.com/FerrySDN/auth-service/internal/ports"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(username, passwordHash string) error {
	_, err := r.db.Exec(`
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
	`, username, passwordHash)
	return err
}


func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	row := r.db.QueryRow(`
		SELECT id, username, password_hash, created_at
		FROM users WHERE username = $1
	`, username)

	var u domain.User
	if err := row.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
