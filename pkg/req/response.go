//	closed: true
//	author:	makarov aleksei
//	target:	this package stores code
//			that describes working with a customized
//			data structure to describe the HTTP(s) response

package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"io"
	"net/http"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the response from HTTP(s) request
type Response struct {
	Code int
	Time int64
	Body []byte
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Parsing response body from raw
// standard http golang lib, status code, calculate request time
func fetchRespFmt(resRaw *http.Response, reqInitTs time.Time) (*Response, error) {

	respBd, respBdErr := io.ReadAll(resRaw.Body)
	if respBdErr != nil {
		return nil, respBdErr
	}

	return &Response{
		Body: respBd,
		Code: resRaw.StatusCode,
		Time: time.Since(reqInitTs).Milliseconds(),
	}, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
