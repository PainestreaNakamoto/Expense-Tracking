package repository

type Account struct {
	AccountID      int     `db:"account_id"`
	Title          string  `db:"title"`
	Description    string  `db:"description"`
	OverAllBalance float32 `db:"overall_balance"`
}

type AccountRepository interface {
	CreateAccount(account Account) (*Account, error)
	InfoAccount(account_id int) (*Account, error)
	Transactions(account_id int) ([]ExpenseTracking, error)
	MakeDeposit(ExpenseTracking) (*ExpenseTracking, error)
	MakeWithDrawal(ExpenseTracking) (*ExpenseTracking, error)
}
