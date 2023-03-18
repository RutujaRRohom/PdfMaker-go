package service

import (
	"context"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit"
	//"github.com/RohomRutuja/Employee/PDF/domain"
)

type Employe interface {
	GenerateEmployee(ctx context.Context, numEmployee int) (employee []Employee, err error)
}

func NewFunc() Employe {
	return &Employee{}
}

func (emp *Employee) GenerateEmployee(ctx context.Context, numEmployee int) (employee []Employee, err error) {

	employee = []Employee{}

	for i := 0; i < numEmployee; i++ {

		employees := Employee{
			EmployeeID:  rand.Intn(1000),
			FirstName:   gofakeit.FirstName(),
			LastName:    gofakeit.LastName(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
		}
		employee = append(employee, employees)

	}
	err = GeneratePDF(employee)
	if err != nil {
		log.Println("error in generating pdf")
		return
	}

	return employee, nil
}
