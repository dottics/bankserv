package v2

import (
	"github.com/johannesscr/micro/msp"
	"net/http"
	"net/url"
)

type Service struct {
	msp.Service
}

const microServiceName string = "bank"

// Config defined the configuration for the `bankserv` package.
type Config struct {
	UserToken string
	APIKey    string
	Values    url.Values
	Header    http.Header
}

// NewService creates a microservice-package (msp) instance. The msp
// is an instance that loads the environmental variables to be able
// to connect to the specific microservice. The msp contains all the
// implementations to correctly exchange with the microservice.
//
// Once a user has logged in the user receives a token, that token needs
// to be passed to the new service so that the token can be added to headers
// to gain access to the microservice.
func NewService(config Config) *Service {
	s := &Service{
		Service: *msp.NewService(msp.Config{
			Name:      microServiceName,
			UserToken: config.UserToken,
			APIKey:    config.APIKey,
			Values:    config.Values,
			Header:    config.Header,
		}),
	}
	return s
}
