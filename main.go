package main

import "fmt"

type Account struct {
	Owner    string
	Balance  float64
	password string
}

func main() {

	var accounts []Account

	for { // Createed a cycle for an infinity menu
		//Typing the menu
		fmt.Println("\n--- Bank Menu ---")
		fmt.Println("1. Create an account")
		fmt.Println("2. Show all accounts")
		fmt.Println("3. Deposite cash to your account")
		fmt.Println("4. Withdraw cash from your account")
		fmt.Println("5. Transfer money from one account to another")
		fmt.Println("6. Exit")
		fmt.Print("What can I do for You?")
		//awaiting the user's choice

		var choice int
		fmt.Scan(&choice) //reading the number, the user write

		//analyzing the choice
		switch choice {
		case 1:

			var name string
			fmt.Println("Creating the new account...")
			fmt.Println("Write your full name down")
			fmt.Scan(&name)
			acc := Account{Owner: name, Balance: 0} //created the new account
			accounts = append(accounts, acc)        // added the created account to the slice

			fmt.Printf("Account for %s successfuly created! Balance %2f$\n", name, acc.Balance) // displaying the confirmation massage to the user

		case 2:
			if len(accounts) == 0 {
				fmt.Println("No accounts has been created yet")
			} else {
				fmt.Println("\n The list of the accounts")
				for i, account := range accounts {
					fmt.Printf("%d. %s - Balance: %2f$\n", i+1, account.Owner, account.Balance)
				}
			}

		case 3:
			if len(accounts) == 0 {
				fmt.Println("No accounts has been created yet")
				break
			}
			var name string
			fmt.Printf("What account do you want to deposite in? \n")
			fmt.Scan(&name)
			//searching for the name
			found := false
			for i := 0; i < len(accounts); i++ {
				if accounts[i].Owner == name {
					//if the account is found asking or the amount
					var amount float64
					fmt.Printf("Write the amount of money to deposite: \n")
					fmt.Scan(&amount)

					// checking if the number is positive
					if amount <= 0 {
						fmt.Println("The number shall be more than 0")
						break
					}
					//increasing the balance
					accounts[i].Balance += amount
					fmt.Printf("Balance for the account '%s' increased by %.2f&. Current balance: %2.f$\n", name, amount, accounts[i].Balance)

					found = true
					break
				}
			}
			if !found {
				fmt.Println("There is no such account on the list")
			}
		case 4:
			if len(accounts) == 0 {
				fmt.Println("No accounts has been created yet")
				break
			}
			var name string
			fmt.Printf("What account do you want to withdraw from? \n")
			fmt.Scan(&name)
			//searching for the name
			found := false
			for i := 0; i < len(accounts); i++ {
				if accounts[i].Owner == name {
					found = true // if we found the account name found will be true
					//if the account is found asking or the amount
					var amount float64
					fmt.Printf("Write the amount of money to withdraw: \n")
					fmt.Scan(&amount)

					// checking if the number is positive or is there enough money to withdraw
					if amount <= 0 {
						fmt.Println("The number shall be more than 0")
						break
					} else if amount > accounts[i].Balance {
						fmt.Println("Not enough money on the account!")
						break
					} else {
						//decreasing the balance
						accounts[i].Balance -= amount
						fmt.Printf("Balance for the account '%s' decreased by %.2f$ Current balance: %2.f$\n", name, amount, accounts[i].Balance)
					}
					break
				}
			}
			if !found { //if we still don't found the account it will tell us
				fmt.Println("There is no such account on the list!!!")
			}

		case 5:
			if len(accounts) < 2 { // checking if there is 2 accounts to start a transfer
				fmt.Println("There shall be atleast 2 accounts for transfer")
				break
			}
			//getting the sender and the receiver account names
			var nameSender, nameReciever string

			fmt.Print("Wrtite the name of the account you want to transfer from! \n")
			fmt.Scan(&nameSender)
			fmt.Print("Wrtite the name of the account you want to transfer to! \n")
			fmt.Scan(&nameReciever)
			fmt.Print("Print the amount of money you want to transfer \n")
			//getting the amount of the money we want to transfer
			var amount float64
			fmt.Scan(&amount)
			// it is need to check the presence of the accounts on the list
			senderIndex := -1
			receiverIndex := -1

			for i, account := range accounts {
				// in this cycle we asure that there is such names on the list, and if so we mark them as indexes on the list that we have
				if account.Owner == nameSender {
					senderIndex = i
				}

				if account.Owner == nameReciever {
					receiverIndex = i
				}

			} //then we check for mistakes and errors with the givven datas such as sender name, receiver name and the amount of money to transfer
			if senderIndex == -1 {
				fmt.Print("The account you want to transfer from has not found!!!")
				break
			}

			if receiverIndex == -1 {
				fmt.Print("The account you want to transfer to has not found!!!")
				break
			}

			if senderIndex == receiverIndex {
				fmt.Print("You can not transfer money FROM and TO the same account!!!")
				break
			}

			if amount <= 0 {
				fmt.Print("The number of transfer should be more than 0 !!!")
				break
			}

			if accounts[senderIndex].Balance < amount {
				fmt.Print("There is no enough money on the account you want to transfer from!!!")
				break
			}
			//here we - the amount from sender balance and + the amount to the receiver balance
			accounts[senderIndex].Balance -= amount
			accounts[receiverIndex].Balance += amount
			// showing the success massage and the new balances for each account after the transfer
			fmt.Printf("Transfer by %.2f$ from account '%s' --to--> account '%s' has completed successfuly! \n", amount, accounts[senderIndex].Owner, accounts[receiverIndex].Owner)
			fmt.Printf("New balances \n %s: %.2f$ \n %s: %.2f$", accounts[senderIndex].Owner, accounts[senderIndex].Balance, accounts[receiverIndex].Owner, accounts[receiverIndex].Balance)

		case 6:
			fmt.Println("Thank you for bankig with us!")
			return //it is used to close the program

		default:
			fmt.Println("Wrong number, Try again!")

		}
	}
}
