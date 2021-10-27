package errgroup_test

import (
	"errors"
	"testing"

	"github.com/evgendn/errgroup"
)

func TestGroup(t *testing.T) {
	testCases := []struct {
		errs    []error
		errsLen int
	}{
		{errs: []error{}, errsLen: 0},
		{errs: []error{nil}, errsLen: 0},
		{errs: []error{errors.New("error1")}, errsLen: 1},
		{errs: []error{nil, errors.New("error1"), nil}, errsLen: 1},
		{errs: []error{errors.New("error1"), errors.New("error2"), errors.New("error3"), errors.New("error4")}, errsLen: 4},
	}
	for i, tc := range testCases {
		g := &errgroup.Group{}

		for _, err := range tc.errs {
			err := err
			g.Go(func() error { return err })
		}

		errs := g.Wait()
		if tc.errsLen != len(errs) {
			t.Errorf("case %d: expected %d errors, got %d", i+1, tc.errsLen, len(errs))
		}
		if tc.errsLen == 0 && errs != nil {
			t.Errorf("case %d: expecting nil but got %v", i, errs)
		}
	}
}
