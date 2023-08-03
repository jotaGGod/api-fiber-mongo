package service

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/mock"
	"api-fiber-mongo/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Account
	}{
		{
			name: "Test should return an account with all the field same as want test attribute",
			want: []entity.Account{
				{
					ID:      "1",
					Name:    "jack",
					Balance: 1000,
				},
			},
		},
	}
	repository.AccountRepository = mock.MockAccountRepository{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accounts := GetAccounts()
			assert.Equal(t, tt.want, accounts)
		})
	}
}

func TestMakeDeposits(t *testing.T) {
	tests := []struct {
		name string
		args int
		want *entity.Account
	}{
		{
			name: "should return an account with updated balance",
			args: 100,
			want: &entity.Account{
				ID:      "1",
				Name:    "jack",
				Balance: 1100,
			},
		},
	}
	repository.AccountRepository = mock.MockAccountRepository{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultAccount := MakeDeposits(tt.args)

			assert.Equal(t, tt.want, resultAccount)
		})
	}
}

func TestCreateAccounts(t *testing.T) {
	tests := []struct {
		name string
		args *entity.Account
		want *entity.Account
	}{
		{
			name: "should return a new account created",
			args: &entity.Account{
				ID:      "2",
				Name:    "Don",
				Balance: 5000,
			},
			want: &entity.Account{
				ID:      "2",
				Name:    "Don",
				Balance: 5000,
			},
		},
	}

	repository.AccountRepository = mock.MockAccountRepository{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultAccount := CreateAccounts(tt.args)

			assert.Equal(t, tt.want, resultAccount)
		})
	}
}

func TestMakeWithdraw(t *testing.T) {
	tests := []struct {
		name string
		args int
		want *entity.Account
	}{
		{
			name: "should return an account with updated balance",
			args: 250,
			want: &entity.Account{
				ID:      "1",
				Name:    "jack",
				Balance: 750,
			},
		},
	}
	repository.AccountRepository = mock.MockAccountRepository{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultAccount := MakeWithdraw(tt.args)

			assert.Equal(t, tt.want, resultAccount)
		})
	}
}
