package account

import "fmt"

// SplitAccount splits an account into a specified number of new accounts
func SplitAccount(origin Account, count int) []Account {
	accounts := make([]Account, count)
	amountPerAccount := origin.Balance / float64(count)

	for i := 0; i < count; i++ {
		id := fmt.Sprintf("%s_%d", origin.ID, i+1)
		accounts[i], _ = NewAccount(id, amountPerAccount)
	}

	return accounts
}

// SplitAccounts splits each account in a list into two
func SplitAccounts(accounts []Account) []Account {
	var newAccounts []Account
	for _, acc := range accounts {
		newAccounts = append(newAccounts, SplitAccount(acc, 2)...)
	}
	return newAccounts
}

// ConsolidateAccounts consolidates a list of accounts into a specified number of new accounts
func ConsolidateAccounts(accounts []Account, count int) []Account {
	newAccounts := make([]Account, count)
	amountPerAccount := ConsolidateAmounts(accounts) / float64(count)

	for i := 0; i < count; i++ {
		id := fmt.Sprintf("Account_C%d", i+1)
		newAccounts[i], _ = NewAccount(id, amountPerAccount)
	}

	return newAccounts
}
