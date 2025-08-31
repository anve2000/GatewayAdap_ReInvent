package model

import "time"

type PaymentStatus string

const (
	StatusApproved PaymentStatus = "APPROVED"
	StatusFailed PaymentStatus = "FAILED"
)


type PaymentResult struct{
	TransactionID string
	Status PaymentStatus
	Amount int64
	Currency string
	Timestamp time.Time
}

type PaymentError struct{
	Code string
	Message string
}