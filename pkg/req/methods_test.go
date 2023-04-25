package req

import "testing"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type (

	//	Data structure for describing the
	//	test case for Get method of the Request struct
	testCaseGet struct {
		arg testCaseGetArg
		exp testCaseGetExp
	}

	//	Data structure for describing the
	//	arg value of test case for GET method of the Request struct
	testCaseGetArg struct {
		a1 Request
	}

	//	Data structure for describing the
	//	exp value of test case for GET method of the Request struct
	testCaseGetExp struct {
		r1 []byte
		r2 error
	}
)

// ------------------------------------------------------------------------ //

var (

	//	Test cases for
	//	GET method of the Request struct
	testCasesGet = []testCaseGet{
		{
			arg: testCaseGetArg{
				a1: Request{
					Href: "https://api.ipify.org",
					ArgQ: map[string]string{"format": "json"},
					Meth: "GET",
					RnUa: true,
				},
			},
			exp: testCaseGetExp{
				r1: []byte("df"),
				r2: nil,
			},
		},
	}
)

// ------------------------------------------------------------------------ //

// Test GET method
// of the Request struct
func TestGet(t *testing.T) {

	for _, testCase := range testCasesGet {

		resp, respErr := Do(testCase.arg.a1)
		if respErr != nil {
			t.Fatalf("%s, %s", string(resp.Body), respErr.Error())
		}

	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
