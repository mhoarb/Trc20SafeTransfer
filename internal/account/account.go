package account

import "fmt"

// PrintAccounts prints the details of the accounts
func PrintAccounts(accounts []Account) {
	for _, acc := range accounts {
		fmt.Printf("%s: %.2f (Address: %s)\n", acc.ID, acc.Balance, acc.Address)
	}
}
