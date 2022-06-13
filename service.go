package bankserv

import (
	"github.com/johannesscr/micro/msp"
)

type Service msp.Service

var microServiceName string = "bank"

// NewService creates a microservice-package (msp) instance. The msp
// is an instance that loads the environmental variables to be able
// to connect to the specific microservice. The msp contains all the
// implementations to correctly exchange with the microservice.
func NewService(token string) *Service {
	return (*Service)(msp.NewService(token, microServiceName))
}

// SetURL sets the scheme and host of the service. Also makes the service
// a mock-able service with `microtest`
func (s *Service) SetURL(scheme, host string) {
	s.URL.Scheme = scheme
	s.URL.Host = host
}
