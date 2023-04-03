package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Make HTTP(s)
// request with GET method
func (r Request) Get() ([]byte, error) {

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

// ------------------------------------------------------------------------ //

// Make HTTP(s)
// request with POST method
func (r Request) Post() ([]byte, error) {
	return nil, nil
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
