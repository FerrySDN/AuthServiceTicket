package auth

import (
	"errors"

	"github.com/FerrySDN/auth-service/internal/ports"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Service struct {
	Repo  ports.UserRepository
	Token ports.TokenService
}

func NewService(r ports.UserRepository, t ports.TokenService) *Service {
	return &Service{Repo: r, Token: t}
}

func (s *Service) Register(username, password string) error {
	if len(username) < 3 || len(password) < 6 {
		return errors.New("username or password too short")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.Repo.Create(username, string(hash))
}

func (s *Service) Login(username, password string) (string, error) {
	u, err := s.Repo.FindByUsername(username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return "", ErrInvalidCredentials
	}

	return s.Token.Generate(u.Username)
}
