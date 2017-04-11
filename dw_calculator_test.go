package bodyfat

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// EPSILON for floating-point equality tolerance
var EPSILON = 0.000001

// TestCase defines each test-case's inputs and expected result
type TestCase struct {
	p  Person       // person attributes
	m  Measurements // body measurements (in millimeters)
	bf float64      // expected body-fat result
}

// TestDWCalculator test the 'durning and womersly' body-fat calculator
func TestDWCalculator(t *testing.T) {
	// define the test cases
	testCases := []TestCase{
		{
			p: Person{
				Age:    16,
				Gender: GenderMale,
			},
			m: Measurements{
				Waist:    5.0,
				FrontArm: 8.0,
				BackArm:  7.0,
				Back:     11.0,
			},
			bf: 18.126964,
		},
		{
			p: Person{
				Age:    23,
				Gender: GenderFemale,
			},
			m: Measurements{
				Waist:    15.0,
				FrontArm: 18.0,
				BackArm:  17.0,
				Back:     21.0,
			},
			bf: 31.909048,
		},
		{
			p: Person{
				Age:    35,
				Gender: GenderMale,
			},
			m: Measurements{
				Waist:    15.0,
				FrontArm: 18.0,
				BackArm:  17.0,
				Back:     21.0,
			},
			bf: 25.279902,
		},
		{
			p: Person{
				Age:    42,
				Gender: GenderMale,
			},
			m: Measurements{
				Waist:    6.0,
				FrontArm: 4.0,
				BackArm:  6.0,
				Back:     16.0,
			},
			bf: 18.46630217,
		},
		{
			p: Person{
				Age:    52,
				Gender: GenderFemale,
			},
			m: Measurements{
				Waist:    4.0,
				FrontArm: 2.0,
				BackArm:  4.0,
				Back:     7.0,
			},
			bf: 19.400740,
		},
	}

	// create a 'durnin and wormersly' body-fat calculator
	calc := NewDWCalculator()

	// iterate through each test-case
	for _, tc := range testCases {
		// run the next sub-test
		name := fmt.Sprintf("Gender '%s', Age '%d'", tc.p.Gender, tc.p.Age)
		t.Run(name, func(t *testing.T) {
			// calculate the body-fat percentage
			bf := calc.Calculate(&tc.p, &tc.m)

			// compare to the expected body-fat
			assert.InEpsilon(t, tc.bf, bf, EPSILON,
				"got '%f', expected '%f'", bf, tc.bf)
		})
	}
}
