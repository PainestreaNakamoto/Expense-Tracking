package domain

import (
	"github.com/PainestreaNakamoto/Expense-Tracking/repository"
)

type expenseTrackingDomain struct {
	repo repository.ExpenseTrackingRepository
}

func InitializeExpenseTrackingDomain(repo repository.ExpenseTrackingRepository) ExpenseTrackingInterface {
	return expenseTrackingDomain{repo: repo}
}

func (self expenseTrackingDomain) ExpenseLists() ([]ExpenseTrackingEntity, error) {
	expenses_tracking_db, err := self.repo.GetAll()
	if err != nil {
		return nil, err
	}
	expense_tracking_lists := []ExpenseTrackingEntity{}
	for _, item := range expenses_tracking_db {
		expense_model := ExpenseTrackingEntity{
			ID:             item.ID,
			DateTime:       item.DateTime,
			Title:          item.Title,
			Classification: item.Classification,
			Income:         item.Income,
			Expense:        item.Expense,
			OverAllBalance: item.OverAllBalance,
		}
		expense_tracking_lists = append(expense_tracking_lists, expense_model)
	}

	return expense_tracking_lists, err
}
