package v1

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetCategoryItem(t *testing.T) {
	tests := []struct {
		name       string
		entityUUID uuid.UUID
		category   string
		from       string
		to         string
		exchange   *microtest.Exchange
		results    []ItemDate
		e          dutil.Error
	}{
		{
			name:       "permission required",
			entityUUID: uuid.MustParse("8a6b6251-94b7-4593-8c27-9a50258cfc19"),
			category:   "food",
			from:       "2021-01-01",
			to:         "2021-01-31",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":[],"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			results: []ItemDate{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name:       "items found",
			entityUUID: uuid.MustParse("8a6b6251-94b7-4593-8c27-9a50258cfc19"),
			category:   "food",
			from:       "2021-01-01",
			to:         "2021-01-31",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message":"items found successfully",
						"data":[
							{
								"item":{
									"uuid":"d2b64e51-8b31-4cbd-be90-439ddb33c3b7",
									"description":"item one",
									"sku":"barcode-here",
									"unit":"each",
									"quantity":21.1,
									"amount":12.3,
									"discount":1.23,
									"category":"food",
									"prediction_category":"groceries",
									"active":true
								},
								"date":"2021-01-01T00:00:00Z"
							},
							{
								"item":{
									"uuid":"c1d396ab-4e32-4d1e-9baf-48a10529cf80",
									"description":"item two",
									"sku":"barcode-here",
									"unit":"each",
									"quantity":2,
									"amount":1300,
									"discount":130,
									"category":"food",
									"prediction_category":"groceries",
									"active":true
								},
								"date":"2021-02-01T00:00:00Z"
							}
						],
						"errors":{}
					}`,
				},
			},
			results: []ItemDate{
				{
					Item: Item{
						UUID:               uuid.MustParse("d2b64e51-8b31-4cbd-be90-439ddb33c3b7"),
						Description:        "item one",
						SKU:                "barcode-here",
						Unit:               "each",
						Quantity:           21.1,
						Amount:             12.3,
						Discount:           1.23,
						Category:           "food",
						PredictionCategory: "groceries",
						Active:             true,
					},
					Date: timeMustParse("2021-01-01T00:00:00.000Z"),
				},
				{
					Item: Item{
						UUID:               uuid.MustParse("c1d396ab-4e32-4d1e-9baf-48a10529cf80"),
						Description:        "item two",
						SKU:                "barcode-here",
						Unit:               "each",
						Quantity:           2,
						Amount:             1300,
						Discount:           130,
						Category:           "food",
						PredictionCategory: "groceries",
						Active:             true,
					},
					Date: timeMustParse("2021-02-01T00:00:00.000Z"),
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

			it, e := s.GetCategoryItems(tc.entityUUID, tc.category, tc.from, tc.to)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			for i, it := range it {
				if tc.results[i].Date != it.Date {
					t.Errorf("expected date %v got %v", tc.results[i].Date, it.Date)
				}
				if !EqualItem(tc.results[i].Item, it.Item) {
					t.Errorf("expected item\n%v\ngot\n%v", tc.results[i].Item, it.Item)
				}
			}
		})
	}
}

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
			name: "create item",
			item: Item{
				TransactionUUID:    uuid.MustParse("df7e4020-3863-49f5-ae6c-6604ab64edf5"),
				Description:        "Sasko Brown Bread",
				SKU:                "barcode-here",
				Unit:               "m/s",
				Quantity:           21.1,
				Amount:             12.3,
				Discount:           1.23,
				Category:           "food",
				PredictionCategory: "groceries",
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"item created","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","description":"Sasko Brown Bread","sku":"barcode-here","unit":"m/s","quantity":21.1,"amount":12.3,"discount":1.23,"category":"food","prediction_category":"groceries","tags":[],"active":true,"create_date":"2022-06-19T15:43:01Z","update_date":"2022-06-19T15:43:01Z"}},"errors":{}}`,
				},
			},
			EItem: Item{
				UUID:               uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID:    uuid.UUID{},
				Description:        "Sasko Brown Bread",
				SKU:                "barcode-here",
				Unit:               "m/s",
				Quantity:           21.1,
				Amount:             12.3,
				Discount:           1.23,
				Category:           "food",
				PredictionCategory: "groceries",
				Active:             true,
				CreateDate:         timeMustParse("2022-06-19T15:43:01.000Z"),
				UpdateDate:         timeMustParse("2022-06-19T15:43:01.000Z"),
				Tags:               Tags{},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

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
				UUID:               uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				Description:        "Sasko White Bread",
				SKU:                "barcode-here",
				Unit:               "m/s",
				Quantity:           21.1,
				Amount:             14.3,
				Discount:           1.43,
				Category:           "food",
				PredictionCategory: "groceries",
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"item updated","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","transaction_uuid":"00000000-0000-0000-0000-000000000000","description":"Sasko White Bread","sku":"barcode-here","unit":"m/s","quantity":21.1,"amount":14.3,"discount":1.43,"category":"food","prediction_category":"groceries","tags":[],"active":true,"create_date":"2022-06-19T15:43:01Z","update_date":"2022-06-19T15:43:01Z"}},"errors":{}}`,
				},
			},
			EItem: Item{
				UUID:               uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID:    uuid.UUID{},
				Description:        "Sasko White Bread",
				SKU:                "barcode-here",
				Unit:               "m/s",
				Quantity:           21.1,
				Amount:             14.3,
				Discount:           1.43,
				Category:           "food",
				PredictionCategory: "groceries",
				Active:             true,
				CreateDate:         timeMustParse("2022-06-19T15:43:01.000Z"),
				UpdateDate:         timeMustParse("2022-06-19T15:43:01.000Z"),
				Tags:               Tags{},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

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

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

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

func TestService_AddItemTags(t *testing.T) {
	tests := []struct {
		name     string
		UUID     uuid.UUID
		xTagUUID []uuid.UUID
		exchange *microtest.Exchange
		item     Item
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			item: Item{},
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
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to process request","data":{},"errors":{"item":["not found"]}}`,
				},
			},
			item: Item{},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"item": {"not found"},
				},
			},
		},
		{
			name: "item tags added",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"item tags removed","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","transaction_uuid":"2b5b5fe0-ba22-4f7f-b1de-499472193202","description":"jbl flip se 2","sku":"barcode","unit":"unit","quantity":1,"amount":2400,"discount":0,"active":true,"create_date":"2022-09-24T10:58:32.000Z","update_date":"2022-09-24T10:58:51.000Z","tags":[{"uuid":"11982575-1b9f-4f67-88fa-4a3228119044","tag":"music"},{"uuid":"91499027-ad4d-4cea-b18e-4a8d474e0874","tag":"technology"}]}},"errors":{}}`,
				},
			},
			item: Item{
				UUID:            uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID: uuid.MustParse("2b5b5fe0-ba22-4f7f-b1de-499472193202"),
				Description:     "jbl flip se 2",
				SKU:             "barcode",
				Unit:            "unit",
				Quantity:        1,
				Amount:          2400,
				Discount:        0,
				Tags: Tags{
					Tag{
						UUID: uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
						Tag:  "music",
					},
					Tag{
						UUID: uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
						Tag:  "technology",
					},
				},
				Active:     true,
				CreateDate: timeMustParse("2022-09-24T10:58:32.000Z"),
				UpdateDate: timeMustParse("2022-09-24T10:58:51.000Z"),
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
			item, e := s.AddItemTags(tc.UUID, tc.xTagUUID)
			if !EqualItem(item, tc.item) {
				t.Errorf("expected item %v got %v", tc.item, item)
			}
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}

func TestService_RemoveItemTags(t *testing.T) {
	tests := []struct {
		name     string
		UUID     uuid.UUID
		xTagUUID []uuid.UUID
		exchange *microtest.Exchange
		item     Item
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			item: Item{},
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
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to process request","data":{},"errors":{"item":["not found"]}}`,
				},
			},
			item: Item{},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"item": {"not found"},
				},
			},
		},
		{
			name: "item tags removed",
			UUID: uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("11982575-1b9f-4f67-88fa-4a3228119044"),
				uuid.MustParse("91499027-ad4d-4cea-b18e-4a8d474e0874"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"item tags removed","data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","transaction_uuid":"2b5b5fe0-ba22-4f7f-b1de-499472193202","description":"jbl flip se 2","sku":"barcode","unit":"unit","quantity":1,"amount":2400,"discount":0,"active":true,"create_date":"2022-09-24T10:58:32.000Z","update_date":"2022-09-24T10:58:51.000Z","tags":[{"tag":"one"}]}},"errors":{}}`,
				},
			},
			item: Item{
				UUID:            uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID: uuid.MustParse("2b5b5fe0-ba22-4f7f-b1de-499472193202"),
				Description:     "jbl flip se 2",
				SKU:             "barcode",
				Unit:            "unit",
				Quantity:        1,
				Amount:          2400,
				Discount:        0,
				Tags: Tags{
					Tag{Tag: "one"},
				},
				Active:     true,
				CreateDate: timeMustParse("2022-09-24T10:58:32.000Z"),
				UpdateDate: timeMustParse("2022-09-24T10:58:51.000Z"),
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
			item, e := s.RemoveItemTags(tc.UUID, tc.xTagUUID)
			if !EqualItem(item, tc.item) {
				t.Errorf("expected item %v got %v", tc.item, item)
			}
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}

func TestService_UpdateItemTags(t *testing.T) {
	tests := []struct {
		name     string
		UUID     uuid.UUID
		xTagUUID []uuid.UUID
		exchange *microtest.Exchange
		item     Item
		e        dutil.Error
	}{
		{
			name: "Forbidden",
			UUID: uuid.MustParse("60dd4e8a-df69-4cef-82a5-ea157e3ed797"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
				uuid.MustParse("2d0e158b-79b2-4745-a06f-1fa37cde786b"),
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body: `{
						"message":"Forbidden: Unable to process request",
						"body":{},
						"errors":{"permission":["Please ensure you have permission"]}
					}`,
				},
			},
			item: Item{},
			e: dutil.NewErr(
				403, "permission",
				[]string{"Please ensure you have permission"},
			),
		},
		{
			name: "Not Found",
			UUID: uuid.MustParse("60dd4e8a-df69-4cef-82a5-ea157e3ed797"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
				uuid.MustParse("2d0e158b-79b2-4745-a06f-1fa37cde786b"),
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body: `{
						"message":"NotFound: Unable to find resource",
						"body":{},
						"errors":{"item":["not found"]}
					}`,
				},
			},
			item: Item{},
			e: dutil.NewErr(
				404, "item",
				[]string{"not found"},
			),
		},
		{
			name: "Successful",
			UUID: uuid.MustParse("60dd4e8a-df69-4cef-82a5-ea157e3ed797"),
			xTagUUID: []uuid.UUID{
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
				uuid.MustParse("2d0e158b-79b2-4745-a06f-1fa37cde786b"),
				uuid.MustParse("f6817f27-9974-44db-95c9-201deb0dff98"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message":"item tags updated",
						"data":{"item":{"uuid":"b5b3df71-d3cc-4069-9912-a0e7237aee2b","transaction_uuid":"2b5b5fe0-ba22-4f7f-b1de-499472193202","description":"jbl flip se 2","sku":"barcode","unit":"unit","quantity":1,"amount":2400,"discount":0,"active":true,"create_date":"2022-09-24T10:58:32.000Z","update_date":"2022-09-24T10:58:51.000Z","tags":[{"tag":"one"}]}},
						"errors":{}
					}`,
				},
			},
			e: nil,
			item: Item{
				UUID:            uuid.MustParse("b5b3df71-d3cc-4069-9912-a0e7237aee2b"),
				TransactionUUID: uuid.MustParse("2b5b5fe0-ba22-4f7f-b1de-499472193202"),
				Description:     "jbl flip se 2",
				SKU:             "barcode",
				Unit:            "unit",
				Quantity:        1,
				Amount:          2400,
				Discount:        0,
				Tags: Tags{
					Tag{Tag: "one"},
				},
				Active:     true,
				CreateDate: timeMustParse("2022-09-24T10:58:32.000Z"),
				UpdateDate: timeMustParse("2022-09-24T10:58:51.000Z"),
			},
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)
	defer ms.Server.Close()

	for i, tc := range tests {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)
			item, e := s.UpdateItemTags(tc.UUID, tc.xTagUUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualItem(item, tc.item) {
				t.Errorf("expected item\n%v\ngot\n%v", tc.item, item)
			}
		})
	}
}
