package v2

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
		ab       AccountBalance
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
			ab: AccountBalance{},
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
						"message": "account balance retrieved",
						"data": {
							"account_balance": {
								"date": "2020-01-01T00:00:00Z",
								"balance": 0,
								"currency": "GBP"
							}
						},
						"errors": {}
					}`,
				},
			},
			ab: AccountBalance{
				Date:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Balance: 0,
				//Currency: "GBP",
			},
			e: nil,
		},
	}

	s := NewService(Config{
		APIKey: "test",
		Header: http.Header{
			"x-dot-api-key": {"test"},
		},
	})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d: %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			ab, e := s.GetAccountBalance(tc.UUID, tc.date)

			if !dutil.ErrorEqual(e, tc.e) {
				t.Errorf("got %v, want %v", e, tc.e)
			}
			if ab != tc.ab {
				t.Errorf("expected\n%v\n got\n%v", tc.ab, ab)
			}
		})
	}
}

func TestService_CreateAccountBalance(t *testing.T) {
	tests := []struct {
		name     string
		abIn     AccountBalance
		exchange *microtest.Exchange
		abOut    AccountBalance
		e        dutil.Error
	}{
		{
			name: "permission required",
			abIn: AccountBalance{
				AccountUUID: uuid.MustParse("7ca43b18-4684-4c32-b894-c15395190d47"),
				Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Balance:     0,
			},
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
			abOut: AccountBalance{},
			e: &dutil.Err{
				Status: http.StatusUnauthorized,
				Errors: dutil.Errors{
					"permission": []string{"permission required"},
				},
			},
		},
		{
			name: "successful",
			abIn: AccountBalance{
				AccountUUID: uuid.MustParse("bf87be5d-2245-4ae0-8882-c17bc5043346"),
				Date:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Balance:     0,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: http.StatusCreated,
					Body: `{
						"message": "account balance created",
						"data": {
							"account_balance": {
								"account_uuid": "bf87be5d-2245-4ae0-8882-c17bc5043346",
								"date": "2020-05-11T00:05:00Z",
								"balance": 1200,
								"currency": "GBP"
							}
						},
						"errors": {}
					}`,
				},
			},
			abOut: AccountBalance{
				AccountUUID: uuid.MustParse("bf87be5d-2245-4ae0-8882-c17bc5043346"),
				Date:        time.Date(2020, 5, 11, 0, 5, 0, 0, time.UTC),
				Balance:     1200,
			},
			e: nil,
		},
	}

	s := NewService(Config{
		APIKey: "test",
		Header: http.Header{
			"x-dot-api-key": {"test"},
		},
	})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d: %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			ab, e := s.CreateAccountBalance(&tc.abIn)

			if !dutil.ErrorEqual(e, tc.e) {
				t.Errorf("got %v, want %v", e, tc.e)
			}
			if ab != tc.abOut {
				t.Errorf("expected\n%v\n got\n%v", tc.abOut, ab)
			}
		})
	}
}
