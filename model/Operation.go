package model

import "time"

type Operation struct {
	Id                        string    `json:"id"`
	ExecutionDate             time.Time `json:"execution_date"`
	ValueDate                 time.Time `json:"value_date"`
	Amount                    float64   `json:"amount"`
	Currency                  string    `json:"currency"`
	AccountNumber             string    `json:"account_number"`
	TransactionType           string    `json:"transaction_type"`
	CounterpartyAccountNumber string    `json:"counterparty_account_number"`
	CounterpartyName          string    `json:"counterparty_name"`
	Communication             string    `json:"communication"`
	Details                   string    `json:"details"`
	Status                    string    `json:"status"`
	ReasonForRefusal          string    `json:"reason_for_refusal"`
}
