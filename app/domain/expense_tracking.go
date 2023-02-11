package domain

type ExpenseTrackingEntity struct {
	ID             int
	DateTime       string
	Title          string
	AccountID      string
	Classification string
	Income         float32
	Expense        float32
	OverAllBalance float32
}

type ExpenseTrackingInterface interface {
	ExpenseLists() ([]ExpenseTrackingEntity, error) // it will show all history
}
