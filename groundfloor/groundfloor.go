package groundfloor

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
