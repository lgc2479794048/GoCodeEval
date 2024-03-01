package repository

import (
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

// Create 插入新记录到数据库
func (r *GormRepository) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

// Update 根据条件更新记录
func (r *GormRepository) Update(model interface{}, conditions interface{}, updates interface{}) error {
	return r.DB.Model(model).Where(conditions).Updates(updates).Error
}

// Delete 根据条件删除记录
func (r *GormRepository) Delete(model interface{}, conditions interface{}) error {
	return r.DB.Where(conditions).Delete(model).Error
}

// FindOne 根据条件查询单个记录
func (r *GormRepository) FindOne(model interface{}, conditions interface{}) error {
	return r.DB.Where(conditions).First(model).Error
}

// Find 根据条件查询多个记录
func (r *GormRepository) Find(models interface{}, conditions interface{}, orders ...string) error {
	query := r.DB.Where(conditions)
	for _, order := range orders {
		query = query.Order(order)
	}
	return query.Find(models).Error
}

// PreloadRelated 预加载关联数据
func (r *GormRepository) PreloadRelated(model interface{}, conditions interface{}, preloads ...string) error {
	query := r.DB.Where(conditions)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	return query.First(model).Error
}

// Transaction 提供一个执行数据库事务的函数
func (r *GormRepository) Transaction(fc func(txRepo *GormRepository) error) (err error) {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit().Error // if Commit returns error update err with Commit error
		}
	}()

	txRepo := NewGormRepository(tx)
	err = fc(txRepo)
	return err
}
