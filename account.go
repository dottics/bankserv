package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// GetUserAccounts gets all the bank accounts for a specific user based on
// the user's UUID passed to the function and returns a slice of Account.
// If an error occurs such as the user not found then an empty slice is returned
// and an error.
func (s *Service) GetUserAccounts(UUID uuid.UUID) (Accounts, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account/user/-"
	// add query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
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
	_, e = s.serv.Decode(r, &res)
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

// GetOrganisationAccounts gets all the bank accounts for a specific
// organisation based on the organisation's UUID and returns a slice of
// Account. If error occurs an error is returned.
func (s *Service) GetOrganisationAccounts(UUID uuid.UUID) (Accounts, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account/organisation/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return Accounts{}, e
	}

	type Data struct {
		Accounts `json:"accounts"`
	}
	res := struct {
		Data   `json:"data"`
		Errors map[string][]string
	}{}

	// decode the response
	_, e = s.serv.Decode(r, &res)
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
	// return the bank accounts on successful
	return res.Data.Accounts, nil
}

// CreateAccount creates a new bank account for either the user or
// organisation based on which UUID is provided. After creating the bank account
// it returns the bank account, or if an error occurs an error is returned.
func (s *Service) CreateAccount(b Account) (Account, dutil.Error) {
	// set path
	s.serv.URL.Path = "/bank-account"
	// marshal data to payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return Account{}, e
	}

	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
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
	_, e = s.serv.Decode(r, &res)
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
	s.serv.URL.Path = "/bank-account/-"
	// marshal payload reader
	p, e := dutil.MarshalReader(b)
	if e != nil {
		return Account{}, e
	}
	// do request
	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)
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
	_, e = s.serv.Decode(r, &res)
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
