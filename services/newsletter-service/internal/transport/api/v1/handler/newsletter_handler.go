package handler

import (
	"4it428-newsletter-api/libs/utils"
	errorsdef "4it428-newsletter-api/services/newsletter-service/internal/service/errors"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/model"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type NewsletterHandler struct {
	newsletterService services.NewsletterService
}

func NewNewsletterHandler(s services.NewsletterService) *NewsletterHandler {
	return &NewsletterHandler{newsletterService: s}
}

func (h *NewsletterHandler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var newsletterRequest model.CreateNewsletterRequest
	if err := json.Unmarshal(b, &newsletterRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(newsletterRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	n, err := h.newsletterService.CreateNewsletter(r.Context(), newsletterRequest.ToNewsletter())
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, model.FromNewsletter(n))
}

func (h *NewsletterHandler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
	newsletters, err := h.newsletterService.ListNewsletters(r.Context())
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, model.FromNewsletterList(newsletters))
}

func (h *NewsletterHandler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	newsletter, err := h.newsletterService.GetNewsletterById(r.Context(), id)
	if err != nil {
		if errors.Is(err, errorsdef.ErrInvalidUUID) {
			utils.WriteErrResponse(w, http.StatusBadRequest, err)
			return
		} else if errors.Is(err, errorsdef.ErrNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, model.FromNewsletter(newsletter))
}

func (h *NewsletterHandler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var request model.UpdateNewsletterRequest
	if err := json.Unmarshal(b, &request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	id := chi.URLParam(r, "id")
	updatedNewsletter, err := h.newsletterService.UpdateNewsletter(r.Context(), id, request.ToNewsletterUpdate())
	if err != nil {
		if errors.Is(err, errorsdef.ErrNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, updatedNewsletter)
}

func (h *NewsletterHandler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := h.newsletterService.DeleteNewsletter(r.Context(), id)
	if err != nil {
		if errors.Is(err, errorsdef.ErrInvalidUUID) {
			utils.WriteErrResponse(w, http.StatusBadRequest, err)
			return
		} else if errors.Is(err, errorsdef.ErrNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}
