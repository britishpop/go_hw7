package transaction

import "sync"

type Transaction struct {
	UserId int64
	Sum    int64
	MCC    string
}

func MakeTransactions(userId int64) []Transaction {
	const usersCount = 100
	const transactionsCount = 100
	const transactionAmount = 1_00
	transactions := make([]Transaction, usersCount*transactionsCount)
	x := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5912",
	}
	y := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "5921",
	}
	z := Transaction{
		UserId: userId,
		Sum:    transactionAmount,
		MCC:    "4121",
	}
	a := Transaction{
		UserId: 129,
		Sum:    transactionAmount,
		MCC:    "4121",
	}
	for index := range transactions {
		switch index % 10 {
		case 0:
			transactions[index] = x
		case 3:
			transactions[index] = y
		case 5:
			transactions[index] = z
		default:
			transactions[index] = a
		}
	}
	return transactions
}

// первая функция - принимает на вход слайс транзакций и id владельца -
// возвращает map с категориями и тратами по ним
func SumByMCC(transactions []Transaction, userId int64) map[string]int64 {
	result := make(map[string]int64)
	for _, value := range transactions {
		if value.UserId == userId {
			result[TranslateMCC(value.MCC)] += value.Sum
		}
	}
	return result
}

// вторая функция - функция с mutex'ом, который защищает любые операции с map
// делит слайс транзакций на несколько кусков и в отдельных горутинах считает map'ы,
// после чего собирает всё в один большой map (вызывает простую функцию)
func MutexSumByMCC(transactions []Transaction, userId int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	partsCount := 10
	partsSize := len(transactions) / partsCount

	for i := 0; i < partsCount; i++ {
		wg.Add(1)
		part := transactions[i*partsSize : (i+1)*partsSize]
		go func() {
			m := SumByMCC(part, userId)
			mu.Lock()
			for k, v := range m {
				result[k] += v
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}

//третья функция - функция с каналами, делит слайс транзакций на несколько кусков
// и в отдельных горутинах считает map'ы по кускам,
// собирает всё в один большой map
func ChanSumByMCC(transactions []Transaction, userId int64) map[string]int64 {
	result := make(map[string]int64)
	ch := make(chan map[string]int64)

	partsCount := 10
	partsSize := len(transactions) / partsCount

	for i := 0; i < partsCount; i++ {
		part := transactions[i*partsSize : (i+1)*partsSize]

		go func(ch chan<- map[string]int64) {
			ch <- SumByMCC(part, userId)
		}(ch)
	}

	finished := 0
	for m := range ch {
		for k, v := range m {
			result[k] += v
		}
		finished++
		if finished == partsCount {
			break
		}
	}

	return result
}

// вторая функция - функция с mutex'ом, который защищает любые операции с map
// делит слайс транзакций на несколько кусков и в отдельных горутинах считает map'ы,
// после чего собирает всё в один большой map (НЕ вызывает простую функцию)
func MutexSumByMCC2(transactions []Transaction, userId int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	partsCount := 10
	partsSize := len(transactions) / partsCount

	for i := 0; i < partsCount; i++ {
		wg.Add(1)
		part := transactions[i*partsSize : (i+1)*partsSize]
		go func() {
			for _, value := range part {
				if value.UserId == userId {
					mu.Lock()
					result[TranslateMCC(value.MCC)] += value.Sum
					mu.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}
