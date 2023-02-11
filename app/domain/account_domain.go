package domain

import "github.com/PainestreaNakamoto/Expense-Tracking/repository"

type accountDomain struct {
	repo repository.ExpenseTracking
}

func InitializeAccountDomain(repo repository.ExpenseTracking) AccountInterface {
	return accountDomain{repo: repo}
}

func (self accountDomain) OpenNewAccount(account AccountEntity) (string, error) {
	return "", nil
}

func (self accountDomain) AccountInfo(account_id int) (string, error) {
	return "", nil
}
func (self accountDomain) Transaction(account_id int) (string, error) {
	return "", nil
}
func (self accountDomain) Deposit(account_id int, expense_tracking ExpenseTrackingEntity) (string, error) {
	return "", nil
}
func (self accountDomain) WithDraw(account_id int, expense_tracking ExpenseTrackingEntity) (string, error) {
	return "", nil
}
