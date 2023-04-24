package card

import "time"

type Card struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	IsLearned bool       `json:"is_learned"`
}

type CardBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CardFields struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	IsLearned *bool  `json:"is_learned"`
}
