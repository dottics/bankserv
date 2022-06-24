package bankserv

import (
	"fmt"
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/microtest"
	"testing"
)

func TestService_GetTags(t *testing.T) {
	tt := []struct {
		name     string
		exchange *microtest.Exchange
		tags     Tags
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
			tags: Tags{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "get default tags",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tags found","data":{"tags":[{"uuid":"0b57c1f7-1545-4270-b3b0-90c9e6a29a56","user_uuid":null,"organisation_uuid":null,"tag":"Food","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":null,"organisation_uuid":null,"tag":"Sport","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":null,"organisation_uuid":null,"tag":"Equipment","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"}]},"errors":{}}`,
				},
			},
			tags: Tags{
				{
					UUID:             uuid.MustParse("0b57c1f7-1545-4270-b3b0-90c9e6a29a56"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.UUID{},
					Tag:              "Food",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.UUID{},
					Tag:              "Sport",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.UUID{},
					Tag:              "Equipment",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
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

			tags, e := s.GetTags()
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTags(tc.tags, tags) {
				t.Errorf("expected tags %v got %v", tc.tags, tags)
			}
		})
	}
}

func TestService_GetUserTags(t *testing.T) {
	tt := []struct {
		name     string
		UUID     uuid.UUID
		exchange *microtest.Exchange
		tags     Tags
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("f080cf50-f537-4f17-90fa-6a2fc5a4690b"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			tags: Tags{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "get user tags",
			UUID: uuid.MustParse("f080cf50-f537-4f17-90fa-6a2fc5a4690b"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tags found","data":{"tags":[{"uuid":"0b57c1f7-1545-4270-b3b0-90c9e6a29a56","user_uuid":"f080cf50-f537-4f17-90fa-6a2fc5a4690b","organisation_uuid":null,"tag":"Food","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":"f080cf50-f537-4f17-90fa-6a2fc5a4690b","organisation_uuid":null,"tag":"Sport","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":"f080cf50-f537-4f17-90fa-6a2fc5a4690b","organisation_uuid":null,"tag":"Equipment","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"}]},"errors":{}}`,
				},
			},
			tags: Tags{
				{
					UUID:             uuid.MustParse("0b57c1f7-1545-4270-b3b0-90c9e6a29a56"),
					UserUUID:         uuid.MustParse("f080cf50-f537-4f17-90fa-6a2fc5a4690b"),
					OrganisationUUID: uuid.UUID{},
					Tag:              "Food",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.MustParse("f080cf50-f537-4f17-90fa-6a2fc5a4690b"),
					OrganisationUUID: uuid.UUID{},
					Tag:              "Sport",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.MustParse("f080cf50-f537-4f17-90fa-6a2fc5a4690b"),
					OrganisationUUID: uuid.UUID{},
					Tag:              "Equipment",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
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

			tags, e := s.GetUserTags(tc.UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTags(tc.tags, tags) {
				t.Errorf("expected tags %v got %v", tc.tags, tags)
			}
		})
	}
}

func TestService_GetOrganisationTags(t *testing.T) {
	tt := []struct {
		name     string
		UUID     uuid.UUID
		exchange *microtest.Exchange
		tags     Tags
		e        dutil.Error
	}{
		{
			name: "permission required",
			UUID: uuid.MustParse("b1039eaa-3584-4d80-bab0-0ae498d8cbae"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			tags: Tags{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "get default tags",
			UUID: uuid.MustParse("b1039eaa-3584-4d80-bab0-0ae498d8cbae"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tags found","data":{"tags":[{"uuid":"0b57c1f7-1545-4270-b3b0-90c9e6a29a56","user_uuid":null,"organisation_uuid":"b1039eaa-3584-4d80-bab0-0ae498d8cbae","tag":"Food","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":null,"organisation_uuid":"b1039eaa-3584-4d80-bab0-0ae498d8cbae","tag":"Sport","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"},{"uuid":"14376f6d-5bdf-42b8-98af-d6ee9839e5e3","user_uuid":null,"organisation_uuid":"b1039eaa-3584-4d80-bab0-0ae498d8cbae","tag":"Equipment","active":true,"create_date":"2022-05-11T13:04:26Z","update_date":"2022-05-11T13:04:26Z"}]},"errors":{}}`,
				},
			},
			tags: Tags{
				{
					UUID:             uuid.MustParse("0b57c1f7-1545-4270-b3b0-90c9e6a29a56"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("b1039eaa-3584-4d80-bab0-0ae498d8cbae"),
					Tag:              "Food",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("b1039eaa-3584-4d80-bab0-0ae498d8cbae"),
					Tag:              "Sport",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
				},
				{
					UUID:             uuid.MustParse("14376f6d-5bdf-42b8-98af-d6ee9839e5e3"),
					UserUUID:         uuid.UUID{},
					OrganisationUUID: uuid.MustParse("b1039eaa-3584-4d80-bab0-0ae498d8cbae"),
					Tag:              "Equipment",
					Active:           true,
					CreateDate:       timeMustParse("2022-05-11T13:04:26Z"),
					UpdateDate:       timeMustParse("2022-05-11T13:04:26Z"),
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

			tags, e := s.GetOrganisationTags(tc.UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
			if !EqualTags(tc.tags, tags) {
				t.Errorf("expected tags %v got %v", tc.tags, tags)
			}
		})
	}
}

func TestService_CreateTag(t *testing.T) {
	tt := []struct {
		name     string
		tag      Tag
		exchange *microtest.Exchange
		ETag     Tag
		e        dutil.Error
	}{
		{
			name: "permission required",
			tag:  Tag{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			ETag: Tag{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "bad request",
			tag:  Tag{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest: Unable to process request","data":{},"errors":{"user_uuid":["required field"],"organisation_uuid":["required field"]}}`,
				},
			},
			ETag: Tag{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"user_uuid":         {"required field"},
					"organisation_uuid": {"required field"},
				},
			},
		},
		{
			name: "create tag",
			tag: Tag{
				Tag:      "dottics test tag",
				UserUUID: uuid.MustParse("8d4f6969-a54e-4c0a-9292-f87455ab0039"),
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 201,
					Body:   `{"message":"tag created","data":{"tag":{"uuid":"31e7685a-2800-46af-a8ed-0a300ecd45c1","user_uuid":"8d4f6969-a54e-4c0a-9292-f87455ab0039","organisation_uuid":null,"tag":"dottics test tag","active":true,"create_date":"2022-06-21T08:52:23Z","update_date":"2022-06-21T08:52:23Z"}},"errors":{}}`,
				},
			},
			ETag: Tag{
				UUID:       uuid.MustParse("31e7685a-2800-46af-a8ed-0a300ecd45c1"),
				UserUUID:   uuid.MustParse("8d4f6969-a54e-4c0a-9292-f87455ab0039"),
				Tag:        "dottics test tag",
				Active:     true,
				CreateDate: timeMustParse("2022-06-21T08:52:23.000Z"),
				UpdateDate: timeMustParse("2022-06-21T08:52:23.000Z"),
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

			tag, e := s.CreateTag(tc.tag)
			if tc.ETag != tag {
				t.Errorf("expected tag %v got %v", tc.ETag, tag)
			}
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}

func TestService_UpdateTag(t *testing.T) {
	tt := []struct {
		name     string
		tag      Tag
		exchange *microtest.Exchange
		ETag     Tag
		e        dutil.Error
	}{
		{
			name: "permission required",
			tag:  Tag{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 403,
					Body:   `{"message":"Forbidden: Unable to process request","data":{},"errors":{"permission":["Please ensure you have permission"]}}`,
				},
			},
			ETag: Tag{},
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "bad request",
			tag:  Tag{},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 400,
					Body:   `{"message":"BadRequest: Unable to process request","data":{},"errors":{"uuid":["required field"]}}`,
				},
			},
			ETag: Tag{},
			e: &dutil.Err{
				Status: 400,
				Errors: map[string][]string{
					"uuid": {"required field"},
				},
			},
		},
		{
			name: "update tag",
			tag: Tag{
				UUID: uuid.MustParse("715e1420-48b9-4fd7-9fff-140013cded72"),
				Tag:  "new tag name",
			},
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tag updated","data":{"tag":{"uuid":"715e1420-48b9-4fd7-9fff-140013cded72","user_uuid":"ac640bd0-9b33-4e19-a702-abb48b4f3b18","organisation_uuid":null,"tag":"new tag name","active":true,"create_date":"0001-01-01T00:00:00Z","update_date":"0001-01-01T00:00:00Z"}},"errors":{}}`,
				},
			},
			ETag: Tag{
				UUID:             uuid.MustParse("715e1420-48b9-4fd7-9fff-140013cded72"),
				UserUUID:         uuid.MustParse("ac640bd0-9b33-4e19-a702-abb48b4f3b18"),
				OrganisationUUID: uuid.UUID{},
				Tag:              "new tag name",
				Active:           true,
				CreateDate:       timeMustParse("2022-0624T11:11:30Z"),
				UpdateDate:       timeMustParse("2022-0624T11:11:30Z"),
			},
		},
	}

	s := NewService("")
	ms := microtest.MockServer(s.serv)

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			ms.Append(tc.exchange)

			_, e := s.UpdateTag(tc.tag)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}

func TestService_DeleteTag(t *testing.T) {
	tt := []struct {
		name     string
		UUID     uuid.UUID
		exchange *microtest.Exchange
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
			e: &dutil.Err{
				Status: 403,
				Errors: map[string][]string{
					"permission": {"Please ensure you have permission"},
				},
			},
		},
		{
			name: "not found",
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 404,
					Body:   `{"message":"NotFound: Unable to find resource","data":{},"errors":{"tag":["not found"]}}`,
				},
			},
			e: &dutil.Err{
				Status: 404,
				Errors: map[string][]string{
					"tag": {"not found"},
				},
			},
		},
		{
			name: "tag deleted",
			UUID: uuid.MustParse("715e1420-48b9-4fd7-9fff-140013cded72"),
			exchange: &microtest.Exchange{
				Response: microtest.Response{
					Status: 200,
					Body:   `{"message":"tag deleted","data":{},"errors":{}}`,
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

			e := s.DeleteTag(tc.UUID)
			if !dutil.ErrorEqual(tc.e, e) {
				t.Errorf("expected error %v got %v", tc.e, e)
			}
		})
	}
}
