package main

import (
	"fmt"
	"go.example.environment/app/repository"
)

func main() {

	fmt.Println("starting..")
	customers, err := repository.GetCustomers()

	if err != nil {
		fmt.Println("ERRO!")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Finding customers..")

		for _, customer := range customers {
			fmt.Println("")
			fmt.Println(customer.ID)
			fmt.Println(customer.Name)
			fmt.Println(customer.Email)
		}
	}
}