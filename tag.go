package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
)

// CreateTag creates a new Tag based on the tag data passed to the method.
func (s *Service) CreateTag(t Tag) (Tag, dutil.Error) {
	// set path
	s.serv.URL.Path = "/tag"
	// marshal payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Tag{}, e
	}
	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
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
	_, e = s.serv.Decode(r, &res)
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
	s.serv.URL.Path = "/tag/-"
	// marshal payload
	p, e := dutil.MarshalReader(t)
	if e != nil {
		return Tag{}, e
	}

	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)
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
	_, e = s.serv.Decode(r, &res)
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
	return nil
}
