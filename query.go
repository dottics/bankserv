package bankserv

import (
	"github.com/dottics/dutil"
	"io"
	"log"
	"net/url"
)

// Query is a function that allows for the development of dynamic bank-service
// queries. These queries are to rapidly develop roll-up and views of
// bank-service data.
//
// The aim is to derive useful insights from transactional data.
func (s *Service) Query(values url.Values) ([]byte, dutil.Error) {
	if values.Get("q") == "" {
		e := dutil.NewErr(400, "q", []string{"q is required"})
		return []byte{}, e
	}
	ms := s.serv
	// set the path
	ms.URL.Path = "/query"
	// set the query parameters/values
	ms.URL.RawQuery = values.Encode()
	res, e := ms.NewRequest("GET", ms.URL.String(), nil, nil)
	if e != nil {
		return []byte{}, e
	}

	xb, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}(res.Body)

	if err != nil {
		e := &dutil.Err{
			Status: 500,
			Errors: map[string][]string{
				"read": {err.Error()},
			},
		}
		return []byte{}, e
	}
	return xb, e
}