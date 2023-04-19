package req

import "time"

type Response struct {
	Time int64
	Code int
	Body []byte
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Make HTTP(s)
// request with GET method
func (r Request) Get() (*Response, error) {

	a := time.Now()

	if err := r.setUpUrl(); err != nil {
		return nil, err
	}

	if err := r.setUpClt(); err != nil {
		return nil, err
	}

	if err := r.setUpReq(); err != nil {
		return nil, err
	}

	if err := r.fetchRes(); err != nil {
		return nil, err
	}

	if err := r.parseRes(); err != nil {
		return nil, err
	}

	b := time.Now()

	d := b.Sub(a)

	return &Response{Time: d.Milliseconds(), Code: r.res.StatusCode, Body: r.rbd}, nil

}

// ------------------------------------------------------------------------ //

// Make HTTP(s)
// request with POST method
func (r Request) Post() ([]byte, error) {

	if err := r.setUpUrl(); err != nil {
		return nil, err
	}

	if err := r.setUpClt(); err != nil {
		return nil, err
	}

	if err := r.setUpReq(); err != nil {
		return nil, err
	}

	if err := r.fetchRes(); err != nil {
		return nil, err
	}

	if err := r.parseRes(); err != nil {
		return nil, err
	}

	return r.rbd, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
