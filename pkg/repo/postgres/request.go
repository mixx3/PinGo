package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RequestPgRepository struct {
	db *gorm.DB
}

func NewRequestRepository(DbDSN string) *ReceiverPgRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: DbDSN, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Request{})
	if err != nil {
		return nil
	}
	return &ReceiverPgRepository{db: db}
}

func (r *RequestPgRepository) Add(schema *api.RequestPostSchema) error {
	request := models.Request{
		StatusExpected:         schema.StatusExpected,
		Body:                   schema.Body,
		ExpectedResponseTimeMs: schema.ExpectedResponseTimeMs,
		ReceiverID:             schema.ReceiverID,
	}
	err := r.db.Create(&request)
	if err != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *RequestPgRepository) GetAll() ([]*api.RequestGetSchema, error) {
	return nil, nil
}

func (r *RequestPgRepository) Get(id int) (*api.RequestGetSchema, error) {
	request := &api.RequestGetSchema{}
	if err := r.db.Model(&models.Log{}).Where("id = ?", id).Find(request); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return request, nil
}

func (r *RequestPgRepository) Delete(id int) error {
	r.db.Delete(&models.Request{}, id)
	return nil
}
