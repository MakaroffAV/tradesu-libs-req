package req

import "testing"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type (

	//	Data structure for
	//	describing the test case for setUpClt method
	testCaseSetUpClt struct {
		arg testCaseSetUpCltArg
		exp testCaseSetUpCltExp
	}

	//	Data structure for
	//	describing the arg value of test case for setUpClt method
	testCaseSetUpCltArg struct {
		a1 Request
	}

	//	Data structure for
	//	describing the exp value of test case for setUpClt method
	testCaseSetUpCltExp struct {
		r1 error
	}
)

// ------------------------------------------------------------------------ //

var (

	//	Test cases for setUpClt method
	testCasesSetUpClt = []testCaseSetUpClt{
		{
			arg: testCaseSetUpCltArg{
				a1: Request{},
			},
			exp: testCaseSetUpCltExp{
				r1: nil,
			},
		},
		{
			arg: testCaseSetUpCltArg{
				a1: Request{
					PrxU: "socks5://127.0.0.1:9050",
				},
			},
			exp: testCaseSetUpCltExp{
				r1: nil,
			},
		},
	}
)

// ------------------------------------------------------------------------ //

// Test setUpClt method
func TestSetUpClient(t *testing.T) {

	for _, testCase := range testCasesSetUpClt {
		if r1 := testCase.arg.a1.setUpClt(); r1 != testCase.exp.r1 {
			t.Fatalf("here")
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
