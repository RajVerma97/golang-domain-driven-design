package aggregate_test

import (
	"errors"
	"github.com/RajVerma97/golang-domain-driven-design/aggregate"
	"testing"
)

func TestNewCustomer(t *testing.T) {

	type testCase struct {
		testName     string
		name         string
		exptectedErr error
	}

	testCases := []testCase{
		{testName: "Empty name validation", name: "", exptectedErr: aggregate.ErrInvalidPerson},
		{testName: "valid name", name: "Rajneesh kumar", exptectedErr: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.exptectedErr) {
				t.Errorf("expected error %v,got %v", tc.exptectedErr, err)
			}

		})
	}

}
