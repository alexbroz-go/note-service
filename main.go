package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID   string   `json:"id"`
	Date string   `json:"date"`
	Tag  []string `json:"tag"`
	Text string   `json:"text"`
}

var timen time.Time = time.Now()
var formattedDate string = timen.Format("2006-01-02")

var notes = []note{
	{ID: "1", Date: formattedDate, Tag: []string{"тег1", "тег2"}, Text: "Пример отзыва один"},
	{ID: "2", Date: formattedDate, Tag: []string{"тег3", "тег4"}, Text: "Пример отзыва два"},
	{ID: "3", Date: formattedDate, Tag: []string{"тег5", "тег6"}, Text: "Пример отзыва три "},
}

func getNotes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, notes)

}

func addNote(context *gin.Context) {
	var newNote note
	if err := context.BindJSON(&newNote); err != nil {
		return
	}
	notes = append(notes, newNote)
	context.IndentedJSON(http.StatusCreated, newNote)
}

func getNoteForText(context *gin.Context) {
	text := context.Param("text")
	note, err := getNoteByText(text)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"massage": "Не найдено"})
		return
	}
	context.IndentedJSON(http.StatusOK, note)

}

func getNoteByText(text string) (*note, error) {
	for i, t := range notes {
		if strings.Contains(t.Text, text) {
			return &notes[i], nil
		}
	}
	return nil, errors.New("note not found")
}

func getNoteForDate(context *gin.Context) {
	data := context.Param("data")
	note, err := getNoteByDate(data)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"massage": "Не найдено"})
		return
	}
	context.IndentedJSON(http.StatusOK, note)

}

func getNoteByDate(data string) (*note, error) {
	for i, t := range notes {
		if t.Date == data {
			return &notes[i], nil
		}
	}
	return nil, errors.New("note not found")
}

func getNoteForTag(context *gin.Context) {
	tag := context.Param("tag")
	note, err := getNoteByTag(tag)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"massage": "Не найдено"})
		return
	}
	context.IndentedJSON(http.StatusOK, note)

}

func getNoteByTag(tag string) (*note, error) {
	for i, t := range notes {
		for _, g := range t.Tag {
			if g == tag {
				return &notes[i], nil
			}
		}
	}
	return nil, errors.New("note not found")
}

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/text/:text", getNoteForText)
	router.GET("/notes/data/:data", getNoteForDate)
	router.GET("/notes/tag/:tag", getNoteForTag)
	router.POST("/notes", addNote)
	router.Run("localhost:1010")
}
