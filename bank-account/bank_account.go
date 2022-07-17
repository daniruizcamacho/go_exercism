package account

import "sync"

type Account struct {
	balance  int64
	isClosed bool
	mu       sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance:  amount,
		isClosed: false,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if a.isClosed {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}

	if (a.balance + amount) < 0 {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.isClosed {
		return 0, false
	}

	balanceBeforeClosing := a.balance
	a.isClosed = true
	a.balance = 0

	return balanceBeforeClosing, true
}
