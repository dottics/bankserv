package v1

import (
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/msp"
)

// GetBanks gets all the banks from the bank-service.
func (s *Service) GetBanks() (Banks, dutil.Error) {
	// get access to the msp service
	// create and make request
	s.URL.Path = "/bank"
	res, e := s.DoRequest("GET", s.URL, nil, nil, nil)
	if e != nil {
		return Banks{}, e
	}

	// response structure
	resp := struct {
		Message string `json:"message"`
		Data    struct {
			Banks Banks `json:"banks"`
		} `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}

	// decode the response
	_, e = msp.Decode(res, &resp)
	if e != nil {
		return Banks{}, e
	}

	// check response for error
	if res.StatusCode != 200 {
		e := &dutil.Err{
			Status: res.StatusCode,
			Errors: resp.Errors,
		}
		return Banks{}, e
	}

	return resp.Data.Banks, nil
}
