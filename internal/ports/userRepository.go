package ports

import "github.com/FerrySDN/auth-service/internal/core/domain"

type UserRepository interface {
	Create(username, passwordHash string) error
	FindByUsername(username string) (*domain.User, error)
}
