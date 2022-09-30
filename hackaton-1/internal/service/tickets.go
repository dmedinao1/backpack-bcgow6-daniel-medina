package service

import (
	"fmt"
	"strconv"
	"strings"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	var ticket Ticket

	for _, currentTicket := range b.Tickets {
		if currentTicket.Id == id {
			ticket = currentTicket
			break
		}
	}

	if ticket.Id == 0 {
		return ticket, fmt.Errorf("error: Tickect con id %v no encontrado", id)
	}

	return ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	toUpdate, err := b.Read(id)

	if err != nil {
		return t, err
	}

	toUpdate.Names = t.Names
	toUpdate.Email = t.Email
	toUpdate.Destination = t.Destination
	toUpdate.Date = t.Date
	toUpdate.Price = t.Price

	return toUpdate, nil
}

func (b *bookings) Delete(id int) (int, error) {
	var indexToDelete int

	for i, currentTicket := range b.Tickets {
		if currentTicket.Id == id {
			indexToDelete = i
			break
		}
	}

	if indexToDelete == 0 {
		return 0, fmt.Errorf("error: Tickect con id %v no encontrado", id)
	}

	b.Tickets = append(b.Tickets[:indexToDelete], b.Tickets[indexToDelete+1:]...)

	return id, nil
}

/*Custom methods*/

func (t *Ticket) InitFromCSVRow(csvRow string) {
	tokens := strings.Split(csvRow, ",")

	if len(tokens) != 6 {
		panic("no hay suficientes columnas para obtener el ticket")
	}

	id, err := strconv.Atoi(tokens[0])

	if err != nil {
		panic(fmt.Sprintf("el id %v no es un número", tokens[0]))
	}

	price, err := strconv.Atoi(tokens[5])

	if err != nil {
		panic(fmt.Sprintf("el precio %v no es un número", tokens[5]))
	}

	t.Id = id
	t.Names = tokens[1]
	t.Email = tokens[2]
	t.Destination = tokens[3]
	t.Date = tokens[4]
	t.Price = price
}

func (t Ticket) ToCsvRow() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
}
