package models

import "time"

type Txn struct {
	TxnTime  time.Time `json:"txntime"`
	Amount   float64   `json:"amount"`
	Items    string    `json:"items"`
	Merchant string    `json:"merchant"`
	UserId   string    `json:"userid"`
	Envelope string    `json:"envelope"`
}
