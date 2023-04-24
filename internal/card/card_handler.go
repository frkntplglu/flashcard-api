package card

import "github.com/gin-gonic/gin"

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

func (h *CardHandler) AddCards(c *gin.Context) {
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

func (h *CardHandler) SetRoutes(r *gin.Engine) {
	r.GET("/cards", h.GetCards)
	r.POST("/cards", h.AddCards)
}
