package logging

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type LevelSuite struct {}
var _ = Suite(&LevelSuite{})

func (a *LevelSuite) TestPaddedString(c *C) {
	c.Check(ERROR.PaddedString(10), Equals, "ERROR     ")
	c.Check(INFO.PaddedString(0), Equals, "INFO    ")
	c.Check(WARNING.PaddedString(3), Equals, "WAR")
}

func (a *LevelSuite) TestJSON(c *C) {
	exp := map[LogLevel]string{
		NONE: `""`,
		LOG: `"LOG"`,
		CRITICAL: `"CRITICAL"`,
		ERROR: `"ERROR"`,
		WARNING: `"WARNING"`,
		INFO: `"INFO"`,
		DEBUG: `"DEBUG"`,
		IGNORED: `"IGNORED"`,
	}
	for k, v := range exp {
		data, err := json.Marshal(k)
		c.Check(err, IsNil)
		c.Check(string(data), Equals, v)
		var ll LogLevel
		err = json.Unmarshal(data, &ll)
		c.Check(err, IsNil)
		c.Check(ll, Equals, k)
	}
	var ll LogLevel
	err := json.Unmarshal([]byte("1"), &ll)
	c.Check(err, ErrorMatches, "^.*can't unmarshal log level.*$")
}

func (a *LevelSuite) TestText(c *C) {
	exp := map[string]LogLevel{
		"": NONE,
		"LOG": LOG,
		"CRITICAL": CRITICAL,
		"ERROR": ERROR,
		"WARNING": WARNING,
		"INFO": INFO,
		"DEBUG": DEBUG,
		"IGNORED": IGNORED,
	}
	for k, v := range exp {
		ll := NONE
		llp := &ll
		err := llp.UnmarshalText(k)
		c.Check(err, IsNil)
		c.Check(ll, Equals, v)
	}
	ll := NONE
	llp := &ll
	err := llp.UnmarshalText("NONE")
	c.Check(err, ErrorMatches, "unknown log level NONE")
}
