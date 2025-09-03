package main

import (
	functions "expense-tracker/functions"
	"flag"
	"fmt"
	"os"
)

func main() {

	filepath := "expenses.json"
	if !functions.CheckFileExist(filepath) {
		os.Create(filepath)
	}

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	action := os.Args[1]

	switch action {
	case "add":
		add := flag.NewFlagSet("add", flag.ExitOnError)

		description := add.String("description", "", "Описание расхода")
		amount := add.Int("amount", 0, "Стоимость расхода")

		add.Parse(os.Args[2:])

		functions.AddExpense(*description, *amount, file)
	case "delete":
		delete := flag.NewFlagSet("delete", flag.ExitOnError)

		id := delete.Int("id", 0, "ID удаляемого расхода")

		delete.Parse(os.Args[2:])

		functions.DeleteExpense(*id, file)
	case "update":
		update := flag.NewFlagSet("update", flag.ExitOnError)

		id := update.Int("id", 0, "ID изменяемого расхода")
		description := update.String("description", "", "Описание изменяемого расхода")
		amount := update.Int("amount", 0, "Стоимость изменяемого расхода")

		update.Parse(os.Args[2:])

		functions.UpdateExpense(*id, *description, *amount, file)
	case "list":
		functions.ReadExpenses(file)
	case "summary":
		summary := flag.NewFlagSet("summary", flag.ExitOnError)

		month := summary.Int("month", 0, "Месяц расходов")

		summary.Parse(os.Args[2:])

		functions.SummaryExpenses(*month, file)
	}
}
