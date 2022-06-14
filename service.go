package bankserv

import (
	"github.com/johannesscr/micro/msp"
)

type Service struct {
	serv *msp.Service
}

const microServiceName string = "bank"

// NewService creates a microservice-package (msp) instance. The msp
// is an instance that loads the environmental variables to be able
// to connect to the specific microservice. The msp contains all the
// implementations to correctly exchange with the microservice.
func NewService(token string) *Service {
	s := &Service{
		serv: msp.NewService(token, microServiceName),
	}
	return s
}
