package mock

import "api-fiber-mongo/entity"

type MockAccountRepository struct{}

func (m MockAccountRepository) GetAccountDb() entity.Account {
	return entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}

func (m MockAccountRepository) MakeDepositDb(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}

func (m MockAccountRepository) CreateAccount(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "2",
		Name:    "Don",
		Balance: 5000,
	}
}

func (m MockAccountRepository) MakeWithdrawDb(conta *entity.Account) *entity.Account {
	return &entity.Account{
		ID:      "1",
		Name:    "jack",
		Balance: 1000,
	}
}
