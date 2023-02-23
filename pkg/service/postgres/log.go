package postgres

import "PinGo/pkg/api"

type logService struct {
	repo api.LogRepository
}

func NewLogService(repo api.LogRepository) api.LogService {
	return &logService{repo: repo}
}

func (s *logService) Create(schema *api.LogPostSchema) error {
	return s.repo.Add(schema)
}

func (s *logService) GetAll() ([]*api.LogGetSchema, error) {
	return s.repo.GetAll()
}

func (s *logService) Get(id int) (*api.LogGetSchema, error) {
	return s.repo.Get(id)
}

func (s *logService) Delete(id int) error {
	return s.repo.Delete(id)
}
