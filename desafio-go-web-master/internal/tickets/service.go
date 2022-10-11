package tickets

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(*gin.Context, string) (int, error)
	AverageDestination(*gin.Context, string) (float64, error)
}

type service struct {
	repository Repository
}

func NewTicketService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetTotalTickets(c *gin.Context, dest string) (int, error) {
	ticketsByDest, err := s.repository.GetTicketByDestination(c, dest)

	if err != nil {
		return 0, err
	}

	return len(ticketsByDest), nil
}

func (s *service) AverageDestination(c *gin.Context, dest string) (float64, error) {
	allTicketsCount, err := s.repository.GetAll(c)

	if err != nil {
		return 0, err
	}

	ticketsCountByDest, err := s.GetTotalTickets(c, dest)

	if err != nil {
		return 0, err
	}

	avg := float64(ticketsCountByDest / len(allTicketsCount))

	return avg, nil
}
