package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// GetUserBankAccounts gets all the bank accounts for a specific user based on
// the user's UUID passed to the function and returns a slice of BankAccount.
// If an error occurs such as the user not found then an empty slice is returned
// and an error.
func (s *Service) GetUserBankAccounts(UUID uuid.UUID) (BankAccounts, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account/user/-"
	// add query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return BankAccounts{}, e
	}

	// response structure
	type Data struct {
		BankAccounts `json:"bank_accounts"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string `json:"errors"`
	}{}

	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return BankAccounts{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return BankAccounts{}, e
	}
	// return bank accounts on successful
	return res.BankAccounts, nil
}

// GetOrganisationBankAccounts gets all the bank accounts for a specific
// organisation based on the organisation's UUID and returns a slice of
// BankAccount. If error occurs an error is returned.
func (s *Service) GetOrganisationBankAccounts(UUID uuid.UUID) (BankAccounts, dutil.Error) {
	return BankAccounts{}, nil
}

// CreateBankAccount creates a new bank account for either the user or
// organisation based on which UUID is provided. After creating the bank account
// it returns the bank account, or if an error occurs an error is returned.
func (s *Service) CreateBankAccount() (BankAccount, dutil.Error) {
	return BankAccount{}, nil
}

// UpdateBankAccount updates a specific bank account's data.
func (s *Service) UpdateBankAccount() (BankAccount, dutil.Error) {
	return BankAccount{}, nil
}

// DeleteBankAccount deletes a specific bank account's data.
func (s *Service) DeleteBankAccount(UUID uuid.UUID) dutil.Error {
	return nil
}
