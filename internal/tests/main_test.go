package tests

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

type IntegrationTest struct {
    suite.Suite
}

func (i *IntegrationTest) SetupTest() {

}

func TestIntegrationTestSuite(t *testing.T) {
    suite.Run(t, new(IntegrationTest))
}
