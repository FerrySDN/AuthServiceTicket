package ports

type TokenService interface {
	Generate(username string, UserId int64) (string, error)
	Validate(token string) (string, error)
}