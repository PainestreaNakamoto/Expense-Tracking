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
		Title:          "Buy a water",
		AccountID:      "848654",
		Classification: "Daily use",
		Income:         0,
		Expense:        5,
	}
	data, err := account_domain.WithDraw(new_record)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}
