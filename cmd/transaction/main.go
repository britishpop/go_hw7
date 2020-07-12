package main

import (
	"fmt"

	"go_hw7.1/pkg/transaction"
)

func main() {
	transactions := transaction.MakeTransactions(1)

	// первая функция
	sum := transaction.SumByMCC(transactions, 1)
	for k, v := range sum {
		fmt.Printf("Sum transaction in %s: %d \r\n", k, v)
	}
}
