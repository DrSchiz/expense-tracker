package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Descriprion string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}

func AddExpense(description string, amount int, file *os.File) {
	expenses := getExpenses(file)

	var expense = Expense{
		Id:          getMaxId(expenses) + 1,
		Descriprion: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	expenses = append(expenses, expense)

	writeFile(expenses, file)
	fmt.Printf("Запись успешно добавлена (ID: %d)\n", expense.Id)
}

func DeleteExpense(id int, file *os.File) {
	expenses := getExpenses(file)

	for i, expense := range expenses {
		if id == expense.Id {
			expenses = append(expenses[:i], expenses[i+1:]...)

			writeFile(expenses, file)
			fmt.Printf("Расход успешно удалён (ID: %d)\n", expense.Id)
			return
		}
	}
	fmt.Println("Нет такой записи")
}

func UpdateExpense(id int, description string, amount int, file *os.File) {
	expenses := getExpenses(file)

	for i, expense := range expenses {
		if id == expense.Id {
			if description != "" {
				expenses[i].Descriprion = description
			}
			if amount != 0 {
				expenses[i].Amount = amount
			}

			writeFile(expenses, file)
			fmt.Printf("Расход успешно изменён (ID: %d)\n", expense.Id)
			return
		}
	}
}

func ReadExpenses(file *os.File) {
	expenses := getExpenses(file)

	if len(expenses) == 0 {
		fmt.Println("У вас нет записей")
	}

	for _, expense := range expenses {
		fmt.Printf("ID: %d	Описание: %s	Цена: %d₽	Дата: %s\n", expense.Id, expense.Descriprion, expense.Amount, expense.Date.Format("2006-01-02"))
	}
}

func SummaryExpenses(month int, file *os.File) {
	expenses := getExpenses(file)

	var monthes []string = []string{
		"Январь",
		"Февраль",
		"Март",
		"Апрель",
		"Май",
		"Июнь",
		"Июль",
		"Август",
		"Сентябрь",
		"Октябрь",
		"Ноябрь",
		"Декабрь",
	}

	var amount int = 0

	if month != 0 {
		for _, expense := range expenses {
			if expense.Date.Month() == time.Month(month) {
				amount += expense.Amount
			}
		}
		fmt.Printf("Сумма расходов за %s: %d₽\n", monthes[month-1], amount)
	} else {
		for _, expense := range expenses {
			amount += expense.Amount
		}
		fmt.Printf("Сумма расходов за всё время: %d₽\n", amount)
	}

}

func getExpenses(file *os.File) []Expense {
	var expenses []Expense

	jsonData := readFile(file)

	json.Unmarshal(jsonData, &expenses)

	return expenses
}

func getMaxId(expenses []Expense) int {
	var ids []int
	var maxId = 0

	for _, expense := range expenses {
		ids = append(ids, expense.Id)
	}

	if len(expenses) != 0 {
		maxId = expenses[0].Id
		for _, id := range ids {
			if maxId < id {
				maxId = id
			}
		}
	}

	return maxId
}

func readFile(file *os.File) []byte {
	jsonData, err := os.ReadFile(file.Name())
	ErrorHandling(err)

	return jsonData
}

func writeFile(expenses []Expense, file *os.File) {
	jsonData, err := json.Marshal(expenses)
	ErrorHandling(err)

	os.WriteFile(file.Name(), jsonData, 1)
}

func ErrorHandling(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}
