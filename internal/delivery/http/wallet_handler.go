package http

import (
	"net/http"

	"pye/internal/app"
	"pye/pkg/response"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	service *app.WalletService
}

func NewWalletHandler(s *app.WalletService) *WalletHandler {
	return &WalletHandler{service: s}
}

func (h *WalletHandler) GetBalance(c *gin.Context) {
	userID := c.GetString("user_id")

	bal, err := h.service.GetBalance(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, 500, "db_error", err.Error(), "")
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"balance": bal,
	})
}
