package dto

type RetriveBankAccount struct {
	ID            string      `json:"id"`
	Type          string      `json:"type"`
	Subtype       string      `json:"subtype"`
	Name          string      `json:"name"`
	Balance       float64     `json:"balance"`
	CurrencyCode  string      `json:"currencyCode"`
	ItemID        string      `json:"itemId"`
	Number        string      `json:"number"`
	MarketingName interface{} `json:"marketingName"`
	TaxNumber     string      `json:"taxNumber"`
	Owner         string      `json:"owner"`
	BankData      BankData    `json:"bankData"`
	CreditData    interface{} `json:"creditData"`
}
