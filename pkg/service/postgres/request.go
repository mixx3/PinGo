package postgres

import (
	api "PinGo/pkg/api"
)

type requestService struct {
	repo api.RequestRepository
}

func NewRequestService(repo api.RequestRepository) api.RequestService {
	return &requestService{repo: repo}
}

func (s *requestService) Create(schema *api.RequestPostSchema) error {
	return s.repo.Add(schema)
}

func (s *requestService) GetAll() ([]*api.RequestGetSchema, error) {
	return s.repo.GetAll()
}

func (s *requestService) Get(id int) (*api.RequestGetSchema, error) {
	return s.repo.Get(id)
}

func (s *requestService) Delete(id int) error {
	return s.repo.Delete(id)
}
