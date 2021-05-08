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

type AccountHandler interface {
	Create() http.HandlerFunc
	List() http.HandlerFunc
	Get() http.HandlerFunc
	Update() http.HandlerFunc
	UpdatePassword() http.HandlerFunc
	Delete() http.HandlerFunc
}

func NewAccountHandler(accountService service.AccountService) AccountHandler {
	return &accountHandler{accountService}
}

type accountHandler struct {
	accountService service.AccountService
}

// @Router /accounts [post]
// @Tags accounts
// @Summary Create account
// @Description TODO
// @Accept json
// @Produce json
// @Param payload body model.AccountCreateRequest true "body request"
// @Success 201 {object} model.AccountResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *accountHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.AccountCreateRequest
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

		res, err := h.accountService.Create(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrEmailRegistered:
				web.MarshalError(w, http.StatusConflict, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		web.MarshalPayload(w, http.StatusCreated, res)
	}
}

// @Router /accounts [get]
// @Tags accounts
// @Summary List accounts
// @Description TODO
// @Produce json
// @Param limit query int false "pagination limit"
// @Param offset query int false "pagination offset"
// @Param name query string false "account name"
// @Success 200 {array} model.AccountResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *accountHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset, err := web.GetPagination(r)
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.AccountListRequest{
			Limit:  limit,
			Offset: offset,
			Name:   web.GetUrlQueryString(r, "name"),
		}

		res, err := h.accountService.List(r.Context(), req)
		if err != nil {
			web.MarshalError(w, http.StatusInternalServerError, err)
			return
		}

		web.MarshalPayload(w, http.StatusOK, res)
	}
}

// @Router /accounts/{account_id} [get]
// @Tags accounts
// @Summary Get account
// @Description TODO
// @Accept json
// @Produce json
// @Param account_id path int true "account id" Format(int64)
// @Success 200 {object} model.AccountResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *accountHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "account_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.AccountGetRequest{ID: id}
		res, err := h.accountService.Get(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrAccountNotFound:
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

// @Router /accounts/{account_id} [put]
// @Tags accounts
// @Summary Update account
// @Description TODO
// @Accept json
// @Produce json
// @Param account_id path int true "account id" Format(int64)
// @Param payload body model.AccountUpdateRequest true "body request"
// @Success 200 {object} model.AccountResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 409 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *accountHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "account_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.AccountUpdateRequest{ID: id}
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

		res, err := h.accountService.Update(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			case constant.ErrEmailRegistered:
				web.MarshalError(w, http.StatusConflict, err)
				return
			case constant.ErrAccountNotFound:
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

// @Router /accounts/{account_id}/password [put]
// @Tags accounts
// @Summary Update account password
// @Description TODO
// @Accept json
// @Produce json
// @Param account_id path int true "account id" Format(int64)
// @Param payload body model.AccountPasswordUpdateRequest true "body request"
// @Success 200 {object} model.AccountResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *accountHandler) UpdatePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "account_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.AccountPasswordUpdateRequest{ID: id}
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

		res, err := h.accountService.UpdatePassword(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized, constant.ErrWrongPassword:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			case constant.ErrAccountNotFound:
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

// @Router /accounts/{account_id} [delete]
// @Tags accounts
// @Summary Delete account
// @Description TODO
// @Produce json
// @Param account_id path int true "account id" Format(int64)
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (h *accountHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := web.GetUrlPathInt64(r, "account_id")
		if err != nil {
			web.MarshalError(w, http.StatusBadRequest, err)
			return
		}

		req := model.AccountDeleteRequest{ID: id}
		err = h.accountService.Delete(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrUnauthorized:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			case constant.ErrAccountNotFound:
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
