package hello

import (
	"github.com/stretchr/testify/suite"
)

type HelloTestSuite struct {
	suite.Suite
}

func (s *HelloTestSuite) TestHello() {
	s.Equal("Hello, World", Greet("World"))
}
