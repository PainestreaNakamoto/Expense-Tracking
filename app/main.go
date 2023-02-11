package main

import (
	"fmt"
	"time"

	"github.com/PainestreaNakamoto/Expense-Tracking/domain"
	"github.com/PainestreaNakamoto/Expense-Tracking/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/expense_tracking_db")
	if err != nil {
		panic(err)
	}
	expense_repository := repository.InitializeExpenseTrackingRepositroyMySQL(db)
	expense_domain := domain.InitializeExpenseTrackingDomain(expense_repository)
	_ = expense_domain

	account_repository := repository.InitializeAccountRepositroyMySQL(db)
	account_domain := domain.InitializeAccountDomain(account_repository)
	new_record := domain.ExpenseTrackingEntity{
		DateTime:       time.Now().Format("2006-01-2 15:04:05"),
		Title:          "From profit of business",
		AccountID:      "999999",
		Classification: "Wedding",
		Income:         10,
		Expense:        0,
	}
	data, err := account_domain.Deposit(new_record)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	data2, err := account_domain.AccountInfo(999999)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)

}
