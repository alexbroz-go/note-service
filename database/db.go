// database/db.go
package database

import (
	"database/sql"
	"log"

	"note-service/models"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB инициализация базы данных
func InitDB() {
	var err error
	connStr := "user=user password=password dbname=notes_db host=localhost port=5438 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	CreateNoteTable()
}

// CloseDB закрывает соединение с базой данных
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// CreateNoteTable создает таблицу заметок
func CreateNoteTable() {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id VARCHAR(50) PRIMARY KEY,
		date VARCHAR(20),
		tag TEXT[],
		text TEXT
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create notes table: %v", err)
	}
}

// AddNote добавляет новую заметку в базу данных
func AddNote(note models.Note) error {
	_, err := db.Exec("INSERT INTO notes (id, date, tag, text) VALUES ($1, $2, $3, $4)", note.ID, note.Date, pq.Array(note.Tag), note.Text)
	return err
}

// GetAllNotes возвращает все заметки из базы данных
func GetAllNotes() ([]models.Note, error) {
	rows, err := db.Query("SELECT id, date, tag, text FROM notes")
	if err != nil {
		return nil, err // Возврат ошибки для отладки
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Date, pq.Array(&note.Tag), &note.Text); err != nil {
			return nil, err // Возврат ошибки для отладки
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err // Проверка на ошибку в итерации
	}

	return notes, nil
}

// GetNoteByText возвращает заметку по тексту
func GetNoteByText(text string) (*models.Note, error) {
	var note models.Note
	err := db.QueryRow("SELECT id, date, tag, text FROM notes WHERE text ILIKE $1 LIMIT 1", "%"+text+"%").Scan(&note.ID, &note.Date, pq.Array(&note.Tag), &note.Text)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// GetNoteByDate возвращает заметку по дате
func GetNoteByDate(date string) (*models.Note, error) {
	var note models.Note
	err := db.QueryRow("SELECT id, date, tag, text FROM notes WHERE date = $1", date).Scan(&note.ID, &note.Date, pq.Array(&note.Tag), &note.Text)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// GetNoteByTag возвращает заметку по тегу
func GetNoteByTag(tag string) ([]models.Note, error) {
	var notes []models.Note
	rows, err := db.Query("SELECT id, date, tag, text FROM notes WHERE $1 = ANY(tag)", tag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		// Сканируем результаты в структуру note
		if err := rows.Scan(&note.ID, &note.Date, pq.Array(&note.Tag), &note.Text); err != nil {
			return nil, err
		}
		notes = append(notes, note) // Добавляем найденную заметку в слайс
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil // Возвращаем слайс заметок
}
