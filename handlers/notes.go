// логика обработки HTTP-запросов.
package handlers

import (
	"net/http"
	"note-service/models"
	"note-service/services"

	"github.com/gin-gonic/gin"
)

// GetNotes возвращает список всех заметок
func GetNotes(context *gin.Context) {
	notes, err := services.GetAllNotes() // Получаем все заметки из сервиса
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // Возвращаем ошибку для отладки
		return
	}
	context.IndentedJSON(http.StatusOK, notes) // Возвращаем заметки в формате JSON
}

// AddNote добавляет новую заметку
func AddNote(context *gin.Context) {
	var newNote models.Note
	if err := context.BindJSON(&newNote); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	err := services.AddNote(newNote)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error adding note"})
		return
	}
	context.IndentedJSON(http.StatusCreated, newNote)
}

// GetNoteForText возвращает заметку по тексту
func GetNoteForText(context *gin.Context) {
	text := context.Param("text")
	note, err := services.GetNoteByText(text)
	if err != nil {
		NotFoundErrorFunc(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, note)
}

// GetNoteForDate возвращает заметку по дате
func GetNoteForDate(context *gin.Context) {
	date := context.Param("data")
	note, err := services.GetNoteByDate(date)
	if err != nil {
		NotFoundErrorFunc(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, note)
}

// GetNoteForTag возвращает заметку по тегу
func GetNoteForTag(context *gin.Context) {
	tag := context.Param("tag")
	notes, err := services.GetNoteByTag(tag) // Извлечение заметок по тегу
	if err != nil {
		NotFoundErrorFunc(context, err) // Обработка ошибки
		return
	}
	if len(notes) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No notes found for this tag"})
		return // Если заметок не найдено, возвращаем 404
	}
	context.IndentedJSON(http.StatusOK, notes) // Возврат найденных заметок
}

// NotFoundErrorFunc обрабатывает ошибки 404
func NotFoundErrorFunc(context *gin.Context, err error) {
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}
