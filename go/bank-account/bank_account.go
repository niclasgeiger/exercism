package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	closed  bool
	mutex   sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
	}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	payout = a.balance
	a.balance = 0
	return payout, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if amount+a.balance < 0 || a.closed {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, true
}
