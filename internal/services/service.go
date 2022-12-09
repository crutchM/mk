package services

import (
	"mk/internal/models"
	"mk/internal/repositories"
)

type AuthService interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type NoteService interface {
	CreateNote(note models.InputNote) (string, error)
	GetNote(id string) (models.Note, error)
	GetAllNotes(userId string) ([]models.Note, error)
}

type Service struct {
	AuthService
	NoteService
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{AuthService: NewAuthService(repo.AuthRepo), NoteService: NewNotesService(repo.NotesRepo)}
}
