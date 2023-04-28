//	closed: true
//	author:	makarov aleksei
//	target:	this package stores code that
//			provides an access point for executing
//			an HTTP(s) request and returning a formatted server response

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Make HTTP(s) request
// with GET, POST method and return response
// in format Response{Code int, Time int64, Body []byte}
func (req Request) GetRespFmt() (*Response, error) {

	reqInitTs := time.Now()

	respRaw, respRawErr := req.fetchRespRaw()
	if respRawErr != nil {
		return nil, respRawErr
	}

	return fetchRespFmt(respRaw, reqInitTs)

}

// ------------------------------------------------------------------------ //
