package req

import (
	"net/http"
	"net/url"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Set up client
// config for HTTP(s) request
func (r *Request) setUpClt() error {

	r.clt = &http.Client{
		Timeout: time.Second * 10,
	}

	switch r.PrxU {
	case "":
		return nil
	default:
		u, uErr := url.Parse(r.PrxU)
		if uErr != nil {
			return uErr
		}
		r.clt.Transport = &http.Transport{Proxy: http.ProxyURL(u)}
		return nil
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
