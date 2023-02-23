package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ReceiverPgRepository struct {
	db *gorm.DB
}

func NewReceiverRepository(DbDSN string) *ReceiverPgRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: DbDSN, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Receiver{})
	if err != nil {
		return nil
	}
	return &ReceiverPgRepository{db: db}
}

func (r *ReceiverPgRepository) Add(schema *api.ReceiverPostSchema) error {
	receiver := models.Receiver{
		Name:     schema.Name,
		SocialID: schema.SocialID,
	}
	err := r.db.Create(&receiver)
	if err != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *ReceiverPgRepository) GetAll() ([]*api.ReceiverGetSchema, error) {
	return nil, nil
}

func (r *ReceiverPgRepository) Get(id int) (*api.ReceiverGetSchema, error) {
	receiver := &api.ReceiverGetSchema{}
	if err := r.db.Model(&models.Receiver{}).Where("id = ?", id).Find(&receiver); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return receiver, nil
}

func (r *ReceiverPgRepository) Delete(id int) error {
	r.db.Delete(&models.Receiver{}, id)
	return nil
}
