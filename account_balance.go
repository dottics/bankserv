package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// GetAccountBalance gets the account balance for a specific bank account based
// on the bank account's UUID passed to the function and returns a pointer to
// the AccountBalance value. The balance returned is the first balance before
// the date passed to the function.
func (s *Service) GetAccountBalance(UUID uuid.UUID, date time.Time, headers *http.Header) (AccountBalance, dutil.Error) {
	// set path
	s.serv.URL.Path = "/account/-/balance"

	// set query string
	qs := s.serv.URL.Query()
	qs.Add("uuid", UUID.String())
	qs.Add("date", date.Format("2006-01-02"))
	s.serv.URL.RawQuery = qs.Encode()

	// do request
	r, e := s.serv.NewRequest(http.MethodGet, s.serv.URL.String(), *headers, nil)
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
	_, e = s.serv.Decode(r, &res)
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
func (s *Service) CreateAccountBalance(ab *AccountBalance, headers *http.Header) (AccountBalance, dutil.Error) {
	// set path
	s.serv.URL.Path = "/account/-/balance"

	// marshal the account balance
	p, e := dutil.MarshalReader(ab)
	if e != nil {
		return AccountBalance{}, e
	}

	// do request
	r, e := s.serv.NewRequest(http.MethodPost, s.serv.URL.String(), *headers, p)
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
	_, e = s.serv.Decode(r, &res)
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
