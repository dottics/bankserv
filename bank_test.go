package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetBanks(t *testing.T) {
	type E struct {
		banks Banks
		e     dutil.Error
	}
	tt := []struct {
		name     string
		exchange *microtest.Exchange
		E
	}{
		{
			name: "403 Forbidden",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			E: E{
				banks: Banks{},
				e: &dutil.Err{
					Errors: map[string][]string{
						"permission": {"Please ensure you have permission"},
					},
				},
			},
		},
		{
			name: "200 Successful",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"banks found","data":{"banks":[{"uuid":"2955f13a-f331-4c28-b007-1fc658a61b30","name":"investec","branch_code":"058150","active":true,"create_date":"2022-01-01T12:00:00Z","update_date":"2022-01-01T12:00:00Z"}]},"errors":{}}`,
				},
			},
			E: E{
				banks: Banks{
					// {"uuid":"2955f13a-f331-4c28-b007-1fc658a61b30","name":"investec","branch_code":"058150","active":true,"create_date":"2022-01-01T12:00:00Z","update_date":"2022-01-01T12:00:00Z"}
					{
						UUID:       uuid.MustParse("2955f13a-f331-4c28-b007-1fc658a61b30"),
						Name:       "investec",
						BranchCode: "058150",
						Active:     true,
						CreateDate: timeMustParse("2022-01-01T12:00:00Z"),
						UpdateDate: timeMustParse("2022-01-01T12:00:00Z"),
					},
				},
				e: nil,
			},
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xb, e := s.GetBanks()
			if !dutil.ErrorEqual(e, tc.E.e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if len(xb) != len(tc.banks) {
				t.Errorf("expected banks to have length %d got %d", len(tc.banks), len(xb))
			}
			if len(tc.banks) > 0 {
				if tc.banks[0] != xb[0] {
					t.Errorf("expected bank %v got %v", tc.banks[0], xb[0])
				}
			}
		})
	}
}
