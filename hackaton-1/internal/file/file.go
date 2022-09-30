package file

import (
	"bufio"
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	tickets := []service.Ticket{}

	file, err := os.Open(f.Path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ticket := service.Ticket{}
		ticket.InitFromCSVRow(fileScanner.Text())
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (f *File) Write(ticket service.Ticket) error {
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	csvRow := ticket.ToCsvRow()

	_, err = file.WriteString(csvRow)

	if err != nil {
		return err
	}

	return nil
}
