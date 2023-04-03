package req

import (
	"fmt"
	"net/url"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func (r *Request) setUpUrl() error {

	if len(r.ArgQ) == 0 {
		r.url = r.Href
	} else {
		qVar := url.Values{}
		for k, v := range r.ArgQ {
			qVar.Add(k, v)
		}
		r.url = fmt.Sprintf("%s?%s", r.Href, qVar.Encode())
	}

	return nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
