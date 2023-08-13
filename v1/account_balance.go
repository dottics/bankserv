package v1

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/msp"
	"net/http"
	"time"
)

// GetAccountBalance gets the account balance for a specific bank account based
// on the bank account's UUID passed to the function and returns a pointer to
// the AccountBalance value. The balance returned is the first balance before
// the date passed to the function.
func (s *Service) GetAccountBalance(UUID uuid.UUID, date time.Time) (AccountBalance, dutil.Error) {
	// set path
	s.URL.Path = "/account/-/balance"

	// set query string
	qs := s.URL.Query()
	qs.Add("uuid", UUID.String())
	qs.Add("date", date.Format("2006-01-02"))

	// do request
	r, e := s.DoRequest(http.MethodGet, s.URL, qs, nil, nil)
	if e != nil {
		return AccountBalance{}, e
	}

	// create response structure
	type Data struct {
		AccountBalance `json:"account_balance"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}

	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return AccountBalance{}, e
	}

	// check the response status
	if r.StatusCode != http.StatusOK {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return AccountBalance{}, e
	}

	return res.Data.AccountBalance, nil
}

// CreateAccountBalance creates a new account balance for a specific bank
// account based on the bank account's UUID passed to the function and returns
// a pointer to the AccountBalance value.
func (s *Service) CreateAccountBalance(ab *AccountBalance) (AccountBalance, dutil.Error) {
	// set path
	s.URL.Path = "/account/-/balance"

	// marshal the account balance
	p, e := dutil.MarshalReader(ab)
	if e != nil {
		return AccountBalance{}, e
	}

	// do request
	r, e := s.DoRequest(http.MethodPost, s.URL, nil, nil, p)
	if e != nil {
		return AccountBalance{}, e
	}

	// create response structure
	type Data struct {
		AccountBalance `json:"account_balance"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}

	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return AccountBalance{}, e
	}

	// check the response status
	if r.StatusCode != http.StatusCreated {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return AccountBalance{}, e
	}
	return res.Data.AccountBalance, nil
}
