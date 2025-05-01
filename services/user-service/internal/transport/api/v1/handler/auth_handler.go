package handler

import (
	"4it428-newsletter-api/libs/utils"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"4it428-newsletter-api/services/user-service/internal/transport/api/v1/model"
	"encoding/json"
	"io"
	"net/http"
)

type AuthHandler struct {
	authService services.IAuthService
}

func NewAuthHandler(s services.IAuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var signUpRequest model.SignUpRequest
	if err := json.Unmarshal(b, &signUpRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(signUpRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := h.authService.SignUp(r.Context(), signUpRequest.ToServiceInput()); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, nil)
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
