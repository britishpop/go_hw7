package transaction

const notCategory = "Категория не указана"

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5921": "Рестораны",
		"5411": "Супермаркеты",
		"5912": "Аптеки",
		"3514": "Отели",
		"4121": "Такси",
	}
	value, ok := mcc[code]

	if ok {
		return value
	}
	return notCategory
}
