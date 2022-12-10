package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, NewError(message))
}

func SendJSONResponse(c *gin.Context, key string, value interface{}) {
	if value == nil {
		logrus.Error("not enough token")
		newErrorResponse(c, http.StatusInternalServerError, "not enough token")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		key: value,
	})
}
