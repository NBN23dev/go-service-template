package services

type Service struct {
}

type Repositories struct {
}

// NewService
func NewService(repo Repositories) (*Service, error) {
	return &Service{}, nil
}
