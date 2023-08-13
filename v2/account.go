package v2

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/msp"
	"net/url"
)

// GetEntityAccounts gets all the accounts for a specific entity. An entity
// represents any user, organisation or group. Based on the UUID passed the
// function returns the Accounts. If an error occurs such as the UUID is invalid
// then the error is returned.
func (s *Service) GetEntityAccounts(UUID uuid.UUID) (Accounts, dutil.Error) {
	// set path
	s.URL.Path = "/account/entity/-"
	// add query string
	qs := url.Values{"uuid": {UUID.String()}}
	// do request
	r, e := s.DoRequest("GET", s.URL, qs, nil, nil)
	if e != nil {
		return Accounts{}, e
	}

	// response structure
	type Data struct {
		Accounts `json:"accounts"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string `json:"errors"`
	}{}

	// decode the response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Accounts{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Accounts{}, e
	}
	// return bank accounts on successful
	return res.Data.Accounts, nil
}

// CreateAccount creates a new bank account for either the user or
// organisation based on which UUID is provided. After creating the bank account
// it returns the bank account, or if an error occurs an error is returned.
func (s *Service) CreateAccount(b Account) (Account, dutil.Error) {
	// set path
	s.URL.Path = "/account"
	// marshal data to payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return Account{}, e
	}

	// do request
	r, e := s.DoRequest("POST", s.URL, nil, nil, p)
	if e != nil {
		return Account{}, e
	}

	type Data struct {
		Account `json:"account"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string
	}{}
	// decode the response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Account{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Account{}, e
	}
	// return bank account on successful
	return res.Data.Account, nil
}

// UpdateAccount updates a specific bank account's data.
func (s *Service) UpdateAccount(b Account) (Account, dutil.Error) {
	// set path
	s.URL.Path = "/account/-"
	// marshal payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return Account{}, e
	}
	// do request
	r, e := s.DoRequest("PUT", s.URL, nil, nil, p)
	if e != nil {
		return Account{}, e
	}

	type Data struct {
		Account `json:"account"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string `json:"errors"`
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Account{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Account{}, e
	}
	// return bank account on successful
	return res.Data.Account, nil
}

// DeleteAccount deletes a specific bank account's data.
func (s *Service) DeleteAccount(UUID uuid.UUID) dutil.Error {
	// set path
	s.URL.Path = "/account/-"
	qs := url.Values{"uuid": {UUID.String()}}

	// do request
	r, e := s.DoRequest("DELETE", s.URL, qs, nil, nil)
	if e != nil {
		return e
	}

	res := struct {
		Errors map[string][]string `json:"errors"`
	}{}

	// decode the response
	_, e = msp.Decode(r, &res)
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
