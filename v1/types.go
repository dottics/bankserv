package v1

import (
	"github.com/google/uuid"
	"time"
)

type Bank struct {
	UUID       uuid.UUID `json:"uuid,omitempty"`
	Name       string    `json:"name,omitempty"`
	BranchCode string    `json:"branch_code,omitempty"`
	SwiftCode  string    `json:"swift_code,omitempty"`
	Active     bool      `json:"active,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
	UpdateDate time.Time `json:"update_date,omitempty"`
}
type Banks []Bank

type Tag struct {
	UUID       uuid.UUID `json:"uuid,omitempty"`
	EntityUUID uuid.UUID `json:"entity_uuid,omitempty"`
	Tag        string    `json:"tag,omitempty"`
	Active     bool      `json:"active,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
	UpdateDate time.Time `json:"update_date,omitempty"`
}
type Tags []Tag

type Item struct {
	UUID            uuid.UUID `json:"uuid,omitempty"`
	TransactionUUID uuid.UUID `json:"transaction_uuid,omitempty"`
	Description     string    `json:"description,omitempty"`
	SKU             string    `json:"sku,omitempty"`
	Unit            string    `json:"unit,omitempty"`
	Quantity        float32   `json:"quantity,omitempty"`
	Amount          float32   `json:"amount,omitempty"`
	Discount        float32   `json:"discount,omitempty"`
	Tags            []Tag     `json:"tags,omitempty"`
	Active          bool      `json:"active,omitempty"`
	CreateDate      time.Time `json:"create_date,omitempty"`
	UpdateDate      time.Time `json:"update_date,omitempty"`
}
type Items []Item

// Transaction describes a transaction, what is most important that the
// ExternalID represents an ID from an external source, if it is the zero string
// it means that the transaction was a manual entry by the user.
type Transaction struct {
	UUID         uuid.UUID `json:"uuid,omitempty"`
	ExternalID   string    `json:"external_id,omitempty"`
	AccountUUID  uuid.UUID `json:"account_uuid,omitempty"`
	Date         time.Time `json:"date,omitempty"`
	BusinessName string    `json:"business_name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Debit        bool      `json:"debit,omitempty"`
	Credit       bool      `json:"credit,omitempty"`
	Amount       float32   `json:"amount,omitempty"`
	Items        []Item    `json:"items,omitempty"`
	Active       bool      `json:"active,omitempty"`
	CreateDate   time.Time `json:"create_date,omitempty"`
	UpdateDate   time.Time `json:"update_date,omitempty"`
}
type Transactions []Transaction

// Account is a description of an account, this should represent any account
// be it a bank account, crypto wallet or anything similar.
type Account struct {
	UUID              uuid.UUID `json:"uuid,omitempty"`
	BankUUID          uuid.UUID `json:"bank_uuid,omitempty"`
	Bank              Bank      `json:"bank,omitempty"`
	EntityUUID        uuid.UUID `json:"entity_uuid,omitempty"`
	Name              string    `json:"name,omitempty"`
	Alias             string    `json:"alias,omitempty"`
	Number            string    `json:"number,omitempty"`
	IntegrationStatus string    `json:"integration_status,omitempty"`
	Active            bool      `json:"active,omitempty"`
	CreateDate        time.Time `json:"create_date,omitempty"`
	UpdateDate        time.Time `json:"update_date,omitempty"`
}
type Accounts []Account

// AccountBalance describes the account's balance at a reference point or date
// in time.
type AccountBalance struct {
	AccountUUID uuid.UUID `json:"account_uuid,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Balance     float32   `json:"balance,omitempty"`
	Active      bool      `json:"active,omitempty"`
}

// timeMustParse is a function the parses a time string formatted based on the
// RFC3339 standard as 2006-01-02T15:04:05Z07:00 to a time.Time and returns
// the time.
func timeMustParse(value string) time.Time {
	// time.RFC3339 == "2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(time.RFC3339, value)
	return t
}
