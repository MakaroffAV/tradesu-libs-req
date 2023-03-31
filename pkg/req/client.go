package req

import (
	"net/http"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

const (

	//	Default timeout
	//	for HTTP(s) request
	defaultTimeout int8 = 10
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Configure client
// by request params
// Depend on request struct
// and p is PrxU, t is Tout fields
func configClient(p *string, t *int8) (*http.Client, error) {

	var requestClient *http.Client

	//	setup client timeout
	if p != nil {
		requestClient.Timeout = time.Second * time.Duration(*t)
	} else {
		requestClient.Timeout = time.Second * time.Duration(defaultTimeout)
	}

	//	setup client transport

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
