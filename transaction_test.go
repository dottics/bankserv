package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetAccountTransactions(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		transactions Transactions
		e            dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			transactions: Transactions{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "account not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"account":["not found"]}}`,
				},
			},
			transactions: Transactions{},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"account": {"not found"},
				},
			},
		},
		{
			name: "transactions found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"","data":{"transactions":[{"uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","account_uuid":"032203af-6002-4abc-9982-73c577add8df","date":"2022-06-18T15:26:22Z","description":"SUPERSPAR JEFFREYS BAYEASTERN CAPEZA","items":[{"uuid":null,"transaction_uuid":null,"description":"two","sku":0,"amount":0,"discount":0,"tags":[{"uuid":null,"tag":"two","active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":true,"create_date":"2022-06-18T15:49:58Z","update_date":"2022-06-18T15:50:06Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","account_uuid":"032203af-6002-4abc-9982-73c577add8df","date":"2022-06-18T15:26:22Z","description":"GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB","items":[{"uuid":null,"transaction_uuid":null,"description":"one","sku":0,"amount":0,"discount":0,"tags":[{"uuid":null,"tag":"one","active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":true,"create_date":"2022-06-18T15:28:34Z","update_date":"2022-06-18T15:29:32Z"}]},"errors":{}}`,
				},
			},
			transactions: Transactions{
				{
					UUID:        uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
					AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
					Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
					Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Active:      true,
					CreateDate:  timeMustParse("2022-06-18T15:49:58.000Z"),
					UpdateDate:  timeMustParse("2022-06-18T15:50:06.000Z"),
					Items: Items{
						{
							Description: "two",
							Tags: Tags{
								{Tag: "two"},
							},
						},
					},
				},
				{
					UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
					Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
					Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
					Active:      true,
					CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
					UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
					Items: Items{
						{
							Description: "one",
							Tags: Tags{
								{Tag: "one"},
							},
						},
					},
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)
	UUID := uuid.MustParse("09bc087c-85b8-4c05-b14b-835cdbd9825c")

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xt, e := s.GetAccountTransactions(UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTransactions(tc.transactions, xt) {
				t.Errorf("expected transactions %v got %v", tc.transactions, xt)
			}
		})
	}
}

func TestService_GetUserTransactions(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		transactions Transactions
		e            dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			transactions: Transactions{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "no transactions",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"","data":{"transactions":[]},"errors":{}}`,
				},
			},
			transactions: Transactions{},
			e:            nil,
		},
		{
			name: "user has transactions",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"","data":{"transactions":[{"uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","account_uuid":"032203af-6002-4abc-9982-73c577add8df","date":"2022-06-18T15:26:22Z","description":"SUPERSPAR JEFFREYS BAYEASTERN CAPEZA","items":[{"uuid":null,"transaction_uuid":null,"description":"two","sku":0,"amount":0,"discount":0,"tags":[{"uuid":null,"tag":"two","active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":true,"create_date":"2022-06-18T15:49:58Z","update_date":"2022-06-18T15:50:06Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","account_uuid":"032203af-6002-4abc-9982-73c577add8df","date":"2022-06-18T15:26:22Z","description":"GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB","items":[{"uuid":null,"transaction_uuid":null,"description":"one","sku":0,"amount":0,"discount":0,"tags":[{"uuid":null,"tag":"one","active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":false,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}],"active":true,"create_date":"2022-06-18T15:28:34Z","update_date":"2022-06-18T15:29:32Z"}]},"errors":{}}`,
				},
			},
			transactions: Transactions{
				{
					UUID:        uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
					AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
					Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
					Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Active:      true,
					CreateDate:  timeMustParse("2022-06-18T15:49:58.000Z"),
					UpdateDate:  timeMustParse("2022-06-18T15:50:06.000Z"),
					Items: Items{
						{
							Description: "two",
							Tags: Tags{
								{Tag: "two"},
							},
						},
					},
				},
				{
					UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
					Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
					Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
					Active:      true,
					CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
					UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
					Items: Items{
						{
							Description: "one",
							Tags: Tags{
								{Tag: "one"},
							},
						},
					},
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)
	UUID := uuid.MustParse("09bc087c-85b8-4c05-b14b-835cdbd9825c")

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xt, e := s.GetUserTransactions(UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTransactions(tc.transactions, xt) {
				t.Errorf("expected transactions %v got %v", tc.transactions, xt)
			}
		})
	}
}

func TestService_CreateTransaction(t *testing.T) {
	tt := []struct {
		name         string
		transaction  Transaction
		exchange     *microtest.Exchange
		ETransaction Transaction
		e            dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			transaction: Transaction{
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
			},
			ETransaction: Transaction{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "bad request",
			transaction: Transaction{
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest: Unable to process request","data":{},"errors":{"uuid":["required field"]}}`,
				},
			},
			ETransaction: Transaction{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"uuid": {"required field"},
				},
			},
		},
		{
			name: "create transaction",
			transaction: Transaction{
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"transaction created","data":{"transaction":{"uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8","account_uuid":"032203af-6002-4abc-9982-73c577add8df","date":"2022-06-18T15:26:22Z","description":"SUPERSPAR JEFFREYS BAYEASTERN CAPEZA","items":[],"active":true,"create_date":"2022-06-18T15:49:58Z","update_date":"2022-06-18T15:50:06Z"}},"errors":{}}`,
				},
			},
			ETransaction: Transaction{
				UUID:        uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:49:58.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:50:06.000Z"),
				Items:       Items{},
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

			txn, e := s.CreateTransaction(tc.transaction)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTransaction(tc.ETransaction, txn) {
				t.Errorf("expected transaction %v got %v", tc.ETransaction, txn)
			}
		})
	}
}

func TestService_UpdateTransaction(t *testing.T) {
	tt := []struct {
		name         string
		transaction  Transaction
		exchange     *microtest.Exchange
		ETransaction Transaction
		e            dutil.Error
	}{
		{
			name: "permission required",
			transaction: Transaction{
				UUID:        uuid.MustParse("7f408ea2-f5e5-4547-8f74-c33fe75c3081"),
				Description: "UPDATE*TXN#TO#NEW#DESCR",
				Date:        timeMustParse("2022-06-19T13:27:19.000Z"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			ETransaction: Transaction{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name:        "bad request",
			transaction: Transaction{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest: Unable to process request","data":{},"errors":{"uuid":["required field"]}}`,
				},
			},
			ETransaction: Transaction{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"uuid": {"required field"},
				},
			},
		},
		{
			name: "update transaction",
			transaction: Transaction{
				UUID:        uuid.MustParse("7f408ea2-f5e5-4547-8f74-c33fe75c3081"),
				Description: "UPDATE*TXN#TO#NEW#DESCR",
				Date:        timeMustParse("2022-06-19T13:27:19.000Z"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"transaction updated","data":{"transaction":{"uuid":"7f408ea2-f5e5-4547-8f74-c33fe75c3081","account_uuid":"6dedbdf5-84ad-435e-8a2f-26d929e18116","date":"2022-06-19T13:27:19Z","description":"UPDATE*TXN#TO#NEW#DESCR","items":[],"active":true,"create_date":"2022-06-18T15:49:58Z","update_date":"2022-06-18T15:50:06Z"}},"errors":{}}`,
				},
			},
			ETransaction: Transaction{
				UUID:        uuid.MustParse("7f408ea2-f5e5-4547-8f74-c33fe75c3081"),
				AccountUUID: uuid.MustParse("6dedbdf5-84ad-435e-8a2f-26d929e18116"),
				Description: "UPDATE*TXN#TO#NEW#DESCR",
				Date:        timeMustParse("2022-06-19T13:27:19.000Z"),
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:49:58.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:50:06.000Z"),
				Items:       Items{},
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

			tx, e := s.UpdateTransaction(tc.transaction)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTransaction(tc.ETransaction, tx) {
				t.Errorf("expected transaction %v got %v", tc.ETransaction, tx)
			}
		})
	}
}

func TestService_DeleteTransaction(t *testing.T) {
	tt := []struct {
		name     string
		UUID     uuid.UUID
		exchange *microtest.Exchange
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("d6b14d91-c5a4-4a02-a6e3-b4dd27cb663c"),
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
			UUID: uuid.UUID{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest: Unable to process request","data":{},"errors":{"transaction":["not found"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"transaction": {"not found"},
				},
			},
		},
		{
			name: "transaction deleted",
			UUID: uuid.MustParse("d6b14d91-c5a4-4a02-a6e3-b4dd27cb663c"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"transaction deleted","data":{},"errors":{}}`,
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

			e := s.DeleteTransaction(tc.UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}
