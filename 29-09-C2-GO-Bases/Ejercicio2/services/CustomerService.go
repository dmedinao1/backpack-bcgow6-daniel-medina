package services

import (
	"fmt"

	"github.com/dmedinao1/backpack-bcgow6-daniel-medina/customers/models"
	"github.com/dmedinao1/backpack-bcgow6-daniel-medina/customers/repositories"
)

func SaveNewClient(
	FullName string,
	DNI string,
	PhoneNumber string,
	Adress string,
) (c *models.Customer, err error) {
	if FullName == "" {
		err = fmt.Errorf("error: El Nombre %v es inválido", FullName)
		return
	}

	if DNI == "" {
		err = fmt.Errorf("error: El DNI %v es inválido", DNI)
		return
	}

	if PhoneNumber == "" {
		err = fmt.Errorf("error: El Teléfono %v es inválido", PhoneNumber)
		return
	}

	if Adress == "" {
		err = fmt.Errorf("error: La Dirección %v es inválido", Adress)
		return
	}

	if userExistsByDNI(DNI) {
		err = fmt.Errorf("error: El DNI %v ya fue registrado", DNI)
		return
	}

	id := repositories.GenerateNextId()

	c = &models.Customer{
		Id:          id,
		FullName:    FullName,
		DNI:         DNI,
		PhoneNumber: PhoneNumber,
		Adress:      Adress,
	}

	repositories.SaveCustomer(c)

	return
}

func userExistsByDNI(dni string) bool {
	allCostumers := repositories.GetAllRegisteredCustomers()

	for _, customer := range *allCostumers {
		if customer.DNI == dni {
			return true
		}
	}

	return false
}
