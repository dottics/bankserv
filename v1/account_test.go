package v1

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

var entityAccount = Account{
	UUID:       uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
	BankUUID:   uuid.MustParse("344a4aa5-1935-4c28-973e-d74247d8db91"),
	EntityUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
	Number:     "098765432109",
	Name:       "private bank account",
	Alias:      "personal account",
}
var entityUpdateAccount = UpdateAccount{
	UUID:       uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
	BankUUID:   uuid.MustParse("344a4aa5-1935-4c28-973e-d74247d8db91"),
	EntityUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
	Number:     "098765432109",
	Name:       "private bank account",
	Alias:      "personal account",
}
var bank = Bank{
	UUID:       uuid.MustParse("344a4aa5-1935-4c28-973e-d74247d8db91"),
	Name:       "investec",
	BranchCode: "580001",
	SwiftCode:  "INVXXJJ",
	Active:     true,
	CreateDate: timeMustParse("2022-06-17T21:57:12.000Z"),
	UpdateDate: timeMustParse("2022-06-17T21:57:12.000Z"),
}

func TestService_GetEntityAccounts(t *testing.T) {
	tt := []struct {
		name     string
		exchange *microtest.Exchange
		accounts Accounts
		e        dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			accounts: Accounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "entity not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body: `{
						"message":"NotFound: Unable to find resource",
						"data":{},
						"errors":{"entity":["not found"]}
					}`,
				},
			},
			accounts: Accounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"entity": {"not found"},
				},
			},
		},
		{
			name: "entity has no bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"entity bank accounts found","data":{"accounts":[]},"errors":{}}`,
				},
			},
			accounts: Accounts{},
			e:        nil,
		},
		{
			name: "entity has bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body: `{
						"message":"entity accounts found",
						"data":{
							"accounts":[
								{
									"uuid":"318b052a-7911-4e09-a76d-f6e6a18c6fcd",
									"bank": {
										"uuid":"344a4aa5-1935-4c28-973e-d74247d8db91",
										"name":"investec",
										"branch_code":"580001",
										"swift_code":"INVXXJJ",
										"active":true,
										"create_date":"2022-06-17T21:57:12.000Z",
										"update_date":"2022-06-17T21:57:12.000Z"
									},
									"entity_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8",
									"number":"012345678911",
									"name":"private bank account",
									"alias":"personal private bank account",
									"balance": {
										"balance": 1200.00,
										"date": "2022-06-17T21:57:12.000Z"
									},
									"active":true,
									"create_date":"2022-05-17T04:35:23.000Z",
									"update_date":"2022-05-17T04:35:23.000Z"
								},
								{
									"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82",
									"bank": {
										"uuid":"344a4aa5-1935-4c28-973e-d74247d8db91",
										"name":"investec",
										"branch_code":"580001",
										"swift_code":"INVXXJJ",
										"active":true,
										"create_date":"2022-06-17T21:57:12.000Z",
										"update_date":"2022-06-17T21:57:12.000Z"
									},
									"entity_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8",
									"number":"012345678912",
									"name":"savings account",
									"alias":"personal savings account",
									"balance": {
										"balance": 1100.00,
										"date": "2022-04-17T21:57:12.000Z"
									},
									"active":true,
									"create_date":"2022-05-17T06:53:32.000Z",
									"update_date":"2022-05-17T06:53:32.000Z"
								}
							]
						},
						"errors":{}
					}`,
				},
			},
			accounts: Accounts{
				Account{
					UUID:       uuid.MustParse("318b052a-7911-4e09-a76d-f6e6a18c6fcd"),
					Bank:       bank,
					EntityUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					Number:     "012345678911",
					Name:       "private bank account",
					Alias:      "personal private bank account",
					Balance: AccountBalance{
						Balance: 1200.00,
						Date:    timeMustParse("2022-06-17T21:57:12.000Z"),
					},
					Active:     true,
					CreateDate: timeMustParse("2022-05-17T04:35:23.000Z"),
					UpdateDate: timeMustParse("2022-05-17T04:35:23.000Z"),
				},
				Account{
					UUID:       uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					Bank:       bank,
					EntityUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					Number:     "012345678912",
					Name:       "savings account",
					Alias:      "personal savings account",
					Balance: AccountBalance{
						Balance: 1100.00,
						Date:    timeMustParse("2022-04-17T21:57:12.000Z"),
					},
					Active:     true,
					CreateDate: timeMustParse("2022-05-17T06:53:32.000Z"),
					UpdateDate: timeMustParse("2022-05-17T06:53:32.000Z"),
				},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8")
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)
			xba, e := s.GetEntityAccounts(UUID)
			// test that the errors are equal
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the bank accounts
			if len(xba) != len(tc.accounts) {
				t.Errorf(
					"expected accounts to have length %d got %d",
					len(tc.accounts),
					len(xba),
				)
			}
			// to check that bank account are equal
			if len(xba) > 0 && tc.accounts[0] != xba[0] {
				t.Errorf(
					"expected account\n%v\ngot\n%v",
					tc.accounts[0],
					xba[0],
				)
			}
		})
	}
}

func TestService_CreateAccount(t *testing.T) {
	tt := []struct {
		name     string
		account  Account // payload data
		exchange *microtest.Exchange
		eAccount Account // expected bank account
		e        dutil.Error
	}{
		{
			name:    "permission required",
			account: entityAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			eAccount: Account{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name:    "entity not found",
			account: entityAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: unable to find resource","data":{},"errors":{"user":["not found"]}}`,
				},
			},
			eAccount: Account{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"user": {"not found"},
				},
			},
		},
		{
			name:    "create entity bank account",
			account: entityAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body: `{
						"message":"account create",
						"data":{
							"account":{
								"uuid":"e6b7f986-307c-4147-a34e-f924790799bb",
								"bank": {
									"uuid":"344a4aa5-1935-4c28-973e-d74247d8db91",
									"name":"investec",
									"branch_code":"580001",
									"swift_code":"INVXXJJ",
									"active":true,
									"create_date":"2022-06-17T21:57:12.000Z",
									"update_date":"2022-06-17T21:57:12.000Z"
								},
								"entity_uuid":"e4bd194d-41e7-4f27-a4a8-161685a9b8b8",
								"number":"098765432109",
								"name":"private bank account",
								"alias":"personal account",
								"number":"098765432109",
								"active":true,
								"create_date":"2022-06-17T21:57:12.000Z",
								"update_date":"2022-06-17T21:57:12.000Z"
							}
						},
						"errors":{}
					}`,
				},
			},
			eAccount: Account{
				UUID:       uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
				Bank:       bank,
				EntityUUID: uuid.MustParse("e4bd194d-41e7-4f27-a4a8-161685a9b8b8"),
				Number:     "098765432109",
				Name:       "private bank account",
				Alias:      "personal account",
				Active:     true,
				CreateDate: timeMustParse("2022-06-17T21:57:12.000Z"),
				UpdateDate: timeMustParse("2022-06-17T21:57:12.000Z"),
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			b, e := s.CreateAccount(tc.account)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if b != tc.eAccount {
				t.Errorf("expected account \n%v\ngot\n%v", tc.eAccount, b)
			}
		})
	}
}

func TestService_UpdateAccount(t *testing.T) {
	tt := []struct {
		name     string
		account  UpdateAccount // payload data
		exchange *microtest.Exchange
		eAccount Account // expected bank account
		e        dutil.Error
	}{
		{
			name:    "permission required",
			account: entityUpdateAccount,
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			eAccount: Account{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			b, e := s.UpdateAccount(tc.account)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if b != tc.eAccount {
				t.Errorf("expected bank account %v got %v", tc.eAccount, b)
			}
		})
	}
}

func TestService_DeleteAccount(t *testing.T) {
	tt := []struct {
		name     string
		exchange *microtest.Exchange
		UUID     uuid.UUID
		e        dutil.Error
	}{
		{
			name: "permission required",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "delete bank account",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"bank account deleted","data":{},"errors":{}}`,
				},
			},
			e: nil,
		},
	}

	s := NewService(Config{})
	ms := microtest.MockServer(s)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb")
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			e := s.DeleteAccount(UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}
