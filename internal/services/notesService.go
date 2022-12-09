package services

import (
	"mk/internal/models"
	"mk/internal/repositories"
)

type NotesService struct {
	repo repositories.NotesRepo
}

func (n NotesService) CreateNote(note models.InputNote) (string, error) {
	return n.repo.CreateNote(note.MapInputToNote())
}

func (n NotesService) GetNote(id string) (models.Note, error) {
	return n.repo.GetNote(id)
}

func (n NotesService) GetAllNotes(userId string) ([]models.Note, error) {
	return n.repo.GetAllNotes(userId)
}

func NewNotesService(repo repositories.NotesRepo) *NotesService {
	return &NotesService{repo: repo}
}
