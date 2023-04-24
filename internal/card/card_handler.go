package card

import (
	"net/http"
	"strconv"

	"github.com/frkntplglu/flashcard-api/internal/models"
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
		c.JSON(http.StatusNotFound, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusNotFound,
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Success: true,
		Data:    cards,
	})
}

func (h *CardHandler) GetCardById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: "Invalid id",
				Code:    http.StatusBadRequest,
			},
		})
		return
	}
	card, err := h.CardRepository.GetCardById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusNotFound,
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Success: true,
		Data:    card,
	})
}

func (h *CardHandler) AddCard(c *gin.Context) {
	var card CardBody

	if c.ShouldBindJSON(&card) != nil {
		c.JSON(http.StatusBadRequest, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: "Invalid body",
				Code:    http.StatusBadRequest,
			},
		})
		return
	}
	err := h.CardRepository.AddCard(card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			},
		})
		return
	}
	c.JSON(http.StatusCreated, models.SuccessResponse{
		Success: true,
		Data:    card,
	})
}

func (h *CardHandler) UpdateCard(c *gin.Context) {
	var cardBody CardFields
	if c.ShouldBindJSON(&cardBody) != nil {
		c.JSON(http.StatusBadRequest, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: "Invalid body",
				Code:    http.StatusBadRequest,
			},
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: "Invalid id",
				Code:    http.StatusBadRequest,
			},
		})
		return
	}
	existingCard, err := h.CardRepository.GetCardById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusNotFound,
			},
		})
		return
	}

	err = h.CardRepository.UpdateCard(existingCard.Id, existingCard, cardBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Success: true,
		Data:    existingCard,
	})
}

func (h *CardHandler) DeleteCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: "Invalid id",
				Code:    http.StatusBadRequest,
			},
		})
		return
	}
	err = h.CardRepository.DeleteCard(id)

	if err != nil {
		c.JSON(http.StatusNotFound, models.FailedResponse{
			Success: false,
			Error: models.Error{
				Message: err.Error(),
				Code:    http.StatusNotFound,
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Success: true,
		Data:    "Card has been deleted successfully",
	})
}

func (h *CardHandler) SetRoutes(r *gin.Engine) {
	r.GET("/cards", h.GetCards)
	r.GET("/cards/:id", h.GetCardById)
	r.POST("/cards", h.AddCard)
	r.PUT("/cards/:id", h.UpdateCard)
	r.DELETE("/cards/:id", h.DeleteCard)
}
