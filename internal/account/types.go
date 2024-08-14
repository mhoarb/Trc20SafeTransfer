package account

import "CryptoAccountSplitter/internal/keygen"

// Account represents a cryptocurrency account
type Account struct {
	ID         string
	PrivateKey string
	PublicKey  string
	Address    string
	Balance    float64
}

// NewAccount creates a new account with generated keys and address
func NewAccount(id string, balance float64) (Account, error) {
	privateKey, publicKeyBytes, err := keygen.GenerateKeyPair()
	if err != nil {
		return Account{}, err
	}

	privateKeyHex := keygen.PrivateKeyToHex(privateKey)
	publicKeyHex := keygen.PublicKeyToHex(publicKeyBytes)
	address := keygen.GenerateAddress(publicKeyBytes)

	return Account{
		ID:         id,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
		Address:    address,
		Balance:    balance,
	}, nil
}

// ConsolidateAmounts sums up the balances of all accounts in a slice
func ConsolidateAmounts(accounts []Account) float64 {
	total := 0.0
	for _, account := range accounts {
		total += account.Balance
	}
	return total
}
