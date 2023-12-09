package v1

import (
	"github.com/google/uuid"
	"time"
)

type Bank struct {
	UUID       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	BranchCode string    `json:"branch_code"`
	SwiftCode  string    `json:"swift_code"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
type Banks []Bank

type Tag struct {
	UUID       uuid.UUID `json:"uuid"`
	EntityUUID uuid.UUID `json:"entity_uuid"`
	Tag        string    `json:"tag"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
type Tags []Tag

type Item struct {
	UUID               uuid.UUID `json:"uuid"`
	TransactionUUID    uuid.UUID `json:"transaction_uuid"`
	Description        string    `json:"description"`
	SKU                string    `json:"sku"`
	Unit               string    `json:"unit"`
	Quantity           float32   `json:"quantity"`
	Amount             float32   `json:"amount"`
	Discount           float32   `json:"discount"`
	Category           string    `json:"category"`
	PredictionCategory string    `json:"prediction_category"`
	Tags               []Tag     `json:"tags"`
	Active             bool      `json:"active"`
	CreateDate         time.Time `json:"create_date"`
	UpdateDate         time.Time `json:"update_date"`
}
type Items []Item

// Transaction describes a transaction, what is most important that the
// ExternalID represents an ID from an external source, if it is the zero string
// it means that the transaction was a manual entry by the user.
type Transaction struct {
	UUID         uuid.UUID `json:"uuid"`
	ExternalID   string    `json:"external_id"`
	AccountUUID  uuid.UUID `json:"account_uuid"`
	Date         time.Time `json:"date"`
	BusinessName string    `json:"business_name"`
	Description  string    `json:"description"`
	Debit        bool      `json:"debit"`
	Credit       bool      `json:"credit"`
	Amount       float32   `json:"amount"`
	Items        []Item    `json:"items"`
	Active       bool      `json:"active"`
	CreateDate   time.Time `json:"create_date"`
	UpdateDate   time.Time `json:"update_date"`
}
type Transactions []Transaction

// Account is a description of an account, this should represent any account
// be it a bank account, crypto wallet or anything similar.
// The Balance field is the current balance of the account.
type Account struct {
	UUID              uuid.UUID      `json:"uuid,"`
	BankUUID          uuid.UUID      `json:"bank_uuid"`
	Bank              Bank           `json:"bank"`
	EntityUUID        uuid.UUID      `json:"entity_uuid"`
	Name              string         `json:"name"`
	Alias             string         `json:"alias"`
	Number            string         `json:"number"`
	IntegrationStatus string         `json:"integration_status"`
	Balance           AccountBalance `json:"balance"`
	Active            bool           `json:"active"`
	CreateDate        time.Time      `json:"create_date"`
	UpdateDate        time.Time      `json:"update_date"`
}
type Accounts []Account

// AccountBalance describes the account's balance at a reference point or date
// in time.
type AccountBalance struct {
	AccountUUID uuid.UUID `json:"account_uuid"`
	Date        time.Time `json:"date"`
	Balance     float32   `json:"balance"`
	Active      bool      `json:"active"`
}

// timeMustParse is a function the parses a time string formatted based on the
// RFC3339 standard as 2006-01-02T15:04:05Z07:00 to a time.Time and returns
// the time.
func timeMustParse(value string) time.Time {
	// time.RFC3339 == "2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(time.RFC3339, value)
	return t
}

// UpdateAccount is only used for updating an account, it is not used for
// creating an account or getting an account.
//
// It is only used on the PUT request for updating an account.
type UpdateAccount struct {
	UUID              uuid.UUID `json:"uuid"`
	BankUUID          uuid.UUID `json:"bank_uuid,omitempty"`
	EntityUUID        uuid.UUID `json:"entity_uuid,omitempty"`
	Name              string    `json:"name,omitempty"`
	Alias             string    `json:"alias,omitempty"`
	Number            string    `json:"number,omitempty"`
	IntegrationStatus string    `json:"integration_status,omitempty"`
	Active            bool      `json:"active"`
}

// CategoryMonthTotal is a struct that describes the total amount spent on a
// category for a specific month and year.
type CategoryMonthTotal struct {
	Category string  `json:"category"`
	Debit    float32 `json:"debit"`
	Credit   float32 `json:"credit"`
	Year     int     `json:"year"`
	Month    int     `json:"month"`
}

type MapCategory struct {
	EntityUUID uuid.UUID `json:"entity_uuid"`
	Replace    string    `json:"replace"`
	With       string    `json:"with"`
}
