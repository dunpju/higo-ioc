package services

type DbService struct {
	DSN string
}

func NewDbService() *DbService {
	return &DbService{DSN:"dsn"}
}
