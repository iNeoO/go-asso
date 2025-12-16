package session

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByToken(token string) (*Session, error) {
	return s.repo.GetByToken(token)
}

func (s *Service) Create(session *Session) (*Session, error) {
	return s.repo.Create(session)
}

func (s *Service) Update(session *Session) (*Session, error) {
	return s.repo.Update(session)
}

func (s *Service) DeleteByToken(token string) error {
	return s.repo.DeleteByToken(token)
}