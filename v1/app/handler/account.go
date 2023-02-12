package handler

import (
	"github.com/PainestreaNakamoto/Expense-Tracking/domain"
	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	account_service domain.AccountInterface
}

func InitializeAccountHandler(account_service domain.AccountInterface) accountHandler {
	return accountHandler{account_service: account_service}
}

func (self accountHandler) AccountInfomation(c *fiber.Ctx) error {
	account_id, err := c.ParamsInt("account_id")
	if err != nil {
		return err
	}
	account_data, err := self.account_service.AccountInfo(account_id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(account_data)
}
