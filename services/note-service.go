// для бизнес-логики и услуг
package services

import (
	"errors"
	"note-service/models"
	"strings"
	"time"
)

var notes = []models.Note{
	{ID: "1", Date: time.Now().Format("2006-01-02"), Tag: []string{"тег1", "тег2"}, Text: "Пример отзыва один"},
	{ID: "2", Date: time.Now().Format("2006-01-02"), Tag: []string{"тег3", "тег4"}, Text: "Пример отзыва два"},
	{ID: "3", Date: time.Now().Format("2006-01-02"), Tag: []string{"тег5", "тег6"}, Text: "Пример отзыва три"},
}

func GetAllNotes() []models.Note {
	return notes
}

func AddNote(newNote models.Note) {
	notes = append(notes, newNote)
}

func GetNoteByText(text string) (*models.Note, error) {
	for i, t := range notes {
		if strings.Contains(t.Text, text) {
			return &notes[i], nil
		}
	}
	return nil, errors.New("note not found")
}

func GetNoteByDate(data string) (*models.Note, error) {
	for i, t := range notes {
		if t.Date == data {
			return &notes[i], nil
		}
	}
	return nil, errors.New("note not found")
}

func GetNoteByTag(tag string) (*models.Note, error) {
	for i, t := range notes {
		for _, g := range t.Tag {
			if g == tag {
				return &notes[i], nil
			}
		}
	}
	return nil, errors.New("note not found")
}
