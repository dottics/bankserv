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

type Tag struct {
	UUID       uuid.UUID `json:"uuid"`
	Tag        string    `json:"tag"`
	Active     bool      `json:"active"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date"`
}

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

type Transaction struct {
	UUID            uuid.UUID `json:"uuid"`
	BankAccountUUID uuid.UUID `json:"bank_account_uuid"`
	Date            time.Time `json:"date"`
	Description     string    `json:"description"`
	Items           []Item    `json:"items"`
	Active          bool      `json:"active"`
	CreateDate      time.Time `json:"create_date"`
	UpdateDate      time.Time `json:"update_date"`
}

type BankAccount struct {
	UUID             uuid.UUID     `json:"uuid"`
	UserUUID         uuid.UUID     `json:"user_uuid"`
	OrganisationUUID uuid.UUID     `json:"organisation_uuid"`
	AccountNumber    string        `json:"account_number"`
	Transactions     []Transaction `json:"transactions"`
	Active           bool          `json:"active"`
	CreateDate       time.Time     `json:"create_date"`
	UpdateDate       time.Time     `json:"update_date"`
}
