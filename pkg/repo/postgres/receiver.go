package postgres

import (
	models "PinGo/pkg/repo"
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
