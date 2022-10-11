package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service tickets.Service
}

func NewHandler(s tickets.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (s *Handler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Handler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
