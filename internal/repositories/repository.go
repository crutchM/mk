package repositories

import (
	"github.com/jmoiron/sqlx"
	"mk/internal/models"
)

type AuthRepo interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string) (models.User, error)
}

type NotesRepo interface {
	CreateNote(note models.Note) (string, error)
	GetNote(id string) (models.Note, error)
	GetAllNotes(userId int) ([]models.Note, error)
}

type Repository struct {
	AuthRepo
	NotesRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{AuthRepo: NewAuthRepository(db), NotesRepo: NewNotesRepository(db)}
}
