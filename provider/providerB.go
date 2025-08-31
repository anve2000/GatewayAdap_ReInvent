package provider

import (
	"gatewayAdapterReInvent/internal/model"
	"strconv"
	"time"
)

type ProviderB struct{}

func NewProviderB() *ProviderB {
	return &ProviderB{}
}

func (p *ProviderB) ProcessPayment(amount int64, currency string) (*model.PaymentResult, *model.PaymentError) {
	raw := struct {
		PaymentID string
		State     string
		Value     struct {
			Amount       string
			CurrencyCode string
		}

		ProcessedAt int64
	}{
		PaymentID: "PAY-789-XYZ",
		State:     "SUCCESS",
		Value: struct {
			Amount       string
			CurrencyCode string
		}{
			Amount:       "100",
			CurrencyCode: "USD",
		},
		ProcessedAt: 17392733,
	}

	floatAmount, _ := strconv.ParseFloat(raw.Value.Amount, 64)
	amountCents := int64(floatAmount * 100)

	return &model.PaymentResult{
		TransactionID: raw.PaymentID,
		Status:        model.PaymentStatus(raw.State),
		Amount:        amountCents,
		Currency:      raw.Value.CurrencyCode,
		Timestamp:     time.Unix(raw.ProcessedAt, 0).UTC(),
	}, nil
}
