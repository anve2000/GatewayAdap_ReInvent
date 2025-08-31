package provider

import (
	"encoding/json"
	"gatewayAdapterReInvent/internal/model"
	"strconv"
	"time"
)

type ProviderA struct{}

func NewProviderA() *ProviderA {
	return &ProviderA{}
}

func (p *ProviderA) ProcessPayment(amount int64, currency string) (*model.PaymentResult, *model.PaymentError) {
	timestampStr := "2024-01-15T10:30:00Z"
	t, err := time.Parse(time.RFC3339, timestampStr)

	if err != nil {
		return nil, &model.PaymentError{
			Code:    "PARSE_ERROR",
			Message: "Failed To Parse timestamp",
		}
	}

	raw := struct {
		TransactionID string
		Status        string
		Amount        int64
		Currency      string
		Timestamp     time.Time
	}{
		TransactionID: "TXN123456",
		Status:        "Approved",
		Amount:        amount,
		Currency:      currency,
		Timestamp:     t,
	}

	return &model.PaymentResult{
		TransactionID: raw.TransactionID,
		Status:        model.PaymentStatus(raw.Status),
		Amount:        raw.Amount,
		Currency:      raw.Currency,
		Timestamp:     t,
	}, nil
}

func ParseProviderAResponse(data []byte) (*model.PaymentResult, *model.PaymentError) {
	var raw struct {
		PaymentId string
		State     string
		Value     struct {
			Amount       string
			CurrencyCode string
		}
		ProcessedAt int64
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, &model.PaymentError{
			Code:    "JSON PARSE ERROR",
			Message: err.Error(),
		}
	}

	floatAmount, err := strconv.ParseFloat(raw.Value.Amount, 64)
	if err != nil {
		return nil, &model.PaymentError{
			Code:    "AMOUNT PARSE ERROR",
			Message: err.Error(),
		}
	}

	amountCents:=int64(floatAmount*100)

	return &model.PaymentResult{
		TransactionID: raw.PaymentId,
		Status: model.PaymentStatus(raw.State),
		Amount: amountCents,
		Currency: raw.Value.CurrencyCode,
		Timestamp: time.Unix(raw.ProcessedAt, 0).UTC(),
	}, nil

}
