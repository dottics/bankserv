package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetUserBankAccounts(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		bankAccounts BankAccounts
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
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "user not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"user":["not found"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"user": {"not found"},
				},
			},
		},
		{
			name: "user has no bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"user bank accounts found","data":{"bank_accounts":[]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e:            nil,
		},
		{
			name: "user has bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"user bank accounts found","data":{"bank_accounts":[{"uuid":"318b052a-7911-4e09-a76d-f6e6a18c6fcd","user_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","organisation_uuid":null,"account_number":"012345678911","active":true,"create_date":"2022-05-17T04:35:23.000Z","update_date":"2022-05-17T04:35:23.000Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","user_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","organisation_uuid":null,"account_number":"012345678912","active":true,"create_date":"2022-05-17T06:53:32.000Z","update_date":"2022-05-17T06:53:32.000Z"}]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{
				BankAccount{
					UUID:             uuid.MustParse("318b052a-7911-4e09-a76d-f6e6a18c6fcd"),
					UserUUID:         uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					OrganisationUUID: uuid.UUID{},
					AccountNumber:    "012345678911",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
				},
				BankAccount{
					UUID:             uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					UserUUID:         uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					OrganisationUUID: uuid.UUID{},
					AccountNumber:    "012345678912",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8")
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)
			xba, e := s.GetUserBankAccounts(UUID)
			// test that the errors are equal
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the bank accounts
			if len(xba) != len(tc.bankAccounts) {
				t.Errorf(
					"expected bank accounts to have length %d got %d",
					len(tc.bankAccounts),
					len(xba),
				)
			}
			// to check that bank account are equal
			if len(xba) > 0 && tc.bankAccounts[0] != xba[0] {
				t.Errorf(
					"expected bank account %v got %v",
					tc.bankAccounts[0],
					xba[0],
				)
			}
		})
	}
}

func TestService_GetOrganisationBankAccounts(t *testing.T) {
	tt := []struct {
		name         string
		exchange     *microtest.Exchange
		bankAccounts BankAccounts
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
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "organisation not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"organisation":["not found"]}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e: &dutil.Err{
				Errors: map[string][]string{
					"organisation": {"not found"},
				},
			},
		},
		{
			name: "organisation has no bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"organisation bank accounts found","data":{"bank_accounts":[]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{},
			e:            nil,
		},
		{
			name: "organisation has bank accounts",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"organisation bank accounts found","data":{"bank_accounts":[{"uuid":"318b052a-7911-4e09-a76d-f6e6a18c6fcd","organisation_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","user_uuid":null,"account_number":"012345678911","active":true,"create_date":"2022-05-17T04:35:23.000Z","update_date":"2022-05-17T04:35:23.000Z"},{"uuid":"d25ac3b1-0a8f-43a3-8da1-d2f22a814a82","organisation_uuid":"ef50ad5f-539a-454d-bb49-c2e3123eaba8","user_uuid":null,"account_number":"012345678912","active":true,"create_date":"2022-05-17T06:53:32.000Z","update_date":"2022-05-17T06:53:32.000Z"}]},"errors":{}}`,
				},
			},
			bankAccounts: BankAccounts{
				BankAccount{
					UUID:             uuid.MustParse("318b052a-7911-4e09-a76d-f6e6a18c6fcd"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					AccountNumber:    "012345678911",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T04:35:23.000Z"),
				},
				BankAccount{
					UUID:             uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8"),
					AccountNumber:    "012345678912",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
					UpdateDate:       timeMustParse("2022-05-17T06:53:32.000Z"),
				},
			},
			e: nil,
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		UUID := uuid.MustParse("ef50ad5f-539a-454d-bb49-c2e3123eaba8")

		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			xba, e := s.GetOrganisationBankAccounts(UUID)
			// test that the errors are equal
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			// test the bank accounts
			if len(xba) != len(tc.bankAccounts) {
				t.Errorf(
					"expected bank accounts to have length %d got %d",
					len(tc.bankAccounts),
					len(xba),
				)
			}
			// to check that bank account are equal
			if len(xba) > 0 && tc.bankAccounts[0] != xba[0] {
				t.Errorf(
					"expected bank account %v got %v",
					tc.bankAccounts[0],
					xba[0],
				)
			}
		})
	}
}
