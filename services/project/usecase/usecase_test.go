package usecase_test

import (
	"testing"

	mock_repository "dot-rahadian-ardya-kotopanjang/services/project/repository/mocks"
	"dot-rahadian-ardya-kotopanjang/services/project/usecase"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_new_project_usecase_with_logger_and_repo(t *testing.T) {
	logger := logrus.New()
	repo := mock_repository.NewIRepository(t)
	usecase := usecase.NewProjectUsecase(logger, repo)

	assert.NotNil(t, usecase)

}
