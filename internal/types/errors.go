package types

import (
	"fmt"
)

type AccountError struct {
	code  int
	error string
}

var _ error = &AccountError{}

var BalanceNotSufficientFunds = AccountError{1, "account balance not sufficient funds"}

func (a AccountError) Error() string {
	return fmt.Sprintf("error code: %d. error: %s", a.code, a.error)
}
