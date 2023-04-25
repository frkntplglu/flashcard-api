package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/frkntplglu/flashcard-api/internal/card"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	r := gin.Default()
	fmt.Println(viper.Get("DB_HOST"))
	fmt.Println(viper.Get("DB_PORT"))
	// Db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.Get("DB_HOST"), viper.GetInt("DB_PORT"), viper.Get("DB_USER"), viper.Get("DB_PASSWORD"), viper.Get("DB_NAME"))
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

	r.Run(fmt.Sprintf(":%v", viper.Get("PORT")))
}
