package http

import (
	"net/http"

	"pye/internal/app"
	"pye/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *app.UserService
}

func NewUserHandler(s *app.UserService) *UserHandler {
	return &UserHandler{service: s}
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "validation_error", "invalid request body", "")
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "db_error", err.Error(), "")
		return
	}

	response.Success(c, http.StatusCreated, user)
}
