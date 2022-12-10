package services

import (
	"mk/internal/models"
	"mk/internal/repositories"
)

type AuthService interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type NoteService interface {
	CreateNote(note models.InputNote) (string, error)
	GetNote(id string) (models.Note, error)
	GetAllNotes(userId int) ([]models.Note, error)
}

type Service struct {
	AuthService
	NoteService
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{AuthService: NewAuthService(repo.AuthRepo), NoteService: NewNotesService(repo.NotesRepo)}
}
