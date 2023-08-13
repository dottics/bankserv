package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/msp"
	"net/url"
)

// CreateItem creates a new Item for a transaction based on the item data passed
// to the function.
func (s *Service) CreateItem(i Item) (Item, dutil.Error) {
	// set path
	s.URL.Path = "/item"
	// marshal payload
	p, e := dutil.MarshalReader(i)
	if e != nil {
		return Item{}, e
	}
	// do request
	r, e := s.DoRequest("POST", s.URL, nil, nil, p)
	if e != nil {
		return Item{}, e
	}

	type Data struct {
		Item `json:"item"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}
	return res.Data.Item, nil
}

// UpdateItem updates an Item based on the item data passed to the function.
func (s *Service) UpdateItem(i Item) (Item, dutil.Error) {
	// set path
	s.URL.Path = "/item/-"
	// marshal payload
	p, e := dutil.MarshalReader(i)
	if e != nil {
		return Item{}, e
	}

	// do request
	r, e := s.DoRequest("PUT", s.URL, nil, nil, p)
	if e != nil {
		return Item{}, e
	}

	type Data struct {
		Item `json:"item"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}

	return res.Data.Item, nil
}

// DeleteItem deletes a specific item based on the UUID of the item that is
// passed. It returns nil if the item delete was successful and return an error
// if any error has occurred.
func (s *Service) DeleteItem(UUID uuid.UUID) dutil.Error {
	// set path
	s.URL.Path = "/item/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	// do request
	r, e := s.DoRequest("DELETE", s.URL, qs, nil, nil)
	if e != nil {
		return e
	}

	res := struct {
		Errors dutil.Errors `json:"errors"`
	}{}
	_, e = msp.Decode(r, &res)
	if e != nil {
		return e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return e
	}
	return nil
}

// AddItemTags takes an Item UUID and a slice of Tag UUID's then exchanges with
// the microservice to get the new Item. It returns the Item value pointed to
// or and error if an error occurs.
func (s *Service) AddItemTags(itemUUID uuid.UUID, tagsUUID []uuid.UUID) (Item, dutil.Error) {
	// set path
	s.URL.Path = "/item/-/tag"
	payload := struct {
		UUID     uuid.UUID   `json:"uuid"`
		TagUUIDs []uuid.UUID `json:"tag_uuids"`
	}{
		UUID:     itemUUID,
		TagUUIDs: tagsUUID,
	}
	p, e := dutil.MarshalReader(payload)
	if e != nil {
		return Item{}, e
	}
	// do request
	r, e := s.DoRequest("POST", s.URL, nil, nil, p)
	if e != nil {
		return Item{}, e
	}
	// decode the response
	type Data struct {
		Item Item `json:"item"`
	}
	res := struct {
		Data   Data         `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}
	return res.Data.Item, nil
}

// RemoveItemTags takes an Item UUID and Tag UUID's then exchanges with the
// bank microservice to remove the items from the Item. It returns the Item
// value pointed to or an error if an error occurs.
func (s *Service) RemoveItemTags(itemUUID uuid.UUID, tagsUUID []uuid.UUID) (Item, dutil.Error) {
	// set path
	s.URL.Path = "/item/-/tag/-"
	// set query string
	xTagUUID := make([]string, 0)
	for _, tagUUID := range tagsUUID {
		xTagUUID = append(xTagUUID, tagUUID.String())
	}
	qs := url.Values{
		"uuid":      []string{itemUUID.String()},
		"tag_uuids": xTagUUID,
	}
	// do request
	r, e := s.DoRequest("DELETE", s.URL, qs, nil, nil)
	if e != nil {
		return Item{}, e
	}
	// decode the response
	type Data struct {
		Item Item `json:"item"`
	}
	res := struct {
		Data   Data         `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}
	return res.Data.Item, nil
}

func (s *Service) UpdateItemTags(itemUUID uuid.UUID, tagsUUID []uuid.UUID) (Item, dutil.Error) {
	// set path
	s.URL.Path = "/item/-/tag/-"
	// construct the payload object
	payload := struct {
		UUID     uuid.UUID   `json:"uuid"`
		TagUUIDs []uuid.UUID `json:"tag_uuids"`
	}{
		UUID:     itemUUID,
		TagUUIDs: tagsUUID,
	}
	p, e := dutil.MarshalReader(payload)
	if e != nil {
		return Item{}, e
	}
	// do request
	r, e := s.DoRequest("PUT", s.URL, nil, nil, p)
	if e != nil {
		return Item{}, e
	}
	// decode the response
	type Data struct {
		Item Item `json:"item"`
	}
	res := struct {
		Data   Data         `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}
	return res.Data.Item, nil
}
