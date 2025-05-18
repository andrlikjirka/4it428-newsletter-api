package model

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
