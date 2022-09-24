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
				SKU:             "barcode-here",
				Unit:            "m/s",
				Quantity:        21.1,
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
			name: "not found",
			item: Item{
				TransactionUUID: uuid.MustParse("df7e4020-3863-49f5-ae6c-6604ab64edf5"),
				Description:     "Sasko Brown Bread",
				SKU:             "barcode-here",
				Unit:            "m/s",
				Quantity:        21.1,
				Amount:          12.3,
				Discount:        1.23,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"transaction":["not found"]}}`,
				},
			},
			EItem: Item{},
			e: &dutil.Err{
				Status: 404,
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
				SKU:             "barcode-here",
				Unit:            "m/s",
				Quantity:        21.1,
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
				SKU:             "barcode-here",
				Unit:            "m/s",
				Quantity:        21.1,
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
		{
			name: "permission required",
			item: Item{
				UUID:        uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				Description: "Sasko Brown Bread",
				SKU:         "barcode-here",
				Unit:        "m/s",
				Quantity:    21.1,
				Amount:      12.3,
				Discount:    1.23,
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
			name: "not found",
			item: Item{
				UUID:        uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				Description: "Sasko Brown Bread",
				SKU:         "barcode-here",
				Unit:        "m/s",
				Quantity:    21.1,
				Amount:      12.3,
				Discount:    1.23,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"item":["not found"]}}`,
				},
			},
			EItem: Item{},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"item": {"not found"},
				},
			},
		},
		{
			name: "item updated",
			item: Item{
				UUID:        uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				Description: "Sasko White Bread",
				SKU:         "barcode-here",
				Unit:        "m/s",
				Quantity:    21.1,
				Amount:      14.3,
				Discount:    1.43,
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"item updated","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","transaction_uuid":"00000000-0000-0000-0000-000000000000","description":"Sasko White Bread","sku":4,"amount":14.3,"discount":1.43,"tags":[],"active":true,"create_date":"2022-06-19T15:43:01Z","update_date":"2022-06-19T15:43:01Z"}},"errors":{}}`,
				},
			},
			EItem: Item{
				UUID:            uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID: uuid.UUID{},
				Description:     "Sasko White Bread",
				SKU:             "barcode-here",
				Unit:            "m/s",
				Quantity:        21.1,
				Amount:          14.3,
				Discount:        1.43,
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

			it, e := s.UpdateItem(tc.item)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualItem(tc.EItem, it) {
				t.Errorf("expected item %v got %v", tc.EItem, it)
			}
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
		{
			name: "permission required",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "not found",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"item":["not found"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"item": {"not found"},
				},
			},
		},
		{

			name: "delete item",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"item deleted","data":{},"errors":{}}`,
				},
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

			e := s.DeleteItem(tc.UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}
