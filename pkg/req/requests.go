package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the HTTP(s) request config
type Request struct {
	//	for configuration url
	Href string
	Qarg map[string]string
	//	for configuration client
	PrxU string
	Tout int8
	//	for configuration request
}

// ------------------------------------------------------------------------ //

// Creating HTTP(s)
// request with GET method
func (r Request) Get() ([]byte, error) {

	return nil, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
