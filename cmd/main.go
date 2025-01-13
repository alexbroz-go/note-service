// основная точка входа в приложение
package main

import (
	"note-service/database"
	"note-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()        // Инициализация базы данных
	defer database.CloseDB() // Закрытие соединения при завершении

	router := gin.Default()
	router.GET("/notes", handlers.GetNotes)
	router.POST("/notes", handlers.AddNote)
	router.GET("/notes/text/:text", handlers.GetNoteForText)
	router.GET("/notes/data/:data", handlers.GetNoteForDate)
	router.GET("/notes/tag/:tag", handlers.GetNoteForTag)
	router.Run("localhost:1010") // Запуск сервера на порту 1010
}
