// определения структур и моделей данных.
package models

type Note struct {
	ID   string   `json:"id"`
	Date string   `json:"date"`
	Tag  []string `json:"tag"`
	Text string   `json:"text"`
}
