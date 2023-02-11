package repository

import (
	"github.com/jmoiron/sqlx"
)

type expenseTrackingRepositoryMySQL struct {
	db *sqlx.DB
}

func InitializeExpenseTrackingRepositroyMySQL(db *sqlx.DB) ExpenseTrackingRepository {
	return expenseTrackingRepositoryMySQL{db: db}
}

func (self expenseTrackingRepositoryMySQL) NextIdentity(expense_item ExpenseTracking) (*ExpenseTracking, error) {
	return nil, nil
}

func (self expenseTrackingRepositoryMySQL) GetAll() ([]ExpenseTracking, error) {
	expenses_list := []ExpenseTracking{}
	query := "select id, date_time, title, account_id, classification, income, expense, overall_balance from expenses_tracking"
	err := self.db.Select(&expenses_list, query)

	if err != nil {
		return nil, err
	}

	return expenses_list, nil
}
