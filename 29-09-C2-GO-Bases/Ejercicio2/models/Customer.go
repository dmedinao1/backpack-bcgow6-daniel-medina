package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Customer struct {
	Id          int
	FullName    string
	DNI         string
	PhoneNumber string
	Adress      string
}

func (c *Customer) FillFromString(raw string) error {
	fields := strings.Split(raw, ",")

	id, err := strconv.Atoi(fields[0])

	if err != nil {
		return err
	}

	c.Id = id
	c.FullName = fields[1]
	c.DNI = fields[2]
	c.PhoneNumber = fields[3]
	c.Adress = fields[4]

	return nil
}

func (c Customer) ToString() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v\n", c.Id, c.FullName, c.DNI, c.PhoneNumber, c.Adress)
}
