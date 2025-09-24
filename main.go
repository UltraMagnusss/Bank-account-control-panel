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
	acc1 := Account{Owner: "Ahmad", Balance: 2500}
	acc2 := Account{Owner: "Muhammad", Balance: 1500}

	fmt.Println("Availible balances:")
	fmt.Printf("%s: %.2f\n", acc1.Owner, acc1.Balance)
	fmt.Printf("%s: %.2f\n", acc2.Owner, acc2.Balance)

	acc1.Deposit(250)
	fmt.Printf("After the deposit %s: %2f\n", acc1.Owner, acc1.Balance)

	success := acc1.Withdraw(350)
	if success {
		fmt.Println("Withdraw has completed successfuly")
	} else {
		fmt.Println("Not enough money on the account")
	}

	fmt.Printf("After withdraw %s: %2f\n", acc1.Owner, acc1.Balance)

	success = Transfer(&acc1, &acc2, 350)
	if success {
		fmt.Println("Transfer has completed successfuly")
	} else {
		fmt.Println("Not enough money on account")
	}

	fmt.Println("After the transfer:")
	fmt.Printf("%s: %.2f\n", acc1.Owner, acc1.Balance)
	fmt.Printf("%s: %.2f\n", acc2.Owner, acc2.Balance)
}
