//	closed: true
//	author:	makarov aleksei
//	target:	this package stores code
//			that describes the process of setting up
//			request, request's client config, fetching raw response

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the HTTP(s) request
type Request struct {

	//	init fields
	Href string
	ArgQ map[string]string
	PrxU string
	Meth string
	ArgH map[string]string
	ArgD string
	Tout int64

	//	calc fields
	clt *http.Client
	req *http.Request
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Generate full url
// by passed domain and path value
func (r *Request) getUrl() string {

	switch len(r.ArgQ) {
	case 0:
		{
			return r.Href
		}
	default:
		{
			qParams := url.Values{}
			for k, v := range r.ArgQ {
				qParams.Add(k, v)
			}
			return fmt.Sprintf("%s?%s", r.Href, qParams.Encode())
		}
	}

}

// ------------------------------------------------------------------------ //

// Set up client
// config for HTTP(s) request
func (r *Request) setClt() error {

	r.clt = &http.Client{}

	switch r.Tout {
	case 0:
		{
			r.clt.Timeout = time.Second * 10
		}
	default:
		{
			r.clt.Timeout = time.Second * time.Duration(r.Tout)
		}
	}

	switch r.PrxU {
	case "":
		{
			return nil
		}
	default:
		{
			prxUrl, prxUrlErr := url.Parse(r.PrxU)
			if prxUrlErr != nil {
				return prxUrlErr
			}
			r.clt.Transport = &http.Transport{Proxy: http.ProxyURL(prxUrl)}
		}
	}

	return nil

}

// ------------------------------------------------------------------------ //

// Set up request
// config for HTTP(s) request
func (r *Request) setReq() error {

	var (
		req    *http.Request
		reqErr error
	)

	switch r.ArgD {
	case "":
		{
			req, reqErr = http.NewRequest(r.Meth, r.getUrl(), nil)
			if reqErr != nil {
				return reqErr
			}
		}
	default:
		{
			req, reqErr = http.NewRequest(r.Meth, r.getUrl(), strings.NewReader(r.ArgD))
			if reqErr != nil {
				return reqErr
			}
		}
	}

	switch len(r.ArgH) {
	case 0:
		{
			req.Header.Add("User-Agent", userAgents[rand.Intn(len(userAgents))])
		}
	default:
		{
			req.Header.Add("User-Agent", userAgents[rand.Intn(len(userAgents))])
			for k, v := range r.ArgH {
				req.Header.Add(k, v)
			}
		}
	}

	r.req = req

	return nil

}

// ------------------------------------------------------------------------ //

// Building client and request config,
// fetching HTTP(s) request raw response
func (r Request) fetchRespRaw() (*http.Response, error) {

	if cltErr := r.setClt(); cltErr != nil {
		return nil, cltErr
	}

	if reqErr := r.setReq(); reqErr != nil {
		return nil, reqErr
	}

	return r.clt.Do(r.req)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
