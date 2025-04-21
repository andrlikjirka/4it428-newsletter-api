package handler

import (
	"4it428-newsletter-api/libs/utils"
	"4it428-newsletter-api/services/user-service/internal/service/iface"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService iface.IUserService
}

func NewUserHandler(s iface.IUserService) *UserHandler {
	return &UserHandler{userService: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Create user endpoint hit")
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "List users endpoint hit")
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := getEmailFromURL(r)
	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Get user by email: %s", email))
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	email := getEmailFromURL(r)
	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Update user by email: %s", email))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	email := getEmailFromURL(r)
	utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Delete user by email: %s", email))
}

func getEmailFromURL(r *http.Request) string {
	return chi.URLParam(r, "email")
}
