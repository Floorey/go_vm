package core

import "fmt"

type Account struct {
	Address string
	Balance int
}
type AccountManager struct {
	Accounts map[string]*Account
}

func NewAccountManager() *AccountManager {
	return &AccountManager{Accounts: make(map[string]*Account)}
}
func (am *AccountManager) CreateAccount(address string, initialBalance int) *Account {
	account := &Account{Address: address, Balance: initialBalance}
	am.Accounts[address] = account
	return account
}
func (am *AccountManager) GetBalance(address string) (int, bool) {
	account, exists := am.Accounts[address]
	if !exists {
		return 0, false
	}
	return account.Balance, true
}
func (am *AccountManager) Transfer(from, to string, amount int) error {
	fromAccount, fromExists := am.Accounts[from]
	toAccount, toExists := am.Accounts[to]

	if !fromExists {
		return fmt.Errorf("account %s does not exist", from)
	}
	if !toExists {
		return fmt.Errorf("account %s does not exist", to)
	}
	if fromAccount.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	fromAccount.Balance -= amount
	toAccount.Balance += amount
	return nil
}
