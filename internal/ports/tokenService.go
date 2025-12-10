package ports

type TokenService interface {
	Generate(username string) (string, error)
	Validate(token string) (string, error)
}