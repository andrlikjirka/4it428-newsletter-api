package handler

import (
	"4it428-newsletter-api/libs/utils"
	errorsdef "4it428-newsletter-api/services/subscription-service/internal/service/errors"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"4it428-newsletter-api/services/subscription-service/internal/transport/api/v1/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type SubscriptionHandler struct {
	subscriptionService services.SubscriptionService
}

func NewSubscriptionHandler(s services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{subscriptionService: s}
}

func (h *SubscriptionHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	subscriptionID := chi.URLParam(r, "subscriptionID")
	err := h.subscriptionService.Unsubscribe(r.Context(), subscriptionID)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}

func (h *SubscriptionHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var request model.SubscriptionRequest
	if err := json.Unmarshal(b, &request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	subscription, err := h.subscriptionService.Subscribe(r.Context(), request.ToSubscription())
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, model.FromSubscription(subscription))
}

func (h *SubscriptionHandler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	newsletterID := utils.GetNewsletterIdFromQueryParam(r)
	if newsletterID == "" {
		utils.WriteErrResponse(w, http.StatusBadRequest, errorsdef.ErrNoNewsletterId)
		return
	}

	subscriptions, err := h.subscriptionService.ListSubscriptions(r.Context(), newsletterID, utils.GetXUserId(r))

	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, model.FromSubscriptionList(subscriptions))
}

func (h *SubscriptionHandler) NotifySubscribers(w http.ResponseWriter, r *http.Request) {
	newsletterID := utils.GetNewsletterIdFromQueryParam(r)

	if newsletterID == "" {
		utils.WriteErrResponse(w, http.StatusBadRequest, errorsdef.ErrNoNewsletterId)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var request model.NotifySubscribersRequest
	if err := json.Unmarshal(b, &request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	err = h.subscriptionService.NotifySubscribers(r.Context(), newsletterID, request.ToNotification())
	if err != nil {
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}
