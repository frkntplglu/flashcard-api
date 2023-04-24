package card

import "database/sql"

type CardRepository struct {
	Db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (r *CardRepository) GetCards() ([]Card, error) {
	return nil, nil
}
