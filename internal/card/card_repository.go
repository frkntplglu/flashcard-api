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

func (r *CardRepository) UpdateCard(id int, card *Card, fields CardFields) error {

	updateQuery := `UPDATE cards SET `
	if fields.Title != "" {
		updateQuery += `title='` + fields.Title + `', `
	}
	if fields.Content != "" {
		updateQuery += `content='` + fields.Content + `', `
	}
	if *fields.IsLearned {
		updateQuery += `is_learned='` + "true" + `', `
	} else {
		updateQuery += `is_learned='` + "false" + `', `
	}
	updateQuery += `updated_at=$1 WHERE id=$2 RETURNING *;`
	fmt.Println(updateQuery)
	row := r.Db.QueryRow(updateQuery, time.Now(), id)

	if row.Err() != nil {
		return row.Err()
	}
	return row.Scan(&card.Title, &card.Content, &card.UpdatedAt, &card.Id, &card.CreatedAt, &card.IsLearned)
}

func (r *CardRepository) DeleteCard(id int) error {
	deleteQuery := `DELETE FROM cards WHERE id=$1`
	_, err := r.Db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
