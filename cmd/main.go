// основная точка входа в приложение
package main

import (
	"note-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/notes", handlers.GetNotes)
	router.GET("/notes/text/:text", handlers.GetNoteForText)
	router.GET("/notes/data/:data", handlers.GetNoteForDate)
	router.GET("/notes/tag/:tag", handlers.GetNoteForTag)
	router.POST("/notes", handlers.AddNote)
	router.Run("localhost:1010")
}
