package main

import (
	"bufio"
	"fmt"
	"os"
	"personal-finance-tracker/finance"
	"strings"
	"time"
)

func readUserInput(reader *bufio.Reader) (finance.Transaction, error) {

	var transaction finance.Transaction

	fmt.Println("Enter date (YYYY-MM-DD):")
	dateStr, err := reader.ReadString('\n')
	if err != nil {
		return transaction, err
	}
	date, err := time.Parse("2006-01-02", strings.TrimSpace(dateStr))
	if err != nil {
		return transaction, fmt.Errorf("invalid date format. Try again")
	}

	fmt.Println("Enter description:")
	description, err := reader.ReadString('\n')
	if err != nil {
		return transaction, err
	}
	description = strings.TrimSpace(description)

	fmt.Println("Enter amount:")
	var amount float64
	fmt.Scan(&amount)

	fmt.Println("Enter type (income or expense):")
	typeStr, err := reader.ReadString('\n')
	if err != nil {
		return transaction, err
	}
	typeStr = strings.TrimSpace(typeStr)

	var transactionType finance.TransactionType
	if typeStr == "income" {
		transactionType = finance.Income
	} else if typeStr == "expense" {
		transactionType = finance.Expense
	} else {
		return transaction, fmt.Errorf("incorrect type, must be 'income' or 'expense'")
	}

	transaction = finance.Transaction{
		Date:        date,
		Description: description,
		Amount:      amount,
		Type:        transactionType,
	}
	return transaction, nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select an option: add, list, update, delete, quit")
		var option string
		fmt.Scan(&option)

		switch option {
		case "add":
			transaction, err := readUserInput(reader)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			finance.AddTransaction(transaction)

		case "list":
			finance.ListTransactions()

		case "update":
			fmt.Println("Enter the number of the transaction: ")
			var index int
			fmt.Scan(&index)

			finance.ShowTransaction(index)

			transaction, err := readUserInput(reader)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			finance.UpdateTransaction(index, transaction)

		case "quit":
			return
		default:
			fmt.Println("Unknown option")
		}
	}
}
