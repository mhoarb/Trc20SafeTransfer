package strategy

import (
	"CryptoAccountSplitter/internal/account"
	"fmt"
)

type CustomSplitStrategy struct {
	SplitCount int
}

// Split splits each account into the specified number of parts
func (s *CustomSplitStrategy) Split(accounts []account.Account) []account.Account {
	var newAccounts []account.Account
	for _, acc := range accounts {
		amountPerAccount := acc.Balance / float64(s.SplitCount)
		for i := 0; i < s.SplitCount; i++ {
			newAccounts = append(newAccounts, account.Account{
				ID:         fmt.Sprintf("%s_%d", acc.ID, i+1),
				Balance:    amountPerAccount,
				PrivateKey: acc.PrivateKey,
				PublicKey:  acc.PublicKey,
				Address:    acc.Address,
			})
		}
	}
	return newAccounts
}
