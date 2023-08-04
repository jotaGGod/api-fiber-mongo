package mock

import "api-fiber-mongo/entity"

type MAccountRepository struct{}

func (m *MAccountRepository) GetAccountDb(id string) *entity.Account {
	return &entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}

func (m *MAccountRepository) MakeDepositDb(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}

func (m *MAccountRepository) CreateAccount(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "2",
		Name:    "Don",
		Balance: 5000,
	}
}

func (m *MAccountRepository) MakeWithdrawDb(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}
