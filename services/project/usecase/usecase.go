package usecase

import (
	"context"

	"dot-rahadian-ardya-kotopanjang/model"
	"dot-rahadian-ardya-kotopanjang/pkg/wrapper"
	"dot-rahadian-ardya-kotopanjang/services/project/repository"

	"github.com/sirupsen/logrus"
)

const (
	errGetProject              = "error when getting project data"
	errCreatingProject         = "error when creating project"
	errUpdatingProject         = "error when update project"
	errUpdatingProjectProgress = "error when update progress"
	errDeletingProject         = "error when deleting project"
)

type IUsecase interface {
	GetAllProject(ctx context.Context) (result wrapper.Result)
	GetProjectById(ctx context.Context, id uint) (result wrapper.Result)
	CreateProject(ctx context.Context, param model.Project) (result wrapper.Result)
	UpdateProject(ctx context.Context, param model.Project) (result wrapper.Result)
	UpdateProjectProgress(ctx context.Context, projectId uint, progress float64) (result wrapper.Result)
	DeleteProject(ctx context.Context, projectId uint) (result wrapper.Result)
}

type usecase struct {
	logger *logrus.Logger
	repo   repository.IRepository
}

// NewProjectUsecase will return IUsecase
func NewProjectUsecase(logger *logrus.Logger, repo repository.IRepository) IUsecase {
	return &usecase{
		logger: logger,
		repo:   repo,
	}
}

// GetAllProject a usecase to get all projects
func (u usecase) GetAllProject(ctx context.Context) (result wrapper.Result) {
	projects, err := u.repo.FindAll()
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errGetProject, wrapper.StatUnexpectedError)
	}

	return wrapper.SuccessResult(projects, wrapper.StatOK)
}

// GetProjectById a usecase to get projects by id
func (u usecase) GetProjectById(ctx context.Context, id uint) (result wrapper.Result) {
	project, err := u.repo.FindById(id)
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errGetProject, wrapper.StatUnexpectedError)
	}
	return wrapper.SuccessResult(project, wrapper.StatOK)
}

// CreateProject a usecase to get create project
func (u usecase) CreateProject(ctx context.Context, param model.Project) (result wrapper.Result) {
	_, err := u.repo.Insert(param)
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errCreatingProject, wrapper.StatUnexpectedError)
	}
	return wrapper.SuccessResult("", wrapper.StatOK)
}

// UpdateProject a usecase to Update project
func (u usecase) UpdateProject(ctx context.Context, param model.Project) (result wrapper.Result) {
	err := u.repo.Update(&param)
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errUpdatingProject, wrapper.StatUnexpectedError)
	}
	return wrapper.SuccessResult("", wrapper.StatOK)
}

// UpdateProject a usecase to update progress project
func (u usecase) UpdateProjectProgress(ctx context.Context, projectId uint, progress float64) (result wrapper.Result) {
	err := u.repo.UpdateByField(projectId, "progress", progress)
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errUpdatingProjectProgress, wrapper.StatUnexpectedError)
	}
	return wrapper.SuccessResult("", wrapper.StatOK)
}

// UpdateProject a usecase to delete project
func (u usecase) DeleteProject(ctx context.Context, projectId uint) (result wrapper.Result) {
	err := u.repo.Delete(projectId)
	if err != nil {
		u.logger.Error(err)
		return wrapper.ErrorResult(err, errDeletingProject, wrapper.StatUnexpectedError)
	}
	return wrapper.SuccessResult("", wrapper.StatOK)
}
