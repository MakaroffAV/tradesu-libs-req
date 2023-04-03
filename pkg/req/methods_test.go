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

		resp, respErr := testCase.arg.a1.Get()
		if respErr != nil {
			t.Fatalf("%s, %s", string(resp), respErr.Error())
		}

	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type (

	//	Data structure for describing the
	//	test case for POST method of the Request struct
	testCasePost struct {
		arg testCasePostArg
		exp testCasePostExp
	}

	//	Data structure for describing the
	//	arg value of test case for POST method of the Request struct
	testCasePostArg struct {
		a1 Request
	}

	//	Data structure for describing the
	//	exp value of test case for POST method of the Request struct
	testCasePostExp struct {
		r1 []byte
		r2 error
	}
)

// ------------------------------------------------------------------------ //

var (

	//	Test cases for
	//	POST method of the Request struct
	testCasesPost = []testCasePost{
		{
			arg: testCasePostArg{
				a1: Request{
					Href: "https://www.pochta.ru/suggestions/v2/postoffice.find-nearest-by-postalcode-vacancies",
					Meth: "POST",
					RnUa: true,
					ArgD: `{"postalCode":"662250","filters":[],"limit":1,"radius":100,"offset":0,"currentDateTime":"2023-04-03T10:17:51"}`,
					ArgH: map[string]string{
						"content-type": "application/json;charset=UTF-8",
					},
				},
			},
			exp: testCasePostExp{
				r1: []byte("df"),
				r2: nil,
			},
		},
	}
)

// ------------------------------------------------------------------------ //

// Test GET method
// of the Request struct
func TestPost(t *testing.T) {

	for _, testCase := range testCasesPost {

		resp, respErr := testCase.arg.a1.Post()
		if respErr != nil {
			t.Fatalf("%s, %s", string(resp), respErr.Error())
		}

	}

}
