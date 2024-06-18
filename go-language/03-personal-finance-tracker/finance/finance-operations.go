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
	fmt.Printf(`
	Date: %s;
	Description: %s;
	Amount: $ %.2f;
	Type: %s.
	`,
		transaction.Date.Format("2006-01-02"),
		transaction.Description,
		transaction.Amount,
		transaction.Type.ToString(),
	)
}

func AddTransaction(transaction Transaction) {

	if transaction.Type == Expense && transaction.Amount > 0 {
		transaction.Amount = -1.0 * transaction.Amount
	}
	transactions = append(transactions, transaction)

	fmt.Printf("\nAdded new transaction with details below:\n")
	printTransactionFormattedDetails(transaction)
	fmt.Println()
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

func UpdateTransaction(index int, transaction Transaction) {
	index = index - 1
	if index < 0 || index > len(transactions)-1 {
		fmt.Println("Invalid index. Try again")
		return
	}

	if transaction.Type == Expense && transaction.Amount > 0 {
		transaction.Amount = -1.0 * transaction.Amount
	}

	transactions[index] = transaction
	fmt.Printf("\nUpdated transaction:\n")
	printTransactionFormattedDetails(transactions[index])
	fmt.Println()
}

func ShowTransaction(index int) {

	index = index - 1
	if index < 0 || index > len(transactions)-1 {
		fmt.Println("Invalid index. Try again!")
		return
	}

	fmt.Printf("\nTransaction details:\n")
	printTransactionFormattedDetails(transactions[index])
	fmt.Println()
}
