package ptrs_test

import (
	"testing"

	"github.com/cpustejovsky/ready-steady-go/tdd/ptrs"
)

func assertBalance(t testing.TB, wallet ptrs.Wallet, want ptrs.Bitcoin) {
	t.Helper()
	got := wallet.Balance

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := ptrs.Wallet{}
		wallet.Deposit(ptrs.Bitcoin(10))
		assertBalance(t, wallet, ptrs.Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		tests := []struct {
			name       string
			wallet     ptrs.Wallet
			withdrawal ptrs.Bitcoin
			want       ptrs.Bitcoin
		}{
			{"standard", ptrs.Wallet{Balance: ptrs.Bitcoin(20)}, ptrs.Bitcoin(10), ptrs.Bitcoin(10)},
			{"insufficient funds", ptrs.Wallet{Balance: ptrs.Bitcoin(10)}, ptrs.Bitcoin(20), ptrs.Bitcoin(10)},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.wallet.Withdraw(tt.withdrawal)
				if (tt.wallet.Balance != tt.want) && err == nil {
					t.Errorf("got %s want %s", tt.wallet.Balance, tt.want)
				}
			})
		}
	})

	t.Run("Withdraw Error", func(t *testing.T) {
		startingBalance := ptrs.Bitcoin(10)
		wallet := ptrs.Wallet{Balance: startingBalance}
		withdrawal := ptrs.Bitcoin(20)
		err := wallet.Withdraw(withdrawal)
		assertError(t, err, ptrs.ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

}
