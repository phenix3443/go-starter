package hello

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	suite.Suite
}

func (s *HelloTestSuite) TestHello() {
	s.Equal("Hello, World", Greet("World"))
}

func TestHelloSuite(t *testing.T) {
	suite.Run(t, new(HelloTestSuite))
}
