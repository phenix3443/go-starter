package greeting

import (
	"github.com/stretchr/testify/suite"
)

type GreetingTestSuite struct {
	suite.Suite
}

func (s *GreetingTestSuite) TestHello() {
	s.Equal("Hello, World", Hello("World"))
}
