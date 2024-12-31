package main

import (
	"errors"
	"fmt"
)

type account struct {
	id               int
	name             string
	balance          float64
	transaction_hist []string
}

var accounts []account

const (
	option_deposit  = 1
	option_withdraw = 2
	option_balance  = 3
	option_history  = 4
	option_exit     = 5
)

func find_account_by_id(id int) (*account, error) {
	for i, acc := range accounts {
		if acc.id == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc, err := find_account_by_id(id)
	if err != nil {
		return err
	}
	acc.balance += amount
	acc.transaction_hist = append(acc.transaction_hist, fmt.Sprintf("Deposited: $%.2f", amount))
	return nil
}

func withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	acc, err := find_account_by_id(id)
	if err != nil {
		return err
	}
	if acc.balance < amount {
		return errors.New("insufficient balance")
	}
	acc.balance -= amount
	acc.transaction_hist = append(acc.transaction_hist, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

func view_balance(id int) (float64, error) {
	acc, err := find_account_by_id(id)
	if err != nil {
		return 0, err
	}
	return acc.balance, nil
}

func view_history(id int) ([]string, error) {
	acc, err := find_account_by_id(id)
	if err != nil {
		return nil, err
	}
	return acc.transaction_hist, nil
}

func main() {
	accounts = append(accounts, account{id: 1, name: "Keshav", balance: 5000})
	accounts = append(accounts, account{id: 2, name: "Bhavna", balance: 10000})

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice, acc_id int
		var amount float64
		fmt.Scan(&choice)

		if choice == option_exit {
			fmt.Println("Exiting the system. thank you!")
			break
		}

		fmt.Print("Enter Account ID: ")
		fmt.Scan(&acc_id)

		switch choice {
		case option_deposit:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := deposit(acc_id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case option_withdraw:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := withdraw(acc_id, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case option_balance:
			if balance, err := view_balance(acc_id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Current Balance: $%.2f\n", balance)
			}

		case option_history:
			if history, err := view_history(acc_id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, record := range history {
					fmt.Println(record)
				}
			}

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
