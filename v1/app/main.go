package main

import (
	"github.com/PainestreaNakamoto/Expense-Tracking/domain"
	"github.com/PainestreaNakamoto/Expense-Tracking/handler"
	_ "github.com/PainestreaNakamoto/Expense-Tracking/handler"
	"github.com/PainestreaNakamoto/Expense-Tracking/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
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
	account_handler := handler.InitializeAccountHandler(account_domain)

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	account_group := v1.Group("/account")
	account_group.Get("/info/:account_id", account_handler.AccountInfomation)
	app.Listen(":8000")

}
