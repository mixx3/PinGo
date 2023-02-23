package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
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
	log := models.Log{
		ReceiverID:     schema.ReceiverID,
		Name:           schema.Name,
		Address:        schema.Address,
		StatusCode:     schema.StatusCode,
		ResponseTimeMs: schema.ResponseTimeMs,
	}
	err := r.db.Create(&log)
	if err != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *LogPgRepository) GetAll() ([]*api.LogGetSchema, error) {
	ids := make([]int64, 0)
	res := make([]*api.LogGetSchema, 0)
	if err := r.db.Model(&models.Log{}).Select("name").Find(&ids).Error; err != nil {
		for id := range ids {
			cur, _ := r.Get(id)
			res = append(res, cur)
		}
		return res, nil
	}
	return res, nil
}
func (r *LogPgRepository) Get(id int) (*api.LogGetSchema, error) {
	log := models.Log{}
	if err := r.db.Where("id = ?", id).First(&log); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return &api.LogGetSchema{
		ID:             id,
		ReceiverID:     log.ReceiverID,
		Name:           log.Name,
		Address:        log.Address,
		StatusCode:     log.StatusCode,
		ResponseTimeMs: log.ResponseTimeMs,
		DtCreated:      log.CreatedAt,
	}, nil
}

func (r *LogPgRepository) Delete(id int) error {
	r.db.Delete(&models.Log{}, id)
	return nil
}
