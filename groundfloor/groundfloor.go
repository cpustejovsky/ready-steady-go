package groundfloor

import (
	"sync/atomic"
)

//tech exploration: 1 file package that implement a single bank account
//global shared variable; functions: Deposit adds to balance; Withdraw to subtract; CheckBalance to return balance

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(deposit int) int {
	b.Balance += deposit
	return b.Balance
}

func (b *BankAccount) Withdraw(withdrawal int) int {
	if withdrawal <= b.Balance {
		b.Balance -= withdrawal
	}
	return b.Balance
}

type BankAccountV2 struct {
	Balance int64
}

func (b *BankAccountV2) Deposit(deposit int64) {
	atomic.AddInt64(&b.Balance, deposit)
}

func (b *BankAccountV2) Withdraw(withdrawal int64) {
	if b.Balance >= int64(withdrawal) {
		atomic.AddInt64(&b.Balance, (withdrawal * -1))
	}

}
