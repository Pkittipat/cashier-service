package responses

import "errors"

var (
	ErrInvalidPaymentLessThanPrice = errors.New("payment shouldn't less that price")
)
