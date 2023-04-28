//	closed: true
//	author:	makarov aleksei
//	target:	this package stores code that tests
//			the provision of an access point to
//			execute an HTTP request and return a formatted server response

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package req

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"testing"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type (

	//	Data structure for describing the
	//	test case for GetRespFmt method of the Request struct
	testCaseGetRespFmt struct {
		arg testCaseGetRespFmtArg
		exp testCaseGetRespFmtExp
	}

	//	Data structure for describing the
	//	arg value of test case for GetRespFmt method of the Request struct
	testCaseGetRespFmtArg struct {
		a1 Request
	}

	//	Data structure for describing the
	//	exp value of test case for GetRespFmt method of the Request struct
	testCaseGetRespFmtExp struct {
		r1 Response
		r2 error
	}
)

// ------------------------------------------------------------------------ //

var (

	//	Test cases for
	//	GetRespFmt method of the Request struct
	testCasesGetRespFmt = []testCaseGetRespFmt{
		{
			arg: testCaseGetRespFmtArg{
				a1: Request{
					Href: "https://api.ipify.org",
					ArgQ: map[string]string{"format": "json"},
					Meth: "GET",
					Tout: 5,
					PrxU: "http://aekejp:bFmiYPeRBf@185.103.165.234:24531",
				},
			},
			exp: testCaseGetRespFmtExp{
				r1: Response{
					Time: 0,
					Code: 200,
					Body: nil,
				},
				r2: nil,
			},
		},
	}
)

// ------------------------------------------------------------------------ //

// Test GetRespFmt
// method of the Request struct
func TestGetRespFmt(t *testing.T) {

	for i, testCase := range testCasesGetRespFmt {

		resp, respErr := testCase.arg.a1.GetRespFmt()
		if respErr != testCase.exp.r2 {
			t.Fatalf(
				`
					Test failed:	Test "TestGetRespFmt" case %d
									returned (%+v, %s) and expected (%+v, %s) values do not match
				`,
				i,
				resp,
				respErr,
				testCase.exp.r1,
				testCase.exp.r2,
			)
		}

	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
