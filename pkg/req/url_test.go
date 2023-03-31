package req

import (
	"testing"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

// Data structure for
// describing the test case for "configUrl" func
type testCaseConfigUrl struct {
	arg testCaseConfigUrlArg
	exp testCaseConfigUrlExp
}

// Data structure for
// describing the arg value of test case for "configUrl" func
type testCaseConfigUrlArg struct {
	a1 string
	a2 map[string]string
}

// Data structure for
// describing the exp value of test case for "configUrl" func
type testCaseConfigUrlExp struct {
	r1 string
	r2 error
}

// ------------------------------------------------------------------------ //

var (

	//	Test cases for "configUrl" func
	testCasesConfigUrl = []testCaseConfigUrl{
		{
			arg: testCaseConfigUrlArg{
				a1: "https://ya.ru",
				a2: nil,
			},
			exp: testCaseConfigUrlExp{
				r1: "https://ya.ru",
				r2: nil,
			},
		},
		{
			arg: testCaseConfigUrlArg{
				a1: "https://ya.ru",
			},
			exp: testCaseConfigUrlExp{
				r1: "https://ya.ru",
				r2: nil,
			},
		},
		{
			arg: testCaseConfigUrlArg{
				a1: "https://ya.ru",
				a2: map[string]string{},
			},
			exp: testCaseConfigUrlExp{
				r1: "https://ya.ru",
				r2: nil,
			},
		},
		{
			arg: testCaseConfigUrlArg{
				a1: "https://ya.ru",
				a2: map[string]string{"some1": "1", "some2": "2"},
			},
			exp: testCaseConfigUrlExp{
				r1: "https://ya.ru?some1=1&some2=2",
				r2: nil,
			},
		},
	}
)

// ------------------------------------------------------------------------ //

// Test "configUrl" func
func TestConfigUrl(t *testing.T) {

	for i, testCase := range testCasesConfigUrl {
		r1, r2 := configURL(&testCase.arg.a1, &testCase.arg.a2)
		if r1 != testCase.exp.r1 || r2 != testCase.exp.r2 {
			t.Fatalf(
				`
					Test failed:	Test case (%d) for "configUrl" func
									returned (%s, %s) and expected (%s, %s) values do not match
				`,
				i,
				r1,
				r2.Error(),
				testCase.exp.r1,
				testCase.exp.r2.Error(),
			)
		}
	}

}
