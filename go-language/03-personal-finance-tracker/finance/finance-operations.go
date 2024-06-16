package finance

import (
	"fmt"
	"time"
)

type TransactionType int

const Income TransactionType = 0
const Expense TransactionType = 1

func (t TransactionType) ToString() string {
	switch t {
	case Income:
		return "Income"
	case Expense:
		return "Expense"
	default:
		return "Unknown"
	}
}

type Transaction struct {
	Date        time.Time
	Description string
	Amount      float64
	Type        TransactionType
}

var transactions []Transaction

func printTransactionFormattedDetails(transaction Transaction) {
	fmt.Println("Date: ", transaction.Date.Format("2006-01-02"))
	fmt.Println("Description: ", transaction.Description)
	fmt.Printf("Amount: $ %.2f\n", transaction.Amount)
	fmt.Println("Type: ", transaction.Type.ToString())
}

func AddTransaction(date time.Time, description string, amount float64, transactionType TransactionType) {

	if transactionType == Expense {
		amount = -1.0 * amount
	}

	transaction := Transaction{Date: date, Description: description, Amount: amount, Type: transactionType}
	transactions = append(transactions, transaction)

	fmt.Println("Added new transaction. With details below")
	printTransactionFormattedDetails(transaction)
}

func ListTransactions() {

	if len(transactions) == 0 {
		fmt.Println("No transactions yet!")
		return
	}

	for i, transaction := range transactions {
		fmt.Printf("\n%d. == transaction ==\n", i+1)
		printTransactionFormattedDetails(transaction)
		fmt.Println()
	}
}
