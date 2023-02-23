package postgres

import (
	"PinGo/pkg/api"
)

type receiverService struct {
	repo api.ReceiverRepository
}

func NewReceiverService(repo api.ReceiverRepository) api.ReceiverService {
	return &receiverService{repo: repo}
}

func (s *receiverService) Create(schema *api.ReceiverPostSchema) error {
	return s.repo.Add(schema)
}
func (s *receiverService) GetAll() ([]*api.ReceiverGetSchema, error) {
	return s.repo.GetAll()
}

func (s *receiverService) Get(id int) (*api.ReceiverGetSchema, error) {
	return s.repo.Get(id)
}

func (s *receiverService) Delete(id int) error {
	return s.repo.Delete(id)
}
