package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_CreateItem(t *testing.T) {
	tt := []struct {
		name     string
		item     Item
		exchange *microtest.Exchange
		EItem    Item
		e        dutil.Error
	}{
		{
			name: "permission required",
			item: Item{
				TransactionUUID: uuid.MustParse("df7e4020-3863-49f5-ae6c-6604ab64edf5"),
				Description:     "Sasko Brown Bread",
				SKU:             2,
				Amount:          12.3,
				Discount:        1.23,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			EItem: Item{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "bad request",
			item: Item{
				TransactionUUID: uuid.MustParse("df7e4020-3863-49f5-ae6c-6604ab64edf5"),
				Description:     "Sasko Brown Bread",
				SKU:             2,
				Amount:          12.3,
				Discount:        1.23,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest","data":{},"errors":{"transaction":["not found"]}}`,
				},
			},
			EItem: Item{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"transaction": {"not found"},
				},
			},
		},
		{
			name: "create transaction",
			item: Item{
				TransactionUUID: uuid.MustParse("df7e4020-3863-49f5-ae6c-6604ab64edf5"),
				Description:     "Sasko Brown Bread",
				SKU:             2,
				Amount:          12.3,
				Discount:        1.23,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"item created","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","description":"Sasko Brown Bread","sku":2,"amount":12.3,"discount":1.23,"tags":[],"active":true,"create_date":"2022-06-19T15:43:01Z","update_date":"2022-06-19T15:43:01Z"}},"errors":{}}`,
				},
			},
			EItem: Item{
				UUID:            uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID: uuid.UUID{},
				Description:     "Sasko Brown Bread",
				SKU:             2,
				Amount:          12.3,
				Discount:        1.23,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-19T15:43:01.000Z"),
				UpdateDate:      timeMustParse("2022-06-19T15:43:01.000Z"),
				Tags:            Tags{},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			it, e := s.CreateItem(tc.item)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualItem(tc.EItem, it) {
				t.Errorf("expected item %v got %v", tc.EItem, it)
			}
		})
	}
}

func TestService_UpdateItem(t *testing.T) {
	tt := []struct {
		name     string
		item     Item
		exchange *microtest.Exchange
		EItem    Item
		e        dutil.Error
	}{
		{},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

		})
	}
}

func TestService_DeleteItem(t *testing.T) {
	tt := []struct {
		name     string
		UUID     uuid.UUID
		exchange *microtest.Exchange
		e        dutil.Error
	}{
		{},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

		})
	}
}
