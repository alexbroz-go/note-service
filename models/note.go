// определения структур и моделей данных.
package models

type Note struct {
	ID   int      `json:"id"`
	Date string   `json:"date"`
	Tag  []string `json:"tag"`
	Text string   `json:"text"`
}
