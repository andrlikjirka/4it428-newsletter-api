package handler

import (
	"4it428-newsletter-api/libs/utils"
	errorsdef "4it428-newsletter-api/services/newsletter-service/internal/service/errors"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/model"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type PostHandler struct {
	postService services.PostService
}

func NewPostHandler(s services.PostService) *PostHandler {
	return &PostHandler{postService: s}
}

func (h *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	newsletterID := chi.URLParam(r, "newsletterID")
	posts, err := h.postService.ListPosts(r.Context(), newsletterID)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, model.FromPostList(posts))
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	newsletterID := chi.URLParam(r, "newsletterID")

	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var postRequest model.CreatePostRequest
	if err := json.Unmarshal(b, &postRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(postRequest); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	p, err := h.postService.CreatePost(r.Context(), postRequest.ToPost(), newsletterID, utils.GetXUserId(r))
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	utils.WriteResponse(w, http.StatusCreated, model.FromPost(p))
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	newsletterID := chi.URLParam(r, "newsletterID")
	postID := chi.URLParam(r, "postID")
	p, err := h.postService.GetPostById(r.Context(), postID, newsletterID)
	if err != nil {
		if errors.Is(err, errorsdef.ErrInvalidUUID) {
			utils.WriteErrResponse(w, http.StatusBadRequest, err)
			return
		} else if errors.Is(err, errorsdef.ErrPostNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, model.FromPost(p))
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var request model.UpdatePostRequest
	if err := json.Unmarshal(b, &request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(request); err != nil {
		utils.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}

	postID := chi.URLParam(r, "postID")
	newsletterID := chi.URLParam(r, "newsletterID")
	updatedPost, err := h.postService.UpdatePost(r.Context(), postID, newsletterID, utils.GetXUserId(r), request.ToPostUpdate())
	if err != nil {
		if errors.Is(err, errorsdef.ErrPostNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, updatedPost)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	newsletterID := chi.URLParam(r, "newsletterID")
	err := h.postService.DeletePost(r.Context(), postID, newsletterID, utils.GetXUserId(r))
	if err != nil {
		if errors.Is(err, errorsdef.ErrInvalidUUID) {
			utils.WriteErrResponse(w, http.StatusBadRequest, err)
			return
		} else if errors.Is(err, errorsdef.ErrPostNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}

func (h *PostHandler) PublishPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	newsletterID := chi.URLParam(r, "newsletterID")

	err := h.postService.PublishPost(r.Context(), postID, newsletterID, utils.GetXUserId(r))
	if err != nil {
		if errors.Is(err, errorsdef.ErrInvalidUUID) {
			utils.WriteErrResponse(w, http.StatusBadRequest, err)
			return
		} else if errors.Is(err, errorsdef.ErrPostNotFound) {
			utils.WriteErrResponse(w, http.StatusNotFound, err)
			return
		}
		utils.WriteErrResponse(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteResponse(w, http.StatusNoContent, nil)
}
