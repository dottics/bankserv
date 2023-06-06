package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"net/http"
	"testing"
	"time"
)

func TestService_GetAccountBalance(t *testing.T) {
	tests := []struct {
		name     string
		UUID     uuid.UUID
		date     time.Time
		exchange *microtest.Exchange
		ab       *AccountBalance
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			date: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: http.StatusUnauthorized,
					Body: `{
						"message": "Unauthorised: Unable to process request",
						"errors": {
							"permission": ["permission required"]
						}
					}`,
				},
			},
			ab: nil,
			e: &dutil.Err{
				Status: http.StatusUnauthorized,
				Errors: dutil.Errors{
					"permission": []string{"permission required"},
				},
			},
		},
		{
			name: "successful",
			UUID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			date: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: http.StatusOK,
					Body: `{
						"message": "account balance retrieved"
						"data": {
							"account_balance": {
								"uuid": "00000000-0000-0000-0000-000000000000",
								"date": "2020-01-01",
								"balance": 0,
								"currency": "GBP"
							}
						},
						"errors": {},
					}`,
				},
			},
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d: %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			ab, e := s.GetAccountBalance(tc.UUID, tc.date, nil)

			if !dutil.ErrorEqual(e, tc.e) {
				t.Errorf("got %v, want %v", e, tc.e)
			}
			if ab != tc.ab {
				t.Errorf("expected %v got %v", tc.ab, ab)
			}
		})
	}
}
