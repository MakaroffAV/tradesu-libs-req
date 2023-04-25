package req

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func buildUrl(href *string, query *map[string]string) (*string, error) {

	//	TODO: check if ArgD not passed in struct init
	if query == nil || len(*query) == 0 {
		return href, nil
	} else {
		queryParams := url.Values{}
		for k, v := range *query {
			queryParams.Add(k, v)
		}
		queryString := fmt.Sprintf("%s?%s", *href, queryParams.Encode())
		return &queryString, nil
	}

}

// ------------------------------------------------------------------------ //
func buildClt(tout int64, prxU string) (*http.Client, error) {

	clt := &http.Client{
		Timeout: time.Second * 10,
	}

	if tout != 0 {
		clt.Timeout = time.Second * time.Duration(tout)
	}

	if prxU != "" {
		pUrl, pUrlErr := url.Parse(prxU)
		if pUrlErr != nil {
			return nil, pUrlErr
		}
		clt.Transport = &http.Transport{Proxy: http.ProxyURL(pUrl)}
	}

	return clt, nil

}

// ------------------------------------------------------------------------ //

func configReq(req Request) (*http.Request, error) {

	url, urlErr := buildUrl(&req.Href, &req.ArgQ)
	if urlErr != nil {
		return nil, urlErr
	}

	//	TODO: rename r
	//	TODO: change 3rd argument
	r, rErr := http.NewRequest(req.Meth, *url, nil)
	if rErr != nil {
		return nil, rErr
	}

	if req.RnUa {
		r.Header.Add("User-Agent", userAgents[rand.Intn(len(userAgents))])
	}

	//	TODO: check if ArgD not passed in struct init (still it nil?)
	//	TODO: if it nill, we cant use len func, change it and change it in buildUrl func

	if req.ArgH == nil || len(req.ArgH) != 0 {
		for k, v := range req.ArgH {
			r.Header.Add(k, v)
		}
	}

	return r, nil

}

// ------------------------------------------------------------------------ //

func fetchReq(req Request, req2 *http.Request) (*http.Response, error) {

	clt, cltErr := buildClt(req.Tout, req.PrxU)
	if cltErr != nil {
		return nil, cltErr
	}

	res, resErr := clt.Do(req2)
	if resErr != nil {
		return nil, resErr
	}

	return res, nil
}

// ------------------------------------------------------------------------ //

func fetchResRaw(req Request) (*http.Response, error) {

	req2, reqErr := configReq(req)
	if reqErr != nil {
		return nil, reqErr
	}

	resp, respErr := fetchReq(req, req2)
	if respErr != nil {
		return nil, respErr
	}

	return resp, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
