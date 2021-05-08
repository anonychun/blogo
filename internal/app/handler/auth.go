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

type AuthHandler interface {
	Login() http.HandlerFunc
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService}
}

type authHandler struct {
	authService service.AuthService
}

// @Router /accounts/auth [post]
// @Tags auth
// @Summary Login account
// @Description TODO
// @Accept json
// @Produce json
// @Param payload body model.AuthRequest true "body request"
// @Success 200 {object} model.AuthResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 401 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
func (h *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req model.AuthRequest
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

		res, err := h.authService.Login(r.Context(), req)
		if err != nil {
			switch err {
			case constant.ErrEmailNotRegistered, constant.ErrWrongPassword:
				web.MarshalError(w, http.StatusUnauthorized, err)
				return
			default:
				web.MarshalError(w, http.StatusInternalServerError, err)
				return
			}
		}

		web.MarshalPayload(w, http.StatusOK, res)
	}
}
