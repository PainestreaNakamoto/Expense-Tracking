package domain

import (
	"github.com/PainestreaNakamoto/Expense-Tracking/repository"
)

type accountDomain struct {
	repo repository.AccountRepository
}

func InitializeAccountDomain(repo repository.AccountRepository) AccountInterface {
	return accountDomain{repo: repo}
}

func (self accountDomain) OpenNewAccount(account AccountEntity) (*AccountEntity, error) {
	new_account := repository.Account{
		AccountID:   account.AccountID,
		Title:       account.Title,
		Description: account.Description,
	}
	_, err := self.repo.CreateAccount(new_account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (self accountDomain) AccountInfo(account_id int) (*AccountEntity, error) {
	account_from_db, err := self.repo.InfoAccount(account_id)
	account_infomation := AccountEntity{
		AccountID:      account_from_db.AccountID,
		Title:          account_from_db.Title,
		Description:    account_from_db.Description,
		OverAllBalance: account_from_db.OverAllBalance,
	}
	if err != nil {
		return nil, err
	}
	return &account_infomation, nil
}

func (self accountDomain) Transaction(account_id int) ([]ExpenseTrackingEntity, error) {
	account, err := self.repo.Transactions(account_id)
	account_info_lists := []ExpenseTrackingEntity{}
	if err != nil {
		return nil, err
	}
	for _, item := range account {
		account_item := ExpenseTrackingEntity{
			ID:             item.ID,
			DateTime:       item.DateTime,
			Title:          item.Title,
			AccountID:      item.AccountID,
			Classification: item.Classification,
			Income:         item.Income,
			Expense:        item.Expense,
			OverAllBalance: item.OverAllBalance,
		}
		account_info_lists = append(account_info_lists, account_item)
	}
	return account_info_lists, nil
}
func (self accountDomain) Deposit(expense_tracking_data ExpenseTrackingEntity) (*ExpenseTrackingEntity, error) {
	new_deposit_record := repository.ExpenseTracking{
		ID:             expense_tracking_data.ID,
		DateTime:       expense_tracking_data.DateTime,
		Title:          expense_tracking_data.Title,
		AccountID:      expense_tracking_data.AccountID,
		Classification: expense_tracking_data.Classification,
		Income:         expense_tracking_data.Income,
	}
	record, err := self.repo.MakeDeposit(new_deposit_record)
	if err != nil {
		return nil, err
	}
	expense_tracking_data.OverAllBalance = record.OverAllBalance

	return &expense_tracking_data, nil
}
func (self accountDomain) WithDraw(expense_tracking_data ExpenseTrackingEntity) (*ExpenseTrackingEntity, error) {
	new_deposit_record := repository.ExpenseTracking{
		ID:             expense_tracking_data.ID,
		DateTime:       expense_tracking_data.DateTime,
		Title:          expense_tracking_data.Title,
		AccountID:      expense_tracking_data.AccountID,
		Classification: expense_tracking_data.Classification,
		Expense:        expense_tracking_data.Expense,
	}
	record, err := self.repo.MakeWithDrawal(new_deposit_record)
	if err != nil {
		return nil, err
	}
	expense_tracking_data.OverAllBalance = record.OverAllBalance

	return &expense_tracking_data, nil
}
