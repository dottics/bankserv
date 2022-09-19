package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// GetBankAccountTransactions gets all the transactions for a specific bank
// account based on the bank account's UUID passed to the function and returns
// a slice of Transaction. If an error occurs the error will not be nil. If the
// bank account has no transactions an empty slice will be returned.
func (s *Service) GetBankAccountTransactions(UUID uuid.UUID) (Transactions, dutil.Error) {
	// set path
	s.serv.URL.Path = "/transaction/bank-account/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return Transactions{}, e
	}
	type Data struct {
		Transactions `json:"transactions"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Transactions{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Transactions{}, e
	}
	return res.Data.Transactions, nil
}

// GetUserTransactions gets all the transactions for a user based on the user's
// UUID. It does this by filtering by all the user's bank accounts. Then returns
// a slice of Transaction. If an error occurs the error wil not be nil. If the
// user's bank accounts have no transactions then an empty slice will be
// returned.
func (s *Service) GetUserTransactions(UUID uuid.UUID) (Transactions, dutil.Error) {
	// set path
	s.serv.URL.Path = "/transaction/user/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()

	r, e := s.serv.NewRequest("GET", s.serv.URL.String(), nil, nil)
	if e != nil {
		return Transactions{}, e
	}

	type Data struct {
		Transactions Transactions `json:"transactions"`
	}
	res := struct {
		Data   Data         `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}

	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Transactions{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Transactions{}, e
	}
	return res.Data.Transactions, nil
}

// CreateTransaction creates a new transaction for a bank account based on the
// transaction data that is passed to the function.
func (s *Service) CreateTransaction(t Transaction) (Transaction, dutil.Error) {
	// set path
	s.serv.URL.Path = "/transaction"
	// marshal payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Transaction{}, e
	}
	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
	if e != nil {
		return Transaction{}, e
	}

	type Data struct {
		Transaction `json:"transaction"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Transaction{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Transaction{}, e
	}
	// return transaction on successful exchange
	return res.Data.Transaction, nil
}

// UpdateTransaction updates a transaction for a bank account based on the
// transaction's UUID and transaction data that is passed to the function.
func (s *Service) UpdateTransaction(t Transaction) (Transaction, dutil.Error) {
	// set path
	s.serv.URL.Path = "/transaction/-"
	// read payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Transaction{}, e
	}

	// do request
	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)

	type Data struct {
		Transaction `json:"transaction"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Transaction{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Transaction{}, e
	}

	return res.Data.Transaction, nil
}

// DeleteTransaction deletes a specific transaction from a bank account. It only
// returns an error if an error has occurred, otherwise it will return nil if
// transaction has successfully been deleted.
func (s *Service) DeleteTransaction(UUID uuid.UUID) dutil.Error {
	// set path
	s.serv.URL.Path = "/transaction/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()

	r, e := s.serv.NewRequest("DELETE", s.serv.URL.String(), nil, nil)
	if e != nil {
		return e
	}

	res := struct {
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
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
	// return on successful response
	return nil
}
