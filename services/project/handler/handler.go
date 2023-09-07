package handler

import (
	"encoding/json"
	"net/http"

	"dot-rahadian-ardya-kotopanjang/model"
	"dot-rahadian-ardya-kotopanjang/pkg/helper"
	"dot-rahadian-ardya-kotopanjang/pkg/wrapper"
	"dot-rahadian-ardya-kotopanjang/services/project/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	usecase usecase.IUsecase
}

// NewProjectHandler is to handle router for project
func NewProjectHandler(router *mux.Router, logger *logrus.Logger, usecase usecase.IUsecase) {
	httpHandler := &Handler{
		logger:  logger,
		usecase: usecase,
	}

	// get all
	router.HandleFunc("/project", httpHandler.GetAll).Methods(http.MethodGet)
	// get by id
	router.HandleFunc("/project/{id}", httpHandler.GetById).Methods(http.MethodGet)
	// insert
	router.HandleFunc("/project", httpHandler.Create).Methods(http.MethodPost)
	// update data
	router.HandleFunc("/project", httpHandler.Update).Methods(http.MethodPut)
	// update progress
	router.HandleFunc("/project/progress/{id}", httpHandler.UpdateProgress).Methods(http.MethodPatch)
	// delete
	router.HandleFunc("/project/{id}", httpHandler.Delete).Methods(http.MethodDelete)
}

// GetAll is a handler to get all projects
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	result = h.usecase.GetAllProject(ctx)
	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "")
}

// GetById is a handler to get project by id
func (h Handler) GetById(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	pathVariables := mux.Vars(r)
	id := helper.ToInt(pathVariables["id"])

	result = h.usecase.GetProjectById(ctx, uint(id))
	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "")
}

// Create is a handler to create project
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	var param model.Project

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		result = wrapper.ErrorResult(wrapper.ErrBadRequest, "", wrapper.StatInvalidRequest)
		wrapper.ResponseError(w, &result)
		return
	}

	result = h.usecase.CreateProject(ctx, param)
	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "")
}

// UpdateProgress is a handler to update whole project
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	var param model.Project

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		result = wrapper.ErrorResult(wrapper.ErrBadRequest, "", wrapper.StatInvalidRequest)
		wrapper.ResponseError(w, &result)
		return
	}

	result = h.usecase.UpdateProject(ctx, param)
	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "")
}

// UpdateProgress is a handler to update progress project
func (h Handler) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	pathVariables := mux.Vars(r)
	id := helper.ToInt(pathVariables["id"])

	var param map[string]any
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		result = wrapper.ErrorResult(wrapper.ErrBadRequest, "", wrapper.StatInvalidRequest)
		wrapper.ResponseError(w, &result)
		return
	}

	progres, ok := param["progress"]
	if !ok {
		result = wrapper.ErrorResult(wrapper.ErrBadRequest, "", wrapper.StatInvalidRequest)
		wrapper.ResponseError(w, &result)
		return
	}
	progressInt := helper.ToFloat64(progres, 2, helper.RoundingAuto)

	result = h.usecase.UpdateProjectProgress(ctx, uint(id), progressInt)

	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "progress updated")
}

// Delete is a handler to delete project
func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	result := wrapper.Result{}
	ctx := r.Context()

	pathVariables := mux.Vars(r)
	id := helper.ToInt(pathVariables["id"])

	result = h.usecase.DeleteProject(ctx, uint(id))

	if result.Err != nil {
		wrapper.ResponseError(w, &result)
		return
	}

	wrapper.ResponseSuccess(w, http.StatusOK, &result, "project deleted")
}
