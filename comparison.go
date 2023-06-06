package bankserv

// EqualTags is a comparison function for Tags or two slices of Tag. Slices
// are not directly comparable, therefore a comparison function is needed.
// Check that the slices have the same length, and the same entry in each
// position. Returns a true for each equality and false if there are
// differences.
func EqualTags(a, b Tags) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		ti := b[i]
		if t != ti {
			return false
		}
	}
	return true
}

// EqualItem is a comparison function for a non-comparable struct, since the
// struct contains a slice of Tag, therefore, compare each field and compare
// the Tags.
func EqualItem(a, b Item) bool {
	if a.UUID != b.UUID {
		return false
	}
	if a.TransactionUUID != b.TransactionUUID {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	if a.Amount != b.Amount {
		return false
	}
	if a.Discount != b.Discount {
		return false
	}
	if a.SKU != b.SKU {
		return false
	}
	if a.Unit != b.Unit {
		return false
	}
	if a.Quantity != b.Quantity {
		return false
	}
	if a.Active != b.Active {
		return false
	}
	if !a.CreateDate.Equal(b.CreateDate) {
		return false
	}
	if !a.UpdateDate.Equal(b.UpdateDate) {
		return false
	}
	if !EqualTags(a.Tags, b.Tags) {
		return false
	}
	return true
}

// EqualItems reports whether two slices of Item have the same items in the
// same position in the slice.
func EqualItems(a, b Items) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if !EqualItem(ai, b[i]) {
			return false
		}
	}
	return true
}

// EqualTransaction reports whether a and b represents the same Transaction.
//
// Cannot do direct equality, because the Items' field is a slice that does not
// support direct equality.
func EqualTransaction(a, b Transaction) bool {
	if a.UUID != b.UUID {
		return false
	}
	if a.ExternalID != b.ExternalID {
		return false
	}
	if a.AccountUUID != b.AccountUUID {
		return false
	}
	if a.Date != b.Date {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	if a.BusinessName != b.BusinessName {
		return false
	}
	if a.Debit != b.Debit {
		return false
	}
	if a.Credit != b.Credit {
		return false
	}
	if a.Amount != b.Amount {
		return false
	}
	if a.Active != b.Active {
		return false
	}
	if a.CreateDate != b.CreateDate {
		return false
	}
	if a.UpdateDate != b.UpdateDate {
		return false
	}
	if !EqualItems(a.Items, b.Items) {
		return false
	}
	return true
}

// EqualTransactions reports whether a and b are the same Transactions in the
// same order within the slice.
func EqualTransactions(a, b *Transactions) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(*a) != len(*b) {
		return false
	}
	for i, ai := range *a {
		if !EqualTransaction(ai, (*b)[i]) {
			return false
		}
	}
	return true
}
