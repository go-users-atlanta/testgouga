package testify

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"strconv"
	"testgouga/math"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
	Units float64
	Price float64
}

func (suite *ExampleTestSuite) SetupTest() {
	suite.Units = 1000
	suite.Price = 23.80
}

func (suite *ExampleTestSuite) TestMath() {
	assert.Equal(suite.T(), math.Calc(suite.Units, suite.Price), 25228)
}

func (suite *ExampleTestSuite) TestHttp() {
	client := &http.Client{}
	units := strconv.FormatFloat(suite.Units, 'f', 2, 64)
	price := strconv.FormatFloat(suite.Price, 'f', 2, 64)
	url := fmt.Sprintf("http://localhost:8081/total?units=%s&price=%s", price, units)
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), "25228.00", string(body))
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
