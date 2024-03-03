package my_bank_errors

import "errors"

var ErrFundsInsufficient = errors.New("the amount to transfer should be greater than the account's balance")
