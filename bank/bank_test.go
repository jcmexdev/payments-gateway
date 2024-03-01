package bank

import (
	"os"
	"testing"
)

var testAccount1, testAccount2 Account

func TestMain(m *testing.M) {
	testAccount1 = Account{
		Customer: Customer{
			Name: "John Doe Client",
		},
		Number:     1001,
		Balance:    0,
		CardNumber: "1234-5678-1234-5678",
	}

	testAccount2 = Account{
		Customer: Customer{
			Name: "John Doe Merchant",
		},
		Number:     1002,
		Balance:    0,
		CardNumber: "1111-2222-3333-4444",
	}

	exitCode := m.Run()

	testAccount1.Balance = 0
	testAccount2.Balance = 0

	os.Exit(exitCode)
}

func TestDeposit(t *testing.T) {
	_ = testAccount1.Deposit(10)
	if testAccount1.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	if err := testAccount1.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	testAccount1.Balance = 0
	_ = testAccount1.Deposit(10)
	_ = testAccount1.Withdraw(5)

	if testAccount1.Balance != 5 {
		t.Error("balance is not being updated after withdraw, expected 5 but got", testAccount1.Balance)
	}
}

func TestWithdraw_LowerThanZero(t *testing.T) {
	_ = testAccount1.Deposit(10)
	err := testAccount1.Withdraw(0)

	if err == nil || err.Error() != "the amount to withdraw should be greater than zero" {
		t.Error("expected error message for withdraw lower than zero")
	}
}

func TestWithdraw_LowerThanBalance(t *testing.T) {
	testAccount1.Balance = 0
	_ = testAccount1.Deposit(10)
	err := testAccount1.Withdraw(11)

	if err == nil || err.Error() != "the amount to withdraw should be greater than the account's balance" {
		t.Error("expected error message for withdraw lower than balance")
	}

}

func TestAccount_TransferSuccess(t *testing.T) {
	testAccount1.Balance = 0
	_ = testAccount1.Deposit(1000)

	err := testAccount1.Transfer(100, &testAccount2)
	if err != nil {
		t.Error("expected no error for transfer")
	}

	if testAccount1.Balance != 900 {
		t.Error("origin balance is not being updated after transfer")
	}

	if testAccount2.Balance != 100 {
		t.Error("destination balance is not being updated after transfer")
	}
}

func TestAccount_TransferFailsWithZeroAmount(t *testing.T) {
	testAccount1.Balance = 0
	testAccount2.Balance = 0
	_ = testAccount1.Deposit(1000)

	err := testAccount1.Transfer(0, &testAccount2)
	if err == nil {
		t.Error("expected error for transfer with zero amount")
	}

	if err.Error() != "the amount to transfer should be greater than zero" {
		t.Error("error message for transfer with zero amount is not being returned")
	}
}

func TestAccount_TransferFailsWithAmountGreaterThanBalance(t *testing.T) {
	testAccount1.Balance = 0
	testAccount2.Balance = 0
	_ = testAccount1.Deposit(500)

	err := testAccount1.Transfer(501, &testAccount2)
	if err == nil {
		t.Error("the amount to transfer should be greater than the account's balance")
	}

	if err.Error() != "the amount to transfer should be greater than the account's balance" {
		t.Error("error message for transfer with amount greater than balance is not being returned")
	}
}
