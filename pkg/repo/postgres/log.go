package postgres

import (
	api "PinGo/pkg/api"
	models "PinGo/pkg/repo"
	"errors"
	"gorm.io/gorm"
)

type logPgRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) api.LogRepository {
	err := db.AutoMigrate(&models.Log{})
	if err != nil {
		return nil
	}
	return &logPgRepository{db: db}
}

func (r *logPgRepository) Add(schema *api.LogPostSchema) error {
	log := models.Log{
		ReceiverID:     schema.ReceiverID,
		Name:           schema.Name,
		Address:        schema.Address,
		StatusCode:     schema.StatusCode,
		ResponseTimeMs: schema.ResponseTimeMs,
	}
	err := r.db.Create(&log)
	if err.Error != nil {
		return errors.New("db error")
	}
	return nil
}

func (r *logPgRepository) GetAll() ([]*api.LogGetSchema, error) {
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

func (r *logPgRepository) Get(id int) (*api.LogGetSchema, error) {
	log := &api.LogGetSchema{}
	if err := r.db.Model(&models.Log{}).Where("id = ?", id).Find(log); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	return log, nil
}

func (r *logPgRepository) Delete(id int) error {
	err := r.db.Delete(&models.Log{}, id)
	return err.Error
}
