package postgres

import (
	models "PinGo/pkg/repo"
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
	err = db.AutoMigrate(&models.Log{})
	if err != nil {
		return nil
	}
	return &ReceiverPgRepository{db: db}
}
