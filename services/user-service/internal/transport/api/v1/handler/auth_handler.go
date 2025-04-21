package handler

import (
	"4it428-newsletter-api/libs/utils"
	"4it428-newsletter-api/services/user-service/internal/service/iface"
	"net/http"
)

type AuthHandler struct {
	authService iface.IAuthService
}

func NewAuthHandler(s iface.IAuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Signup endpoint hit")
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Signin endpoint hit")
}

func (h *AuthHandler) SocialSignIn(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Social signin endpoint hit")
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Logout endpoint hit")
}

func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Verify  endpoint hit")
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Refresh endpoint hit")
}
