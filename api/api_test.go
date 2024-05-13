package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	suite.Suite
	response *http.Response
}

func (suite *APISuite) TestGetAPI() {
	link := "https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=DEMO_KEY"

	suite.response = GetAPI(link)
	suite.Equal(suite.response.StatusCode, http.StatusOK)
}

func TestGetAPISuite(t *testing.T) {
	suite.Run(t, new(APISuite))
}
