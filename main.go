package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	if a.Balance < amount {
		return false
	}
	a.Balance -= amount
	return true
}

func Transfer(from *Account, to *Account, amount float64) bool {
	success := from.Withdraw(amount)
	if success {
		to.Deposit(amount)
		return true
	}
	return false
}

func main() {
	for { // Createed a cycle for an infinity menu
		//Typing the menu
		fmt.Println("Menu:")
		fmt.Println("1. Create account")
		fmt.Println("2. Exit")

		//awaiting the user's choice
		var choice int
		fmt.Println("Choose the action: ")
		fmt.Scan(&choice) //reading the number, the user write

		//analyzing the choice
		if choice == 1 {
			fmt.Println("Creating the new account...") //here will be code writen after a while
		} else if choice == 2 {
			fmt.Println("Good bye!")
			break //exit from the cycle, program shut down
		} else {
			fmt.Println("Wrong number, Try again!")
		}

	}
}
