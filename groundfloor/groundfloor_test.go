package groundfloor_test

import (
	"runtime"
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

func TestConcurrentDepositWithDraw(t *testing.T) {
	ba := groundfloor.BankAccount{
		Balance: 0,
	}
	g := runtime.GOMAXPROCS(0)
	d := make(chan int, g)
	w := make(chan int, g)
	deposits := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	withdrawals := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100} // go func() {

		go func() {
		defer close(d)
		for _, deposit := range deposits {
			d <- deposit
		}
	}()

	go func() {
		defer close(w)
		for _, withdrawal := range withdrawals {
			w <- withdrawal
		}
	}()

	for desposit := range d {
		
	}
	for _, d := range deposits {
		go ba.Deposit(d)
	}
	for _, w := range withdrawals {
		go ba.Withdraw(w)
	}
	time.Sleep(5 * time.Second)
	got := ba.Balance
	want := 100
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
