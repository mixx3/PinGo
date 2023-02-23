package postgres

import (
	api "PinGo/pkg/api"
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

func (r *LogPgRepository) Add(schema *api.LogPostSchema) error {
	return nil
}
func (r *LogPgRepository) GetAll() ([]*api.LogGetSchema, error) {
	return nil, nil
}
func (r *LogPgRepository) Get(id int) (*api.LogGetSchema, error) {
	return nil, nil
}

func (r *LogPgRepository) Delete(id int) error {
	return nil
}
