package strategy

import (
	"CryptoAccountSplitter/internal/account"
	"fmt"
)

// SimpleSplitStrategy splits each account into two
type SimpleSplitStrategy struct{}

// Split splits each account into two
func (s *SimpleSplitStrategy) Split(accounts []account.Account) []account.Account {
	var newAccounts []account.Account
	for _, acc := range accounts {
		amountPerAccount := acc.Balance / 3
		newAccounts = append(newAccounts, account.Account{
			ID:         fmt.Sprintf("%s_1", acc.ID),
			Balance:    amountPerAccount,
			PrivateKey: acc.PrivateKey,
			PublicKey:  acc.PublicKey,
			Address:    acc.Address,
		}, account.Account{
			ID:         fmt.Sprintf("%s_2", acc.ID),
			Balance:    amountPerAccount,
			PrivateKey: acc.PrivateKey,
			PublicKey:  acc.PublicKey,
			Address:    acc.Address,
		})
	}
	return newAccounts
}
