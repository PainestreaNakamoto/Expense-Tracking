package main

import (
	"fmt"

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
	data, err := expense_domain.ExpenseLists()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
