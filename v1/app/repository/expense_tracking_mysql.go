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
	query := "insert into expenses_tracking (date_time, title, account_id, classification, income, expense, overall_balance) values (?,?,?,?,?,?,?)"
	result, err := self.db.Exec(query, expense_item.DateTime, expense_item.Title, expense_item.AccountID, expense_item.Classification, expense_item.Income, expense_item.Expense, expense_item.OverAllBalance)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	expense_item.ID = int(id)

	return &expense_item, nil
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
