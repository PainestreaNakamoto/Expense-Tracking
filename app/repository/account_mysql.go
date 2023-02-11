package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	db *sqlx.DB
}

func InitializeAccountRepositroyMySQL(db *sqlx.DB) AccountRepository {
	return accountRepository{db: db}
}
func (self accountRepository) CreateAccount(account Account) (*Account, error) {

	account_in_db := Account{}
	err := self.db.Get(&account_in_db, "select account_id from accounts where account_id=?", account.AccountID)

	if err == sql.ErrNoRows {
		query := "insert into accounts (account_id, title, description, overall_balance) values (?,?,?,?)"
		_, err = self.db.Exec(query, account.AccountID, account.Title, account.Description, 0)
		if err != nil {
			return nil, err
		}

		return &account, nil

	} else {
		return nil, errors.New("Can't Open Account or The account_id is already exits.")
	}

}
func (self accountRepository) InfoAccount(account_id int) (*Account, error) {
	query := "select account_id, title, description, overall_balance from accounts where account_id=?"
	account_info := Account{}
	err := self.db.Get(&account_info, query, account_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("The account_id %v is not found", account_id))
		}
		return nil, err
	}
	return &account_info, nil
}

func (self accountRepository) Transactions(account_id int) ([]ExpenseTracking, error) {
	query := "select id, date_time, title, account_id, classification, income, expense, overall_balance from expenses_tracking where account_id=?"
	account_transaction_lists := []ExpenseTracking{}
	err := self.db.Select(&account_transaction_lists, query, account_id)
	if err != nil {
		return nil, err
	}
	return account_transaction_lists, nil
}

func (self accountRepository) MakeDeposit(expense_data ExpenseTracking) (*ExpenseTracking, error) {
	if expense_data.Income > 0 && expense_data.Expense <= 0 {
		last_record := ExpenseTracking{}
		query := "select overall_balance from expenses_tracking order by id desc limit 1"
		err := self.db.Get(&last_record, query)
		if err == sql.ErrNoRows {
			// Not do anything
		} else if err != nil {
			return nil, err
		}

		query = "insert into expenses_tracking ( date_time, title, account_id, classification, income, expense, overall_balance) values (?,?,?,?,?,?,?)"
		expense_data.OverAllBalance = last_record.OverAllBalance + expense_data.Income
		result, err := self.db.Exec(query, expense_data.DateTime, expense_data.Title, expense_data.AccountID, expense_data.Classification, expense_data.Income, expense_data.Expense, expense_data.OverAllBalance)
		_ = result

		if err != nil {
			return nil, err
		}
		query = "select overall_balance from accounts where account_id=?"
		account_data := Account{}
		err = self.db.Get(&account_data, query, expense_data.AccountID)
		fmt.Println(account_data)
		if err != nil {
			return nil, err
		}
		query = "update accounts set overall_balance=? WHERE account_id=?"
		account_data.OverAllBalance += expense_data.Income
		fmt.Println(account_data.OverAllBalance)
		result, err = self.db.Exec(query, account_data.OverAllBalance, expense_data.AccountID)
		_ = result
		if err != nil {
			return nil, err
		}
		return &expense_data, nil

	}
	return nil, errors.New("Can't deposit money")

}

func (self accountRepository) MakeWithDrawal(expense_data ExpenseTracking) (*ExpenseTracking, error) {
	if expense_data.Expense > 0 && expense_data.Income <= 0 {
		last_record := ExpenseTracking{}
		query := "select overall_balance from expenses_tracking order by id desc limit 1"
		err := self.db.Get(&last_record, query)
		if err != nil {
			return nil, err
		}

		query = "insert into expenses_tracking ( date_time, title, account_id, classification, income, expense, overall_balance) values (?,?,?,?,?,?,?)"
		expense_data.OverAllBalance = last_record.OverAllBalance - expense_data.Expense
		result, err := self.db.Exec(query, expense_data.DateTime, expense_data.Title, expense_data.AccountID, expense_data.Classification, expense_data.Income, expense_data.Expense, expense_data.OverAllBalance)
		_ = result

		if err != nil {
			return nil, err
		}

		return &expense_data, nil
	}
	return nil, errors.New("Can't withdraw money")
}
