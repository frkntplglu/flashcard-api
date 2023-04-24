package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/frkntplglu/flashcard-api/internal/card"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "flashcard"
)

func main() {
	r := gin.Default()

	// Db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	fmt.Println("Connected!")

	// Repositories
	cardRepository := card.NewCardRepository(db)

	// Handlers
	cardHandler := card.NewCardHandler(cardRepository)
	cardHandler.SetRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
