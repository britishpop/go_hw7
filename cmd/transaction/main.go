package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"

	"go_hw7.1/pkg/transaction"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	transactions := transaction.MakeTransactions(1)

	// // первая функция
	// fmt.Println("Результаты первой функции")
	// sum := transaction.SumByMCC(transactions, 1)
	// for k, v := range sum {
	// 	fmt.Printf("Sum transaction in %s: %d \r\n", k, v)
	// }
	// вторая функция
	fmt.Println("Результаты второй функции")
	sum2 := transaction.MutexSumByMCC(transactions, 1)
	for k, v := range sum2 {
		fmt.Printf("Sum transaction in %s: %d \r\n", k, v)
	}
	// // третья функция
	// fmt.Println("Результаты третьей функции")
	// sum3 := transaction.ChanSumByMCC(transactions, 1)
	// for k, v := range sum3 {
	// 	fmt.Printf("Sum transaction in %s: %d \r\n", k, v)
	// }
	// // четвертая функция
	// fmt.Println("Результаты четвертой функции")
	// sum4 := transaction.MutexSumByMCC2(transactions, 1)
	// for k, v := range sum4 {
	// 	fmt.Printf("Sum transaction in %s: %d \r\n", k, v)
	// }
}
