// логика обработки HTTP-запросов.
package handlers

import (
	"net/http"
	"note-service/models"
	"note-service/services"

	"github.com/gin-gonic/gin"
)

func GetNotes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, services.GetAllNotes())
}

func AddNote(context *gin.Context) {
	var newNote models.Note
	if err := context.BindJSON(&newNote); err != nil {
		return
	}
	services.AddNote(newNote)
	context.IndentedJSON(http.StatusCreated, newNote)
}

func GetNoteForText(context *gin.Context) {
	text := context.Param("text")
	note, err := services.GetNoteByText(text)
	if err != nil {
		NotFoundErrorFunc(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, note)
}

func GetNoteForDate(context *gin.Context) {
	data := context.Param("data")
	note, err := services.GetNoteByDate(data)
	if err != nil {
		NotFoundErrorFunc(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, note)
}

func GetNoteForTag(context *gin.Context) {
	tag := context.Param("tag")
	note, err := services.GetNoteByTag(tag)
	if err != nil {
		NotFoundErrorFunc(context, err)
		return
	}
	context.IndentedJSON(http.StatusOK, note)
}

func NotFoundErrorFunc(context *gin.Context, err error) {
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
}
