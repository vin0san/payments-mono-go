package http

import (
	"pye/pkg/response"

	"github.com/gin-gonic/gin"
)

type registerReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "validation_error", "invalid input", "")
		return
	}

	u, err := h.service.Register(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, 500, "db_error", err.Error(), "")
		return
	}

	response.Success(c, 201, u)
}

// Login handler

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var req loginReq
	c.ShouldBindJSON(&req)

	token, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, 401, "auth_error", "invalid credentials", "")
		return
	}

	response.Success(c, 200, gin.H{"token": token})
}
