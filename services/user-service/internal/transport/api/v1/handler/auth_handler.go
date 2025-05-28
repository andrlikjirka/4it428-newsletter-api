package handler

import (
	"4it428-newsletter-api/pkg/utils"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"4it428-newsletter-api/services/user-service/internal/transport/api/v1/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
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

	signUpInfo, err := h.authService.SignUp(r.Context(), signUpRequest.ToServiceInput())
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, signUpInfo)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var signInRequest model.SignInRequest
	if err := json.Unmarshal(b, &signInRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(signInRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	signInInfo, err := h.authService.SignIn(r.Context(), signInRequest.Email, signInRequest.Password)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, signInInfo)
}

func (h *AuthHandler) SocialSignIn(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Social signin endpoint hit")
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var refreshTokenRequest model.RefreshTokenRequest
	if err := json.Unmarshal(b, &refreshTokenRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(refreshTokenRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	result, err := h.authService.RefreshToken(r.Context(), refreshTokenRequest.RefreshToken)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, result)
}

func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		utils.WriteErrResponse(w, http.StatusUnauthorized, errors.New("missing or invalid authorization header"))
		return
	}
	idToken := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := h.authService.Verify(r.Context(), idToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(claims)
	// Set headers from claims
	if user_id, ok := claims["user_id"].(uuid.UUID); ok {
		fmt.Println(user_id)
		w.Header().Set("X-User-ID", user_id.String())
	}
	if email, ok := claims["email"].(string); ok {
		w.Header().Set("X-User-Email", email)
	}
	if emailVerif, ok := claims["email_verified"].(bool); ok {
		w.Header().Set("X-User-Email_Verified", strconv.FormatBool(emailVerif))
	}
	if role, ok := claims["role"].(string); ok { // if you use custom claims
		w.Header().Set("X-User-Role", role)
	}
	w.WriteHeader(http.StatusOK) // return 200 OK and no body
}
