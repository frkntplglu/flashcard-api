package card

import (
	"database/sql"
	"fmt"
	"time"
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

func (r *CardRepository) GetCardById(id int) (*Card, error) {
	selectQuery := `SELECT * FROM cards WHERE id=$1;`

	row := r.Db.QueryRow(selectQuery, id)
	var card Card
	err := row.Scan(&card.Title, &card.Content, &card.UpdatedAt, &card.Id, &card.CreatedAt, &card.IsLearned)

	if err != nil {
		return nil, err
	}

	return &card, nil

}

func (r *CardRepository) AddCard(cardBody CardBody) error {
	insertQuery := `insert into "cards"("title", "content") values($1, $2)`
	_, err := r.Db.Exec(insertQuery, cardBody.Title, cardBody.Content)
	if err != nil {
		return err
	}

	return nil
}

func (r *CardRepository) UpdateCard(id int, cardBody CardBody) error {
	updateQuery := `update cards set title=$1, content=$2, updated_at=$3 where id=$4`
	_, err := r.Db.Exec(updateQuery, cardBody.Title, cardBody.Content, time.Now(), id)
	fmt.Println("err :", err)
	if err != nil {
		return err
	}

	return nil
}
