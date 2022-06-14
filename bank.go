package bankserv

import "github.com/dottics/dutil"

// GetBanks gets all the banks from the bank-service.
func (s *Service) GetBanks() (Banks, dutil.Error) {
	// get access to the msp service
	ms := s.serv
	// create and make request
	ms.URL.Path = "/bank"
	res, e := ms.NewRequest("GET", ms.URL.String(), nil, nil)
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
	_, e = ms.Decode(res, &resp)
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
