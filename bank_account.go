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
	return res.Data.BankAccounts, nil
}

// GetOrganisationBankAccounts gets all the bank accounts for a specific
// organisation based on the organisation's UUID and returns a slice of
// BankAccount. If error occurs an error is returned.
func (s *Service) GetOrganisationBankAccounts(UUID uuid.UUID) (BankAccounts, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account/organisation/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return BankAccounts{}, e
	}

	type Data struct {
		BankAccounts `json:"bank_accounts"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string
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
	// return the bank accounts on successful
	return res.Data.BankAccounts, nil
}

// CreateBankAccount creates a new bank account for either the user or
// organisation based on which UUID is provided. After creating the bank account
// it returns the bank account, or if an error occurs an error is returned.
func (s *Service) CreateBankAccount(b BankAccount) (BankAccount, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account"
	// marshal data to payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return BankAccount{}, e
	}

	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
	if e != nil {
		return BankAccount{}, e
	}

	type Data struct {
		BankAccount `json:"bank_account"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return BankAccount{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return BankAccount{}, e
	}
	// return bank account on successful
	return res.Data.BankAccount, nil
}

// UpdateBankAccount updates a specific bank account's data.
func (s *Service) UpdateBankAccount(b BankAccount) (BankAccount, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account/-"
	// marshal payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return BankAccount{}, e
	}
	// do request
	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)
	if e != nil {
		return BankAccount{}, e
	}

	type Data struct {
		BankAccount `json:"bank_account"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return BankAccount{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return BankAccount{}, e
	}
	// return bank account on successful
	return res.Data.BankAccount, nil
}

// DeleteBankAccount deletes a specific bank account's data.
func (s *Service) DeleteBankAccount(UUID uuid.UUID) dutil.Error {
	// set path
	s.serv.URL.Path = "/bank-account/-"
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()

	// do request
	r, e := s.serv.NewRequest("DELETE", s.serv.URL.String(), nil, nil)
	if e != nil {
		return e
	}

	res := struct {
		Errors map[string][]string `json:"errors"`
	}{}

	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return e
	}
	return nil
}
