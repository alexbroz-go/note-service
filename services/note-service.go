// для бизнес-логики и услуг
package services

import (
	"note-service/database"
	"note-service/models"
)

// Возврат новой заметки
func AddNote(newNote models.Note) error {
	return database.AddNote(newNote)
}

// Возврат всех заметок
func GetAllNotes() ([]models.Note, error) {
	return database.GetAllNotes()
}

// Возврат заметки по тексту
func GetNoteByText(text string) (*models.Note, error) {
	return database.GetNoteByText(text)
}

// Возврат заметки по дате
func GetNoteByDate(date string) (*models.Note, error) {
	return database.GetNoteByDate(date)
}

// Возврат списка заметок
func GetNoteByTag(tag string) ([]models.Note, error) {
	return database.GetNoteByTag(tag)
}
