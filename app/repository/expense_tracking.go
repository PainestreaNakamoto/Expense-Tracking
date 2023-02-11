package repository

type ExpenseTracking struct {
	ID             int     `db:"id"`
	DateTime       string  `db:"date_time"`
	Title          string  `db:"title"`
	AccountID      string  `db:"account_id"`
	Classification string  `db:"classification"`
	Income         float32 `db:"income"`
	Expense        float32 `db:"expense"`
	OverAllBalance float32 `db:"overall_balance"`
}

type ExpenseTrackingRepository interface {
	NextIdentity(ExpenseTracking) (*ExpenseTracking, error)
	GetAll() ([]ExpenseTracking, error)
}
