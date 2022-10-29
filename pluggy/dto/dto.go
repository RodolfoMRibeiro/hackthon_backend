package dto

import "time"

type CreateItemUserDto struct {
	ID              string      `json:"id"`
	Connector       Connector   `json:"connector"`
	CreatedAt       time.Time   `json:"createdAt"`
	UpdatedAt       time.Time   `json:"updatedAt"`
	Status          string      `json:"status"`
	ExecutionStatus string      `json:"executionStatus"`
	LastUpdatedAt   interface{} `json:"lastUpdatedAt"`
	WebhookURL      interface{} `json:"webhookUrl"`
	Error           Error       `json:"error"`
	ClientUserID    interface{} `json:"clientUserId"`
	StatusDetail    interface{} `json:"statusDetail"`
	Parameter       interface{} `json:"parameter"`
}
type Credentials struct {
	Label             string `json:"label"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Placeholder       string `json:"placeholder"`
	Validation        string `json:"validation"`
	ValidationMessage string `json:"validationMessage"`
	Optional          bool   `json:"optional"`
}
type Health struct {
	Status string      `json:"status"`
	Stage  interface{} `json:"stage"`
}
type Connector struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	PrimaryColor   string        `json:"primaryColor"`
	InstitutionURL string        `json:"institutionUrl"`
	Country        string        `json:"country"`
	Type           string        `json:"type"`
	Credentials    []Credentials `json:"credentials"`
	ImageURL       string        `json:"imageUrl"`
	HasMFA         bool          `json:"hasMFA"`
	Health         Health        `json:"health"`
	Products       []string      `json:"products"`
	CreatedAt      time.Time     `json:"createdAt"`
}
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
