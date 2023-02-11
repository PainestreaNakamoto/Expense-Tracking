package domain

type AccountEntity struct {
	AccountID      int
	Title          string
	Description    string
	OverAllBalance float32
}

type AccountInterface interface {
	// For only open new account
	OpenNewAccount(account AccountEntity) (string, error)

	// For get infomation of account such as account info and transaction of account
	AccountInfo(account_id int) (string, error)
	Transaction(account_id int) (string, error)

	// For handle or interactive with account and record expense tracking
	Deposit(account_id int, expense_tracking ExpenseTrackingEntity) (string, error)
	WithDraw(account_id int, expense_tracking ExpenseTrackingEntity) (string, error)
}