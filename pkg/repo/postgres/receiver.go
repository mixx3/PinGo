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
	log := models.Log{}
	err := r.db.Create(&log)
	if err != nil {
		return errors.New("WTF")
	}
	return nil
	return nil
}

func (r *ReceiverPgRepository) GetAll() ([]*api.ReceiverGetSchema, error) {
	return nil, nil
}

func (r *ReceiverPgRepository) Get(id int) (*api.ReceiverGetSchema, error) {
	return nil, nil
}

func (r *ReceiverPgRepository) Delete(id int) error {
	return nil
}
