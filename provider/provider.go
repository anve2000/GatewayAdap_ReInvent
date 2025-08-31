package provider

import (
	"gatewayAdapterReInvent/internal/model"
)

type PaymentProvider interface{
	ProcessPayment(amount int64, currency string)(*model.PaymentResult, *model.PaymentError)
}