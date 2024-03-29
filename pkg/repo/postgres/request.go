package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
	"gorm.io/gorm"
)

type requestPgRepository struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) api.RequestRepository {
	err := db.AutoMigrate(&models.Request{})
	if err != nil {
		return nil
	}
	return &requestPgRepository{db: db}
}

func (r *requestPgRepository) Add(schema *api.RequestPostSchema) error {
	request := models.Request{
		StatusExpected:         schema.StatusExpected,
		Body:                   schema.Body,
		Name:                   schema.Name,
		Address:                schema.Address,
		RepeatTimeMs:           schema.RepeatTimeMs,
		ExpectedResponseTimeMs: schema.ExpectedResponseTimeMs,
		ReceiverID:             schema.ReceiverID,
	}
	err := r.db.Create(&request)
	if err.Error != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *requestPgRepository) GetAll() ([]*api.RequestGetSchema, error) {
	return nil, nil
}

func (r *requestPgRepository) Get(id int) (*api.RequestGetSchema, error) {
	request := &api.RequestGetSchema{}
	if err := r.db.Model(&models.Log{}).Where("id = ?", id).Find(request); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return request, nil
}

func (r *requestPgRepository) Delete(id int) error {
	err := r.db.Delete(&models.Request{}, id)
	return err.Error
}
