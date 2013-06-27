package parkour

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ParkourSuite struct{}

var _ = Suite(&ParkourSuite{})

func (s *ParkourSuite) TestGetOnlyAppCfg(c *C) {
	c.Check(len(configFiles), Equals, 1)

	c.Check(GetString("aString"), Equals, "www.example.com")
	c.Check(GetInt("aIntNum"), Equals, 42)
	c.Check(GetFloat("aFloatNum"), Equals, 3.14)
	c.Check(GetBool("aBool"), Equals, true)

	defer func() {
		shouldThrowException := false
		if r := recover(); r != nil {
			shouldThrowException = true
		}
		c.Check(shouldThrowException, Equals, true)
	}()

	c.Check(GetString("non-exist-key"), Equals, false)
}

func (s *ParkourSuite) TestGetFullCfg(c *C) {
	c.Check(len(configFiles), Equals, 3)
	c.Check(exist(envCfgPath), Equals, true)
	c.Check(exist(localCfgPath), Equals, true)

	c.Check(GetString("aString"), Equals, "stringInDevCfg")
	c.Check(GetInt("aIntNum"), Equals, 42)
	c.Check(GetFloat("aFloatNum"), Equals, 3.1402107)

	defer func() {
		shouldThrowException := false
		if r := recover(); r != nil {
			shouldThrowException = true
		}
		c.Check(shouldThrowException, Equals, true)
	}()

	c.Check(GetString("non-exist-key"), Equals, false)
}

func (s *ParkourSuite) TestGetSpecifiedEnvCfg(c *C) {
	c.Check(len(configFiles), Equals, 3)
	c.Check(envCfgPath, Equals, "configs/prod.cfg")
	c.Check(exist(envCfgPath), Equals, true)

	c.Check(GetString("aString"), Equals, "stringInProdCfg")
	c.Check(GetInt("aIntNum"), Equals, 42)
	c.Check(GetFloat("aFloatNum"), Equals, 3.1402107)

	defer func() {
		shouldThrowException := false
		if r := recover(); r != nil {
			shouldThrowException = true
		}
		c.Check(shouldThrowException, Equals, true)
	}()

	c.Check(GetString("non-exist-key"), Equals, false)
}
