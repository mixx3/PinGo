package postgres

import (
	models "PinGo/pkg/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LogPgRepository struct {
	db *gorm.DB
}

func NewLogRepository(DbDSN string) *LogPgRepository {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: DbDSN, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil
	}
	err = db.AutoMigrate(&models.Log{})
	if err != nil {
		return nil
	}
	return &LogPgRepository{db: db}
}
