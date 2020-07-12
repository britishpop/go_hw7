package transaction

type Transaction struct {
	Id     int64
	UserId int64
	Type   string
	Sum    int64
	Status string
	MCC    string
	Date   int64
}

func Sum(transactions []Transaction) int64 {
	sum := int64(0)
	for _, t := range transactions {
		sum += t.Sum
	}
	return sum
}

func SumByMCC(transactions []Transaction, userId int64) map[string]int64 {
	transByMcc := make(map[string][]Transaction)
	for _, value := range transactions {
		if value.UserId == userId {
			transByMcc[TranslateMCC(value.MCC)] = append(transByMcc[TranslateMCC(value.MCC)], value)
		}
	}
	result := make(map[string]int64)
	for key, value := range transByMcc {
		result[key] = Sum(value)
	}
	return result
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
		switch index % 13 {
		case 0:
			transactions[index] = x
		case 6:
			transactions[index] = y
		case 7:
			transactions[index] = z
		default:
			transactions[index] = a
		}
	}
	return transactions
}
