package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		got := wallet.Withdraw(Bitcoin(100))

		assertError(t, got, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Expected %s got %s ", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got error, but didn't wanted one.")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error, but didn't get one.")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
