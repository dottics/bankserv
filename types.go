package bankserv

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
	UUID            uuid.UUID `json:"uuid"`
	TransactionUUID uuid.UUID `json:"transaction_uuid"`
	Description     string    `json:"description"`
	SKU             string    `json:"sku"`
	Unit            string    `json:"unit"`
	Quantity        float32   `json:"quantity"`
	Amount          float32   `json:"amount"`
	Discount        float32   `json:"discount"`
	Tags            []Tag     `json:"tags"`
	Active          bool      `json:"active"`
	CreateDate      time.Time `json:"create_date"`
	UpdateDate      time.Time `json:"update_date"`
}
type Items []Item

type Transaction struct {
	UUID         uuid.UUID `json:"uuid"`
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

type Account struct {
	UUID       uuid.UUID `json:"uuid"`
	BankUUID   uuid.UUID `json:"bank_uuid"`
	Bank       Bank      `json:"bank"`
	EntityUUID uuid.UUID `json:"entity_uuid"`
	Name       string    `json:"name"`
	Alias      string    `json:"alias"`
	Number     string    `json:"number"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
type Accounts []Account

// timeMustParse is a function the parses a time string formatted based on the
// RFC3339 standard as 2006-01-02T15:04:05Z07:00 to a time.Time and returns
// the time.
func timeMustParse(value string) time.Time {
	// time.RFC3339 == "2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(time.RFC3339, value)
	return t
}
