package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mk/internal/models"
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

func (h *Handler) createNote(c *gin.Context) {
	var input models.InputNote

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateNote(input)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(c, "newNote", id)
}

func (h *Handler) getNote(c *gin.Context) {
	id := c.Param("id")

	note, err := h.services.GetNote(id)
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(c, "note", note)

}
