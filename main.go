package main

import (
	"fmt"

	"CryptoAccountSplitter/internal/account"
)

func main() {
	// Initialize the origin account
	origin, _ := account.NewAccount("OriginAccount", 2000)
	stage1Accounts := account.SplitAccount(origin, 2)
	stage2Accounts := account.SplitAccounts(stage1Accounts)
	stage3Accounts := account.SplitAccounts(stage2Accounts)

	stage4Accounts := account.ConsolidateAccounts(stage3Accounts, 4)
	stage5Accounts := account.ConsolidateAccounts(stage4Accounts, 2)

	finalAccount := account.ConsolidateAccounts(stage5Accounts, 1)
	fmt.Println(finalAccount[0].Balance, finalAccount[0].Address, finalAccount[0].PublicKey)
}
