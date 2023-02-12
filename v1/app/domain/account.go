package domain

type AccountEntity struct {
	AccountID      int
	Title          string
	Description    string
	OverAllBalance float32
}

type AccountInterface interface {
	// For only open new account
	OpenNewAccount(account AccountEntity) (*AccountEntity, error)

	// For get infomation of account such as account info and transaction of account
	AccountInfo(account_id int) (*AccountEntity, error)
	Transaction(account_id int) ([]ExpenseTrackingEntity, error)

	// For handle or interactive with account and record expense tracking
	Deposit(ExpenseTrackingEntity) (*ExpenseTrackingEntity, error)
	WithDraw(ExpenseTrackingEntity) (*ExpenseTrackingEntity, error)
}
