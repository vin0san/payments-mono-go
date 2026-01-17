package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorBody struct {
	Type    string `json:"type"`
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type APIResponse struct {
	Object  string      `json:"object" example:"response"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
	TraceID string      `json:"trace_id,omitempty"`
}

func Success(c *gin.Context, statusCode int, data interface{}) {
	traceID := c.GetString("trace_id")
	c.JSON(statusCode, APIResponse{
		Object:  "response",
		Success: true,
		Data:    data,
		TraceID: traceID,
	})
}


func Error(ctx *gin.Context, status int, errType, errMsg, details string) {
	traceID := ctx.GetString("trace_id")
	ctx.JSON(status, APIResponse{
		Object:  "response",
		Success: false,
		Data:    nil,
		Error: &ErrorBody{
			Type:    errType,
			Error:   errMsg,
			Details: details,
		},
		TraceID: traceID,
	})
}

func BadRequest(c *gin.Context, message string, details ...interface{}) {
	var det string
	if len(details) > 0 {
		det = details[0].(string)
	}
	Error(c, 400, ErrBadRequest, message, det)
}
