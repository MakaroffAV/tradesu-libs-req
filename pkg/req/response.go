package req

import "io/ioutil"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func (r *Request) parseRes() error {

	if r.res == nil {
		return r.resErr
	}

	if r.res.StatusCode != 200 {
		return r.resErr
	}

	r.rbd, r.rbdErr = ioutil.ReadAll(r.res.Body)
	if r.rbdErr != nil {
		return r.rbdErr
	}

	return nil
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
