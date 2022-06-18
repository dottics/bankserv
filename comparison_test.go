package bankserv

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestEqualTags(t *testing.T) {
	tt := []struct {
		name string
		a    Tags
		b    Tags
		o    bool
	}{
		{
			name: "different lengths",
			a: Tags{
				{Tag: "one"},
				{Tag: "two"},
			},
			b: Tags{
				{Tag: "three"},
			},
			o: false,
		},
		{
			name: "different order",
			a: Tags{
				{Tag: "one"},
				{Tag: "two"},
			},
			b: Tags{
				{Tag: "two"},
				{Tag: "one"},
			},
			o: false,
		},
		{
			name: "same order",
			a: Tags{
				{Tag: "one"},
				{Tag: "two"},
			},
			b: Tags{
				{Tag: "one"},
				{Tag: "two"},
			},
			o: true,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := EqualTags(tc.a, tc.b)
			if tc.o != o {
				t.Errorf("expected output %t got %t", tc.o, o)
			}
		})
	}
}

func TestEqualItem(t *testing.T) {
	tt := []struct {
		name string
		a    Item
		b    Item
		o    bool
	}{
		{
			name: "different UUID",
			a: Item{
				UUID: uuid.MustParse("51d51af6-aeee-4ddb-8f02-d379f6b8673f"),
			},
			b: Item{
				UUID: uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
			},
			o: false,
		},
		{
			name: "different TransactionUUID",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
			},
			o: false,
		},
		{
			name: "different Description",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
			},
			o: false,
		},
		{
			name: "different Amount",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          236.189,
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
			},
			o: false,
		},
		{
			name: "different Discount",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        236.189,
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        192.12,
			},
			o: false,
		},
		{
			name: "different SKU",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.95,
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
			},
			o: false,
		},
		{
			name: "different Active",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.95,
				Active:          true,
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          false,
			},
			o: false,
		},
		{
			name: "different CreateDate",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.95,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:54:11.000Z"),
			},
			o: false,
		},
		{
			name: "different UpdateDate",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.95,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:55:01.000Z"),
			},
			o: false,
		},
		{
			name: "same item no tags",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags:            Tags{},
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags:            Tags{},
			},
			o: true,
		},
		{
			name: "same item different tags",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags: Tags{
					{Tag: "one"},
				},
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags:            Tags{},
			},
			o: false,
		},
		{
			name: "same item same tags",
			a: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags: Tags{
					{Tag: "one"},
				},
			},
			b: Item{
				UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
				TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
				Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
				Amount:          37.6,
				Discount:        3.45,
				SKU:             3.15,
				Active:          true,
				CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
				UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
				Tags: Tags{
					{Tag: "one"},
				},
			},
			o: true,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := EqualItem(tc.a, tc.b)
			if tc.o != o {
				t.Errorf("expected output %t got %t", tc.o, o)
			}
		})
	}
}

func TestEqualItems(t *testing.T) {
	tt := []struct {
		name string
		a    Items
		b    Items
		o    bool
	}{
		{
			name: "different lengths",
			a: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
				{
					UUID:            uuid.MustParse("afa3a1f9-1822-48b4-874b-1c5ff974af30"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "two"},
					},
				},
			},
			b: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
			},
			o: false,
		},
		{
			name: "different items",
			a: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
				{
					UUID:            uuid.MustParse("afa3a1f9-1822-48b4-874b-1c5ff974af30"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "two"},
					},
				},
			},
			b: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
				{
					UUID:            uuid.MustParse("2fa81848-229a-464a-8881-f04046d4f430"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "two"},
					},
				},
			},
			o: false,
		},
		{
			name: "same items",
			a: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
				{
					UUID:            uuid.MustParse("2fa81848-229a-464a-8881-f04046d4f430"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "two"},
					},
				},
			},
			b: Items{
				{
					UUID:            uuid.MustParse("a03d4ac5-1d5b-465c-9e0a-c7658912c47d"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "one"},
					},
				},
				{
					UUID:            uuid.MustParse("2fa81848-229a-464a-8881-f04046d4f430"),
					TransactionUUID: uuid.MustParse("d441f6ba-4e40-477a-aa46-916b2dc56bb5"),
					Description:     "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
					Amount:          37.6,
					Discount:        3.45,
					SKU:             3.15,
					Active:          true,
					CreateDate:      timeMustParse("2022-06-18T13:53:57.000Z"),
					UpdateDate:      timeMustParse("2022-06-18T13:54:45.000Z"),
					Tags: Tags{
						{Tag: "two"},
					},
				},
			},
			o: true,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := EqualItems(tc.a, tc.b)
			if tc.o != o {
				t.Errorf("expected output %t got %t", tc.o, o)
			}
		})
	}
}

func TestEqualTransaction(t *testing.T) {
	tt := []struct {
		name string
		a    Transaction
		b    Transaction
		o    bool
	}{
		{
			name: "zero transaction",
			a:    Transaction{},
			b:    Transaction{},
			o:    true,
		},
		{
			name: "different UUID",
			a: Transaction{
				UUID: uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
			},
			b: Transaction{
				UUID: uuid.MustParse("e6b7f986-307c-4147-a34e-f924790799bb"),
			},
			o: false,
		},
		{
			name: "different Account UUID",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("5ed51d15-d033-4a4f-9a5a-a060bb9fc467"),
			},
			o: false,
		},
		{
			name: "different Date",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:35.000Z"),
			},
			o: false,
		},
		{
			name: "different Description",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "SUPERSPAR JEFFREYS BAYEASTERN CAPEZA",
			},
			o: false,
		},
		{
			name: "different Active",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      false,
			},
			o: false,
		},
		{
			name: "different CreateDate",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:50.000Z"),
			},
			o: false,
		},
		{
			name: "different UpdateDate",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:48.000Z"),
			},
			o: false,
		},
		{
			name: "same Transaction",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
			},
			o: true,
		},
		{
			name: "different Items",
			a: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
				Items: Items{
					{Description: "one"},
				},
			},
			b: Transaction{
				UUID:        uuid.MustParse("d25ac3b1-0a8f-43a3-8da1-d2f22a814a82"),
				AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
				Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
				Description: "GOOGLE *GOOGLE STORAGEG.CO/HELPPAY#GB",
				Active:      true,
				CreateDate:  timeMustParse("2022-06-18T15:28:34.000Z"),
				UpdateDate:  timeMustParse("2022-06-18T15:29:32.000Z"),
				Items: Items{
					{Description: "two"},
				},
			},
			o: false,
		},
		{
			name: "different Item Tags",
			a: Transaction{
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
			b: Transaction{
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
							{Tag: "two"},
						},
					},
				},
			},
			o: false,
		},
		{
			name: "equal down to tags",
			a: Transaction{
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
			b: Transaction{
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
			o: true,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := EqualTransaction(tc.a, tc.b)
			if tc.o != o {
				t.Errorf("expected output %t got %t", tc.o, o)
			}
		})
	}
}

func TestEqualTransactions(t *testing.T) {
	tt := []struct {
		name string
		a    Transactions
		b    Transactions
		o    bool
	}{
		{
			name: "different lengths",
			a: Transactions{
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
			},
			b: Transactions{
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
			o: false,
		},
		{
			name: "different transaction order",
			a: Transactions{
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
			},
			b: Transactions{
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
			o: false,
		},
		{
			name: "different transaction elements",
			a: Transactions{
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
					UUID:        uuid.MustParse("5ed51d15-d033-4a4f-9a5a-a060bb9fc467"),
					AccountUUID: uuid.MustParse("032203af-6002-4abc-9982-73c577add8df"),
					Date:        timeMustParse("2022-06-18T15:26:22.000Z"),
					Description: "ANOTHER TXN DESCRIPTION",
					Active:      true,
					CreateDate:  timeMustParse("2022-06-18T15:51:03.000Z"),
					UpdateDate:  timeMustParse("2022-06-18T15:52:07.000Z"),
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
			b: Transactions{
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
			o: false,
		},
		{
			name: "same transactions",
			a: Transactions{
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
			b: Transactions{
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
			o: true,
		},
	}

	for i, tc := range tt {
		name := fmt.Sprintf("%d %s", i, tc.name)
		t.Run(name, func(t *testing.T) {
			o := EqualTransactions(tc.a, tc.b)
			if tc.o != o {
				t.Errorf("expected output %t got %t", tc.o, o)
			}
		})
	}
}
