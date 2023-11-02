package v1

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"net/url"
	"testing"
)

func TestService_GetCategoryMonthTotals(t *testing.T) {
	tests := []struct {
		name                 string
		entityUUID           uuid.UUID
		query                url.Values
		exchange             *microtest.Exchange
		xCategoryMonthTotals []CategoryMonthTotal
		e                    dutil.Error
	}{
		{
			name:       "403 no permission",
			entityUUID: uuid.MustParse("74c5eb22-f574-42a3-8b80-384719c097c6"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"errors":{"permission":["no permission"]}}`,
				},
			},
			xCategoryMonthTotals: []CategoryMonthTotal{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"no permission"},
				},
			},
		},
		{
			name:       "200 not found",
			entityUUID: uuid.MustParse("c77e891c-66be-4828-b3be-5ac07cdd8194"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"category view found","data":{"result":[]},"errors":{}}`,
				},
			},
			xCategoryMonthTotals: []CategoryMonthTotal{},
			e:                    nil,
		},
		{
			name:       "200 successful",
			entityUUID: uuid.MustParse("1d912b91-4e4a-43fe-abab-1b6500334a0c"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message":"category view found",
						"data":{
							"result":[
								{
									"entity_uuid":"1d912b91-4e4a-43fe-abab-1b6500334a0c",
									"category":"bread",
									"year":2022,
									"month":3,
									"total":200.01
								},
								{
									"entity_uuid":"1d912b91-4e4a-43fe-abab-1b6500334a0c",
									"category":"food",
									"year":2022,
									"month":4,
									"total":1300.98
								}
							]
						},
						"errors":{}
					}`,
				},
			},
			xCategoryMonthTotals: []CategoryMonthTotal{
				{
					Category: "bread",
					Year:     2022,
					Month:    3,
					Total:    200.01,
				},
				{
					Category: "food",
					Year:     2022,
					Month:    4,
					Total:    1300.98,
				},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xCategoryMonthTotals, e := s.GetCategoryMonthTotals(tc.entityUUID, tc.query)

			if !dutil.ErrorEqual(e, tc.e) {
				t.Errorf("expected error %v, got %v", tc.e, e)
			}

			if len(xCategoryMonthTotals) != len(tc.xCategoryMonthTotals) {
				t.Errorf("expected %d category month totals got %d", len(tc.xCategoryMonthTotals), len(xCategoryMonthTotals))
			}

			for i, x := range tc.xCategoryMonthTotals {
				xi := xCategoryMonthTotals[i]
				if x != xi {
					t.Errorf("expected %#v got %#v", x, xi)
				}
			}
		})
	}
}
