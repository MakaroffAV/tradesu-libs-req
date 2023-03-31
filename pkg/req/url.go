package req

import (
	"fmt"
	"net/url"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Configure full
// url by request params
// Depend on request struct
// and h is Href, q is Qarg fields
func configURL(h *string, q *map[string]string) (string, error) {

	if q == nil {
		return *h, nil
	} else {
		if len(*q) == 0 {
			return *h, nil
		} else {
			qVars := url.Values{}
			for k, v := range *q {
				qVars.Add(k, v)
			}
			hFull := fmt.Sprintf("%s?%s", *h, qVars.Encode())
			return hFull, nil
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
