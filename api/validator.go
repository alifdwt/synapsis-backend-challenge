package api

import (
	"github.com/alifdwt/synapsis-backend-challenge/util"
	"github.com/go-playground/validator/v10"
)

var validPaymentMethod validator.Func = func(fl validator.FieldLevel) bool {
	if paymentMethod, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedPaymentMethod(paymentMethod)
	}
	return false
}
