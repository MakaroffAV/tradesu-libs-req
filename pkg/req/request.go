package req

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the HTTP(s) request config
type Request struct {

	//	init params
	Href string
	ArgQ map[string]string
	PrxU string
	Meth string
	ArgH map[string]string
	ArgD string
	RnUa bool

	//	calc params
	url    string
	urlErr error
	req    *http.Request
	reqErr error
	clt    *http.Client
	cltErr error
	res    *http.Response
	resErr error
	rbd    []byte
	rbdErr error
}

// ------------------------------------------------------------------------ //

func (r *Request) setUpReq() error {

	if r.ArgD == "" {
		r.req, r.reqErr = http.NewRequest(r.Meth, r.url, nil)
	} else {
		r.req, r.reqErr = http.NewRequest(r.Meth, r.url, strings.NewReader(r.ArgD))
	}

	if r.RnUa {
		j := userAgents[rand.Intn(len(userAgents))]
		fmt.Println(j)
		print(j)
		r.req.Header.Add("User-Agent", j)
	}

	if len(r.ArgH) != 0 {
		for k, v := range r.ArgH {
			r.req.Header.Add(k, v)
		}
	}

	return nil

}

// ------------------------------------------------------------------------ //

func (r *Request) fetchRes() error {

	if r.urlErr != nil || r.reqErr != nil || r.cltErr != nil {
		return nil
	}

	r.res, r.resErr = r.clt.Do(r.req)
	if r.resErr != nil {
		return nil
	}

	return nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
