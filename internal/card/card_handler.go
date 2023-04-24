package card

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	CardRepository *CardRepository
}

func NewCardHandler(cardRepository *CardRepository) *CardHandler {
	return &CardHandler{CardRepository: cardRepository}
}

func (h *CardHandler) GetCards(c *gin.Context) {
	cards, err := h.CardRepository.GetCards()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"cards": cards,
	})
}

func (h *CardHandler) GetCardById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}
	card, err := h.CardRepository.GetCardById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"card": card,
	})
}

func (h *CardHandler) AddCard(c *gin.Context) {
	var card CardBody
	if c.ShouldBindJSON(&card) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid body",
		})
		return
	}
	err := h.CardRepository.AddCard(card)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"card": card,
	})
}

func (h *CardHandler) UpdateCard(c *gin.Context) {
	var card CardBody
	if c.ShouldBindJSON(&card) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid body",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}
	existingCard, err := h.CardRepository.GetCardById(id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	err = h.CardRepository.UpdateCard(existingCard.Id, card)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"card": card,
	})
}

func (h *CardHandler) SetRoutes(r *gin.Engine) {
	r.GET("/cards", h.GetCards)
	r.GET("/cards/:id", h.GetCardById)
	r.POST("/cards", h.AddCard)
	r.PUT("/cards/:id", h.UpdateCard)
}
