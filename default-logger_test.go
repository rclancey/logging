package logging

import (
	"bytes"
	"context"
	"log"
	"strings"

	. "gopkg.in/check.v1"
)

type DefaultSuite struct {}
var _ = Suite(&DefaultSuite{})

func (a *DefaultSuite) TestLevel(c *C) {
	SetLevel(DEBUG)
	c.Check(Level(), Equals, DEBUG)
}

func (a *DefaultSuite) TestOutput(c *C) {
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	c.Check(Writer(), Equals, buf)
}

func (a *DefaultSuite) TestPrefix(c *C) {
	SetPrefix("unittest")
	c.Check(Prefix(), Equals, "unittest")
}

func (a *DefaultSuite) TestRawLog(c *C) {
	dbuf := NewBuffer()
	SetOutput(dbuf)
	SetLevel(WARNING)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("defaulttest")
	lbuf := NewBuffer()
	l := NewLogger(lbuf, DEBUG)
	l.SetFlags(log.Ldate | log.Ltime)
	l.SetPrefix("customtest")
	RawLog(nil, DEBUG, 1)
	c.Check(len(dbuf.Bytes()), Equals, 0)
	c.Check(len(lbuf.Bytes()), Equals, 0)
	RawLog(nil, ERROR, "ab", "cd")
	dbuf.Wait()
	c.Check(strings.TrimSpace(string(dbuf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    defaulttest default-logger_test.go:[0-9]+: abcd$")
	c.Check(len(lbuf.Bytes()), Equals, 0)
	ctx := NewContext(context.Background(), l)
	RawLog(ctx, DEBUG, "ab", "cd")
	lbuf.Wait()
	c.Check(strings.TrimSpace(string(lbuf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} DEBUG    customtest abcd$")
}

func (a *DefaultSuite) TestRawLogln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(WARNING)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	RawLogln(nil, ERROR, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestRawLogf(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(WARNING)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	RawLogf(nil, ERROR, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestPrint(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Print(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestPrintln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Println(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestPrintf(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Printf(nil, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestDebug(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Debug(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestDebugln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Debugln(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestDebugf(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Debugf(nil, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestInfo(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Info(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestInfoln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Infoln(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestInfof(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Infof(nil, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestWarn(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Warn(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestWarnln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Warnln(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestWarnf(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Warnf(nil, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestError(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Error(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestErrorln(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Errorln(nil, "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestErrorf(c *C) {
	buf := NewBuffer()
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Errorf(nil, "%s / %s", "ab", "cd")
	buf.Wait()
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest default-logger_test.go:[0-9]+: ab / cd$")
}

func (a *DefaultSuite) TestFatal(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Fatal(nil, "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: abcd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *DefaultSuite) TestFatalln(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Fatalln(nil, "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: ab cd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *DefaultSuite) TestFatalf(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	Fatalf(nil, "%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: ab / cd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *DefaultSuite) TestPanic(c *C) {
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	c.Check(func() { Panic(nil, "ab", "cd") }, Panics, "abcd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: abcd$")
}

func (a *DefaultSuite) TestPanicln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	c.Check(func() { Panicln(nil, "ab", "cd") }, Panics, "ab cd\n")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: ab cd$")
}

func (a *DefaultSuite) TestPanicf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	SetOutput(buf)
	SetLevel(DEBUG)
	SetFlags(log.Ldate | log.Lshortfile)
	SetPrefix("unittest")
	c.Check(func() { Panicf(nil, "%s / %s", "ab", "cd") }, Panics, "ab / cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest default-logger_test.go:[0-9]+: ab / cd$")
}
