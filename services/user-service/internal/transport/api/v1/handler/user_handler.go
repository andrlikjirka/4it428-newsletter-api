package handler

import (
	"4it428-newsletter-api/libs/utils"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"4it428-newsletter-api/services/user-service/internal/transport/api/v1/model"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(s services.IUserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var userRequest model.CreateUserRequest
	if err := json.Unmarshal(b, &userRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(userRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.userService.CreateUser(r.Context(), userRequest.ToUser()); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, userRequest)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	userList := h.userService.ListUsers(r.Context())
	response := model.FromUserList(userList)
	utils.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := getEmailFromURL(r)
	user, err := h.userService.GetUserByEmail(r.Context(), email)
	if err != nil {
		utils.WriteResponse(w, http.StatusNotFound, err)
		return
	}
	response := model.FromUser(user)
	utils.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var userRequest model.UpdateUserRequest
	if err := json.Unmarshal(b, &userRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(userRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	email := getEmailFromURL(r)
	updatedUser, err := h.userService.UpdateUser(r.Context(), email, userRequest.ToUserUpdate())
	if err != nil {
		utils.WriteResponse(w, http.StatusNotFound, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, updatedUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := getEmailFromURL(r)
	err := h.userService.DeleteUser(r.Context(), email)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusNotFound, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}

// ===========================================
// PRIVATE HELPER FUNCTIONS

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")
	return email
}
