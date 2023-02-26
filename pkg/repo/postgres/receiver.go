package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
	"gorm.io/gorm"
)

type receiverPgRepository struct {
	db *gorm.DB
}

func NewReceiverRepository(db *gorm.DB) api.ReceiverRepository {
	err := db.AutoMigrate(&models.Receiver{})
	if err != nil {
		return nil
	}
	return &receiverPgRepository{db: db}
}

func (r *receiverPgRepository) Add(schema *api.ReceiverPostSchema) error {
	receiver := models.Receiver{
		Name:     schema.Name,
		SocialID: schema.SocialID,
	}
	err := r.db.Create(&receiver)
	if err.Error != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *receiverPgRepository) GetAll() ([]*api.ReceiverGetSchema, error) {
	return nil, nil
}

func (r *receiverPgRepository) Get(id int) (*api.ReceiverGetSchema, error) {
	receiver := &api.ReceiverGetSchema{}
	if err := r.db.Model(&models.Receiver{}).Where("id = ?", id).Find(&receiver); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return receiver, nil
}

func (r *receiverPgRepository) Delete(id int) error {
	err := r.db.Delete(&models.Receiver{}, id)
	return err.Error
}
