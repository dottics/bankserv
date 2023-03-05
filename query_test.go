package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/johannesscr/micro/microtest"
	"net/url"
	"testing"
)

func TestService_Query(t *testing.T) {
	tests := []struct {
		name     string
		values   url.Values
		exchange *microtest.Exchange
		e        dutil.Error
		bytes    string
	}{
		{
			name:     "no q defined",
			values:   url.Values{},
			exchange: nil,
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"q": {"q is required"},
				},
			},
			bytes: "",
		},
		{
			name: "only q is defined",
			values: url.Values{
				"q": {"tag_summary"},
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tag summary successful","data":{"tag_one":110,"tag_two":123},"errors":{}}`,
				},
			},
			bytes: `{"message":"tag summary successful","data":{"tag_one":110,"tag_two":123},"errors":{}}`,
			e:     nil,
		},
		{
			name: "multiple url query params defined",
			values: url.Values{
				"q":        {"tag_summary"},
				"uuid":     {"da2e9f64-7be3-4bf5-b468-62d2b2539c77"},
				"end_date": {"2022-10-12"},
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tag summary successful","data":{"tag_one":110,"tag_two":123},"errors":{}}`,
				},
			},
			bytes: `{"message":"tag summary successful","data":{"tag_one":110,"tag_two":123},"errors":{}}`,
			e:     nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			// add the exchange to the queue of the mock service
			ms.Append(tc.exchange)
			xb, e := s.Query(tc.values)
			// ensure the error was what was expected
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// ensure that the response body is what was expected.
			xbRes := string(xb)
			if xbRes != tc.bytes {
				t.Errorf("expected body response '%v' got '%v'", tc.bytes, xbRes)
			}
		})
	}
}
