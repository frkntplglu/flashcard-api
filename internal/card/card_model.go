package card

import "time"

type Card struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Status    int        `json:"status"`
}
