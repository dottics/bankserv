package v1

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"github.com/johannesscr/micro/msp"
	"net/url"
)

// GetTags gets all the system default tags generic to all users and
// organisations for the classification of items.
func (s *Service) GetTags() (Tags, dutil.Error) {
	// set path
	s.URL.Path = "/tag"
	r, e := s.DoRequest("GET", s.URL, nil, nil, nil)
	if e != nil {
		return Tags{}, e
	}

	type Data struct {
		Tags `json:"tags"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Tags{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Tags{}, e
	}
	return res.Data.Tags, nil
}

// GetEntityTags gets all the tags for an entity based on the entity's UUID. If
// there are no user tags an empty slice is returned. If an error occurs
// an empty slice is returned and a nonzero error.
func (s *Service) GetEntityTags(UUID uuid.UUID) (Tags, dutil.Error) {
	// set path
	s.URL.Path = "/tag/entity/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	// do request
	r, e := s.DoRequest("GET", s.URL, qs, nil, nil)
	if e != nil {
		return Tags{}, e
	}

	type Data struct {
		Tags `json:"tags"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Tags{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Tags{}, e
	}
	return res.Data.Tags, nil
}

// CreateTag creates a new Tag based on the tag data passed to the method.
func (s *Service) CreateTag(t Tag) (Tag, dutil.Error) {
	// set path
	s.URL.Path = "/tag"
	// marshal payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Tag{}, e
	}
	// do request
	r, e := s.DoRequest("POST", s.URL, nil, nil, p)
	if e != nil {
		return Tag{}, e
	}

	type Data struct {
		Tag `json:"tag"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Tag{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Tag{}, e
	}
	return res.Data.Tag, nil
}

// UpdateTag updates a tag's data with the tag data passed to the method.
func (s *Service) UpdateTag(t Tag) (Tag, dutil.Error) {
	// set path
	s.URL.Path = "/tag/-"
	// marshal payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Tag{}, e
	}

	r, e := s.DoRequest("PUT", s.URL, nil, nil, p)
	if e != nil {
		return Tag{}, e
	}

	type Data struct {
		Tag `json:"tag"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = msp.Decode(r, &res)
	if e != nil {
		return Tag{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Tag{}, e
	}

	return res.Data.Tag, nil
}

// DeleteTag deletes a tag based on the UUID passed to the method.
func (s *Service) DeleteTag(UUID uuid.UUID) dutil.Error {
	// set path
	s.URL.Path = "/tag/-"
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
	// decode response
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
