package logging

import (
	. "gopkg.in/check.v1"
)

type ColorSuite struct {}
var _ = Suite(&ColorSuite{})

func (a *ColorSuite) TestNewColorizer(c *C) {
	cz, err := NewColorizer(ColorHotPink, ColorTurquoise, FontBold | FontItalic)
	c.Check(err, IsNil)
	c.Check(cz, NotNil)
	c.Check(cz.escape, Equals, "\033[38;5;199;48;5;80;1;3m")
	c.Check(cz.Colorize("abcd"), Equals, "\033[38;5;199;48;5;80;1;3mabcd\033[0m")
	cz, err = NewColorizer(ColorCode("fuschia"), ColorDefault, FontDefault)
	c.Check(err, ErrorMatches, `unknown foreground color 'fuschia'`)
	cz, err = NewColorizer(ColorDefault, ColorCode("rainbow"), FontDefault)
	c.Check(err, ErrorMatches, `unknown background color 'rainbow'`)
	cz, err = NewColorizer(ColorDefault, ColorDefault, FontDefault)
	c.Check(cz.escape, Equals, "")
	c.Check(cz.Colorize("abcd"), Equals, "abcd")
}
