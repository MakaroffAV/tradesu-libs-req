package req

import (
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the HTTP(s) request
type Request struct {

	//	init params
	Href string
	ArgQ map[string]string
	PrxU string
	Meth string
	ArgH map[string]string
	ArgD string
	RnUa bool
	Tout int64
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Make HTTP(s)
// request with GET, POST ... method
func Do(reqCnf Request) (*Response, error) {

	reqInitTs := time.Now()

	resRaw, resRawErr := fetchResRaw(reqCnf)
	if resRawErr != nil {
		return nil, resRawErr
	}

	resFmt, resFmtErr := fetchResFmt(resRaw, reqInitTs)
	if resFmtErr != nil {
		return nil, resFmtErr
	}

	return resFmt, nil

}

// ------------------------------------------------------------------------ //
