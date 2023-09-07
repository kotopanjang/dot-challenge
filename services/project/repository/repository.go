package repository

import (
	"dot-rahadian-ardya-kotopanjang/model"
	"dot-rahadian-ardya-kotopanjang/pkg/helper"
	"dot-rahadian-ardya-kotopanjang/pkg/redis"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRepository interface {
	FindAll() ([]model.Project, error)
	FindById(id uint) (*model.Project, error)
	Insert(projectParam model.Project) (uint, error)
	Update(project *model.Project) error
	UpdateByField(projectId uint, field string, value any) error
	Delete(id uint) error
}

type repository struct {
	logger *logrus.Logger
	db     *gorm.DB
	cache  *redis.RedisStorage
}

// NewProjectRepository will return IRepository
func NewProjectRepository(db *gorm.DB, cache *redis.RedisStorage, logger *logrus.Logger) IRepository {
	return &repository{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

// FindAll is to find all project
func (p repository) FindAll() ([]model.Project, error) {
	projects := []model.Project{}
	result := p.db.Preload("Members").Find(&projects)

	if result.Error != nil {
		p.logger.Error(result.Error)
		return projects, result.Error
	}

	return projects, nil
}

// FindById is to find project by id.
// first, will retrieve from cache. if not found, it will load from database
func (p repository) FindById(id uint) (*model.Project, error) {
	project := &model.Project{}

	// retrieve from redis
	proj, err := p.getFromRedis(id)
	if err == nil && project != nil {
		log.Println("getting data from REDIS. id: " + helper.ToString(id))
		return proj, nil
	}

	// retrieve from database
	log.Println("getting data from DATABASE. id: " + helper.ToString(id))
	result := p.db.Preload("Members").First(project, id)
	if result.Error != nil {
		p.logger.Error(result.Error)
		return project, result.Error
	}

	// add to redis
	p.addToRedis(*project)

	return project, nil
}

// Insert is to Insert Project
func (p repository) Insert(projectParam model.Project) (uint, error) {
	result := p.db.Create(&projectParam)

	if result.Error != nil {
		p.logger.Error(result.Error)
		return 0, result.Error
	}

	return projectParam.ID, nil
}

// Update is to Update Project with all fields
func (p repository) Update(project *model.Project) error {
	// delete from redis
	p.deleteFromRedis(project.ID)

	result := p.db.Save(&project)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}

	return result.Error
}

// UpdateByField is to Update project by field name.
// field is fieldname in databaes
func (p repository) UpdateByField(projectId uint, field string, value any) error {
	// delete from redis
	p.deleteFromRedis(projectId)

	result := p.db.Model(&model.Project{}).Where("id = ?", projectId).Update(field, value)
	if result.Error != nil {
		p.logger.Error(result.Error)
	}

	return result.Error
}

// Delete is to disable project.
// data is not actually deleted, but flagged as deleted with inserting date to column `DeletedAt`
func (p repository) Delete(id uint) error {
	// delete from redis
	p.deleteFromRedis(id)

	project := &model.Project{}
	result := p.db.First(&project, id)

	if result.Error != nil {
		p.logger.Error(result.Error)
		return result.Error
	}

	result = p.db.Delete(&project)
	return result.Error
}

// addToRedis is to add key value pair to redis
func (p repository) addToRedis(proj model.Project) error {
	err := p.cache.Set(helper.ToString(proj.ID), &proj)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	return nil
}

// deleteFromRedis is to delete data from redis
func (p repository) deleteFromRedis(id uint) error {
	err := p.cache.Del(helper.ToString(id))
	if err != nil {
		p.logger.Error(err)
		return err
	}
	return nil
}

// getFromRedis is to get data from redis
func (p repository) getFromRedis(id uint) (*model.Project, error) {
	project := &model.Project{}
	err := p.cache.GetByKey(helper.ToString(id), project)
	if err != nil {
		return nil, err
	}

	return project, nil
}
