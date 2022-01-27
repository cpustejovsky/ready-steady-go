package groundfloor_test

import (
	"sync"
	"testing"
	"time"

	"github.com/cpustejovsky/ready-steady-go/groundfloor"
)

func TestDeposit(t *testing.T) {
	ba := groundfloor.BankAccount{
		Balance: 0,
	}
	got := ba.Deposit(10)
	want := 10

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

}

func TestWithdraw(t *testing.T) {
	ba := groundfloor.BankAccount{
		Balance: 10,
	}
	got := ba.Withdraw(10)
	want := 0

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

}

var deposits = []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
var withdrawals = []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}

func TestConcurrentDepositWithDraw(t *testing.T) {
	ba := groundfloor.BankAccount{
		Balance: 0,
	}

	var wg sync.WaitGroup

	for _, d := range deposits {
		wg.Add(1)
		go func(d int) {
			defer wg.Done()
			ba.Deposit(d)
		}(d)
	}
	for _, w := range withdrawals {
		wg.Add(1)
		go func(w int) {
			defer wg.Done()
			ba.Withdraw(w)
		}(w)
	}

	wg.Wait()

	got := ba.Balance
	want := 100
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestConcurrentDepositWithDrawV2(t *testing.T) {
	ba := groundfloor.BankAccountV2{
		Balance: 0,
	}

	for _, d := range deposits {
		go func(d int64) {
			ba.Deposit(d)
		}(int64(d))
	}
	for _, w := range withdrawals {
		go func(w int64) {
			ba.Withdraw(w)
		}(int64(w))
	}
	time.Sleep(1 * time.Second)

	got := ba.Balance
	var want int64 = 100
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
