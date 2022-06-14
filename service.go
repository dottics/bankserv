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
//
// Once a user has logged in the user receives a token, that token needs
// to be passed to the new service so that the token can be added to headers
// to gain access to the microservice.
func NewService(token string) *Service {
	s := &Service{
		serv: msp.NewService(token, microServiceName),
	}
	return s
}
