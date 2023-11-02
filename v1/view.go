package v1

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/msp"
	"net/url"
)

// GetCategoryMonthTotals gets all the category month totals for a specific
// entity based on the entity's UUID. The function returns a slice of
// CategoryMonthTotal if the UUID is valid, otherwise it returns the dutil.Error
// that is encountered.
func (s *Service) GetCategoryMonthTotals(entityUUID uuid.UUID, query url.Values) ([]CategoryMonthTotal, dutil.Error) {
	s.URL.Path = fmt.Sprintf("view/%s/category/month", entityUUID.String())

	// do request
	r, e := s.DoRequest("GET", s.URL, query, nil, nil)
	if e != nil {
		return []CategoryMonthTotal{}, e
	}

	type data struct {
		Result []CategoryMonthTotal `json:"result"`
	}
	res := struct {
		Data   data         `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}

	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return []CategoryMonthTotal{}, e
	}

	// check response status
	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return []CategoryMonthTotal{}, e
	}
	// return category month totals on success
	return res.Data.Result, nil
}
