package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anonychun/go-blog-api/internal/app/model"
	"github.com/anonychun/go-blog-api/internal/app/service"
	"github.com/anonychun/go-blog-api/internal/constant"
	"github.com/anonychun/go-blog-api/internal/validation"
	"github.com/anonychun/go-blog-api/internal/web"
)

type PostHandler interface {
	Create() http.HandlerFunc
	List() http.HandlerFunc
	Get() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
}

func NewPostHandler(postService service.PostService) PostHandler {
	return &postHandler{postService}
}

type postHandler struct {
	postService service.PostService
}

// @Router /posts [post]
// @Tags posts
// @Summary Create post
// @Description TODO
// @Accept json
// @Produce json
// @Param payload body model.PostCreateRequest true "body request"
// @Success 201 {object} model.PostResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *postHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.PostCreateRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, constant.ErrRequestBody)
			return
		}

		err = validation.Struct(req)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		res, err := h.postService.Create(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		web.MarshalPayload(w, http.StatusCreated, res)
	}
}

// @Router /posts [get]
// @Tags posts
// @Summary List posts
// @Description TODO
// @Produce json
// @Param limit query int false "pagination limit"
// @Param offset query int false "pagination offset"
// @Param title query string false "post title"
// @Success 200 {array} model.PostResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *postHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset, err := web.GetPagination(r)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.PostListRequest{
			Limit:  limit,
			Offset: offset,
			Title:  web.GetUrlQueryString(r, "title"),
		}

		res, err := h.postService.List(r.Context(), req)
		if err != nil {
			web.MarshalError(w, http.StatusInternalServerError, err)
			return
		}

		web.MarshalPayload(w, http.StatusOK, res)
	}
}

// @Router /posts/{post_id} [get]
// @Tags posts
// @Summary Get post
// @Description TODO
// @Accept json
// @Produce json
// @Param post_id path int true "post id" Format(int64)
// @Success 200 {object} model.PostResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *postHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "post_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.PostGetRequest{ID: id}
		res, err := h.postService.Get(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrPostNotFound:
				web.MarshalError(w, http.StatusNotFound, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		web.MarshalPayload(w, http.StatusOK, res)
	}
}

// @Router /posts/{post_id} [put]
// @Tags posts
// @Summary Update post
// @Description TODO
// @Accept json
// @Produce json
// @Param post_id path int true "post id" Format(int64)
// @Param payload body model.PostUpdateRequest true "body request"
// @Success 200 {object} model.PostResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *postHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "post_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.PostUpdateRequest{ID: id}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, constant.ErrRequestBody)
			return
		}

		err = validation.Struct(req)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		res, err := h.postService.Update(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			case constant.ErrPostNotFound:
				web.MarshalError(w, http.StatusNotFound, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		web.MarshalPayload(w, http.StatusOK, res)
	}
}

// @Router /posts/{post_id} [delete]
// @Tags posts
// @Summary Delete post
// @Description TODO
// @Produce json
// @Param post_id path int true "post id" Format(int64)
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *postHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "post_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.PostDeleteRequest{ID: id}
		err = h.postService.Delete(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			case constant.ErrPostNotFound:
				web.MarshalError(w, http.StatusNotFound, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
