package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID      int
	Name    string
	Balance float64
	History []string
}

var accounts []Account

func deposit(accID int, amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accID {
			accounts[i].Balance += amount
			accounts[i].History = append(accounts[i].History, fmt.Sprintf("Deposited: %.2f", amount))
			return nil
		}
	}
	return errors.New("Account not found")
}

func withdraw(accID int, amount float64) error {
	if amount <= 0 {
		return errors.New("Withdrawal amount must be greater than zero")
	}
	for i, acc := range accounts {
		if acc.ID == accID {
			if acc.Balance < amount {
				return errors.New("Insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].History = append(accounts[i].History, fmt.Sprintf("Withdrew: %.2f", amount))
			return nil
		}
	}
	return errors.New("Account not found")
}

func viewHistory(accID int) error {
	for _, acc := range accounts {
		if acc.ID == accID {
			fmt.Println("Transaction History for", acc.Name)
			for _, record := range acc.History {
				fmt.Println(record)
			}
			return nil
		}
	}
	return errors.New("Account not found")
}

func main() {
	accounts = append(accounts, Account{ID: 1, Name: "Siva", Balance: 15000})
	accounts = append(accounts, Account{ID: 2, Name: "Sid", Balance: 12000})

	const (
		DepositOption     = 1
		WithdrawOption    = 2
		ViewBalanceOption = 3
		ViewHistoryOption = 4
		ExitOption        = 5
	)

	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case DepositOption:
			var accID int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accID)
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			if err := deposit(accID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful")
			}
		case WithdrawOption:
			var accID int
			var amount float64
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accID)
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			if err := withdraw(accID, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful")
			}
		case ViewBalanceOption:
			var accID int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accID)
			for _, acc := range accounts {
				if acc.ID == accID {
					fmt.Printf("Account ID: %d, Name: %s, Balance: %.2f\n", acc.ID, acc.Name, acc.Balance)
				}
			}
		case ViewHistoryOption:
			var accID int
			fmt.Print("Enter account ID: ")
			fmt.Scan(&accID)
			if err := viewHistory(accID); err != nil {
				fmt.Println("Error:", err)
			}
		case ExitOption:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
