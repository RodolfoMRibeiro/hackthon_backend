package dto

import "time"

type TransactionData struct {
	Total                int                    `json:"total"`
	TotalPages           int                    `json:"totalPages"`
	Page                 int                    `json:"page"`
	AccountOfTransaction []AccountOfTransaction `json:"results"`
}
type AccountOfTransaction struct {
	ID                 string      `json:"id"`
	Description        string      `json:"description"`
	DescriptionRaw     string      `json:"descriptionRaw"`
	CurrencyCode       string      `json:"currencyCode"`
	Amount             int         `json:"amount"`
	Date               time.Time   `json:"date"`
	Category           interface{} `json:"category"`
	Balance            float64     `json:"balance"`
	AccountID          string      `json:"accountId"`
	ProviderCode       interface{} `json:"providerCode"`
	Status             string      `json:"status"`
	PaymentData        interface{} `json:"paymentData"`
	Type               string      `json:"type"`
	CreditCardMetadata interface{} `json:"creditCardMetadata"`
	Merchant           interface{} `json:"merchant"`
}
