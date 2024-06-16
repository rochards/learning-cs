package main

import (
	"bufio"
	"fmt"
	"os"
	"personal-finance-tracker/finance"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select an option: add, list, complete, delete, quit")
		var option string
		fmt.Scan(&option)

		switch option {
		case "add":
			fmt.Println("Enter date (YYYY-MM-DD):")
			dateStr, _ := reader.ReadString('\n')
			date, err := time.Parse("2006-01-02", strings.TrimSpace(dateStr))
			if err != nil {
				fmt.Println("Invalid date format. Try again!")
				continue
			}

			fmt.Println("Enter description:")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Println("Enter amount")
			var amount float64
			fmt.Scan(&amount)

			fmt.Println("Enter type: income or expense")
			typeStr, _ := reader.ReadString('\n')
			typeStr = strings.TrimSpace(typeStr)

			var transactionType finance.TransactionType
			if typeStr == "income" {
				transactionType = finance.Income
			} else if typeStr == "expense" {
				transactionType = finance.Expense
			} else {
				fmt.Println("Incorrect type")
				continue
			}

			finance.AddTransaction(date, description, amount, transactionType)

		case "list":
			finance.ListTransactions()
		case "quit":
			return
		default:
			fmt.Println("Unknown option")
		}
	}
}
