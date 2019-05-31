package main

import (
	"environment.go.example/app/repository"
	"fmt"
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