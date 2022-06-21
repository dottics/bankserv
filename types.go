package bankserv

import (
	"github.com/google/uuid"
	"time"
)

type Bank struct {
	UUID       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	BranchCode string    `json:"branch_code"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}
type Banks []Bank

type Tag struct {
	UUID             uuid.UUID `json:"uuid"`
	UserUUID         uuid.UUID `json:"user_uuid"`
	OrganisationUUID uuid.UUID `json:"organisation_uuidUUID"`
	Tag              string    `json:"tag"`
	Active           bool      `json:"active"`
	CreateDate       time.Time `json:"create_date"`
	UpdateDate       time.Time `json:"update_date"`
}
type Tags []Tag

type Item struct {
	UUID            uuid.UUID `json:"uuid"`
	TransactionUUID uuid.UUID `json:"transaction_uuid"`
	Description     string    `json:"description"`
	SKU             float32   `json:"sku"`
	Amount          float32   `json:"amount"`
	Discount        float32   `json:"discount"`
	Tags            []Tag     `json:"tags"`
	Active          bool      `json:"active"`
	CreateDate      time.Time `json:"create_date"`
	UpdateDate      time.Time `json:"update_date"`
}
type Items []Item

type Transaction struct {
	UUID        uuid.UUID `json:"uuid"`
	AccountUUID uuid.UUID `json:"bank_account_uuid"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Items       []Item    `json:"items"`
	Active      bool      `json:"active"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
}
type Transactions []Transaction

type BankAccount struct {
	UUID             uuid.UUID `json:"uuid"`
	UserUUID         uuid.UUID `json:"user_uuid"`
	OrganisationUUID uuid.UUID `json:"organisation_uuid"`
	AccountNumber    string    `json:"account_number"`
	Active           bool      `json:"active"`
	CreateDate       time.Time `json:"create_date"`
	UpdateDate       time.Time `json:"update_date"`
}
type BankAccounts []BankAccount

// timeMustParse is a function the parses a time string formatted based on the
// RFC3339 standard as 2006-01-02T15:04:05Z07:00 to a time.Time and returns
// the time.
func timeMustParse(value string) time.Time {
	// time.RFC3339 == "2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(time.RFC3339, value)
	return t
}
