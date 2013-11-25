package prettytest

import (
	"github.com/remogatto/prettytest"
	"testgouga/math"
	"testing"
)

// Start of setup
type testSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}

func (t *testSuite) TestCalculations() {
	var units float64 = 1000
	var price float64 = 23.80
	var expectedResult float64 = 25228
	t.Equal(math.Calc(units, price), expectedResult)
}
