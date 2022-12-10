package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getNotes(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		logrus.Info(err)
		newErrorResponse(c, http.StatusBadRequest, "invalid token")
		return
	}
	notes, err := h.services.GetAllNotes(id)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(c, "notes", notes)

}
