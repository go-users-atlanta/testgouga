package gocheck

import (
	. "launchpad.net/gocheck"
	"testgouga/math"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{ Units, Price float64 }

var _ = Suite(&MySuite{1000, 23.80})

type FixturedSuite struct{ Units, Price float64 }

var _ = Suite(&FixturedSuite{})

//Basic Test

func (s *MySuite) TestCalculations(c *C) {
	c.Check(math.Calc(s.Units, s.Price), Equals, float64(25228))
}

// Test with Fixtures
//func (s *SuiteType) SetUpSuite(c *C) - Run once when the suite starts running.
//func (s *SuiteType) SetUpTest(c *C) - Run before each test or benchmark starts running.
//func (s *SuiteType) TearDownTest(c *C) - Run after each test or benchmark runs.
//func (s *SuiteType) TearDownSuite(c *C) - Run once after all tests or benchmarks have finished running.
func (s *FixturedSuite) SetUpTest(c *C) {
	s.Units = 1000
	s.Price = 10.2
}

func (s *FixturedSuite) TestCalculations(c *C) {
	c.Check(math.Calc(s.Units, s.Price), Equals, float64(10812))
}

// Test benchmarking

func (s *FixturedSuite) BenchmarkCalculations(c *C) {
	for i := 0; i < c.N; i++ {
		math.Calc(s.Units, s.Price)
	}
}
