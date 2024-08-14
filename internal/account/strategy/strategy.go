package strategy

import "CryptoAccountSplitter/internal/account"

// SplitStrategy is the interface for splitting accounts
type SplitStrategy interface {
	Split(accounts []account.Account) []account.Account
}

// ConsolidateStrategy is the interface for consolidating accounts
type ConsolidateStrategy interface {
	Consolidate(accounts []account.Account) []account.Account
}
