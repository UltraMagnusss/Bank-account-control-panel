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

	var accounts []Account

	for { // Createed a cycle for an infinity menu
		//Typing the menu
		fmt.Println("\n--- Bank Menu ---")
		fmt.Println("1. Create account")
		fmt.Println("2. Exit")
		fmt.Print("What can I do for You?")
		//awaiting the user's choice

		var choice int
		fmt.Scan(&choice) //reading the number, the user write

		//analyzing the choice
		if choice == 1 {
			var name string
			fmt.Println("Creating the new account...")
			fmt.Println("Write your full name down")
			fmt.Scan(&name)
			acc := Account{Owner: name, Balance: 0} //created the new account
			accounts = append(accounts, acc)        // added the created account to the slice

			fmt.Printf("Account for %s successfuly created! Balance %2f\n", name, acc.Balance) // displaying the confirmation massage to the user
		} else if choice == 2 {
			fmt.Println("Good bye!")
			break //exit from the cycle, program shut down
		} else {
			fmt.Println("Wrong number, Try again!")
		}

	}
}
