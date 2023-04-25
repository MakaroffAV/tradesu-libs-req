package req

import (
	"io/ioutil"
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
func fetchResFmt(resRaw *http.Response, reqInitTs time.Time) (*Response, error) {

	resBd, resBdErr := ioutil.ReadAll(resRaw.Body)
	if resBdErr != nil {
		return nil, resBdErr
	}

	return &Response{
		Body: resBd,
		Code: resRaw.StatusCode,
		Time: time.Since(reqInitTs).Milliseconds(),
	}, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
