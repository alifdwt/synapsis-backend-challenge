package util

const (
	COD           = "COD"
	BANK_TRANSFER = "BANK_TRANSFER"
	E_WALLET      = "E_WALLET"
)

func IsSupportedPaymentMethod(paymentMethod string) bool {
	switch paymentMethod {
	case COD, BANK_TRANSFER, E_WALLET:
		return true
	}
	return false
}
