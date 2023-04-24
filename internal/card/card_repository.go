package card

import (
	"database/sql"
	"fmt"
)

type CardRepository struct {
	Db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{Db: db}
}

func (r *CardRepository) GetCards() ([]Card, error) {
	rows, err := r.Db.Query("SELECT * FROM cards")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cards := make([]Card, 0)

	for rows.Next() {
		card := Card{}
		err = rows.Scan(&card.Title, &card.Content, &card.UpdatedAt, &card.Id, &card.CreatedAt, &card.IsLearned)
		if err != nil {
			fmt.Println("err :", err)
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *CardRepository) AddCard(cardBody CardBody) error {
	insertQuery := `insert into "cards"("title", "content") values($1, $2)`
	result, err := r.Db.Exec(insertQuery, cardBody.Title, cardBody.Content)
	fmt.Println("result :", result)
	if err != nil {
		return err
	}

	return nil
}
