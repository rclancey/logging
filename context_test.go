package logging

import (
	"bytes"
	"context"

	. "gopkg.in/check.v1"
)

type ContextSuite struct {}
var _ = Suite(&ContextSuite{})

func (a *ContextSuite) TestContext(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, INFO)
	ctx := NewContext(context.Background(), l)
	cl := FromContext(ctx)
	c.Check(cl, Equals, l)
	cl = FromContext(nil)
	c.Check(cl, NotNil)
	c.Check(cl, FitsTypeOf, l)
	c.Check(cl, Not(Equals), l)
	cl = FromContext(context.Background())
	c.Check(cl, NotNil)
	c.Check(cl, FitsTypeOf, l)
	c.Check(cl, Not(Equals), l)
}
