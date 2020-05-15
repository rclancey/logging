package logging

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }
type LoggingSuite struct {}
var _ = Suite(&LoggingSuite{})

func (a *LoggingSuite) TestNewLogger(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, INFO)
	c.Check(l.Writer(), Equals, buf)
	c.Check(l.Level(), Equals, INFO)
	c.Check(l.Prefix(), Equals, "")
	xbuf := bytes.NewBuffer([]byte{})
	l.SetOutput(xbuf)
	c.Check(l.Writer(), Equals, xbuf)
	l.SetLevel(ERROR)
	c.Check(l.Level(), Equals, ERROR)
	l.SetFlags(log.LUTC | log.Ldate | log.Llongfile)
	c.Check(l.TimeZone(), Equals, time.UTC)
	c.Check(l.TimeFormat(), Equals, "2006/01/02")
	c.Check(l.SourceFormat().format, Equals, NewSourceFormatter("%{fullpath}:%{linenumber}:").format)
	l.SetPrefix("unittest")
	c.Check(l.Prefix(), Equals, "unittest")
}

func (a *LoggingSuite) TestWiths(c *C) {
	tz, err := time.LoadLocation("America/New_York")
	c.Assert(err, IsNil)
	buf := bytes.NewBuffer([]byte{})
	xbuf := bytes.NewBuffer([]byte{})
	l1 := NewLogger(buf, INFO)
	l2 := l1.WithOutput(xbuf)
	l3 := l2.WithLevel(ERROR)
	l4 := l3.WithLevelColor(ERROR, ColorHotPink, ColorDefault, FontDefault)
	l5 := l4.WithTimeFormat("Mon Jan 02, 2006 03:04 PM")
	l6 := l5.WithTimeZone(tz)
	l7 := l6.WithTimeColor(ColorBlue, ColorDefault, FontDefault)
	l8 := l7.WithSourceFormat("%{package}:%{qualifiedfunction}:")
	l9 := l8.WithSourceColor(ColorRed, ColorCyan, FontItalic)
	l10 := l9.WithPrefix("unittest")
	l11 := l10.WithPrefixColor(ColorDefault, ColorDefault, FontDefault)
	l12 := l11.WithMessageColor(ColorDefault, ColorDefault, FontItalic)
	l13 := l12.WithFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	l14 := l13.WithFlags(log.Ldate)
	l15 := l14.WithFlags(log.Ltime)
	l16 := l15.WithFlags(log.Ltime | log.Lmicroseconds)
	l17 := l16.WithFlags(0)
	c.Check(l1.Writer(), Equals, buf)
	c.Check(l1.Level(), Equals, INFO)
	c.Check(l1.LevelColor(ERROR).foreground, Equals, ColorRed)
	c.Check(l1.TimeFormat(), Equals, "2006/01/02 15:04:05")
	c.Check(l1.TimeZone(), Equals, time.Local)
	c.Check(l1.TimeColor(), IsNil)
	c.Check(l1.SourceFormat().format, Equals, NewSourceFormatter("%{filename}:%{linenumber}:").format)
	c.Check(l1.SourceColor(), IsNil)
	c.Check(l1.Prefix(), Equals, "")
	c.Check(l1.PrefixColor(), IsNil)
	c.Check(l1.MessageColor(), IsNil)

	c.Check(l2.Writer(), Equals, xbuf)
	c.Check(l2.Level(), Equals, INFO)

	c.Check(l3.Writer(), Equals, xbuf)
	c.Check(l3.Level(), Equals, ERROR)

	c.Check(l4.LevelColor(ERROR).foreground, Equals, ColorHotPink)
	c.Check(l5.TimeFormat(), Equals, "Mon Jan 02, 2006 03:04 PM")
	c.Check(l6.TimeZone(), Equals, tz)
	c.Check(l7.TimeColor().foreground, Equals, ColorBlue)
	c.Check(l7.TimeColor().background, Equals, ColorDefault)
	c.Check(l7.TimeColor().font, Equals, FontDefault)
	c.Check(l8.SourceFormat().format, Equals, NewSourceFormatter("%{package}:%{qualifiedfunction}:").format)
	c.Check(l9.SourceColor().foreground, Equals, ColorRed)
	c.Check(l9.SourceColor().background, Equals, ColorCyan)
	c.Check(l9.SourceColor().font, Equals, FontItalic)
	c.Check(l10.Prefix(), Equals, "unittest")
	c.Check(l11.PrefixColor().foreground, Equals, ColorDefault)
	c.Check(l12.MessageColor().font, Equals, FontItalic)
	c.Check(l13.TimeZone(), Equals, time.Local)
	c.Check(l13.TimeFormat(), Equals, "2006/01/02 15:04:05.000000")
	c.Check(l14.TimeFormat(), Equals, "2006/01/02")
	c.Check(l15.TimeFormat(), Equals, "15:04:05")
	c.Check(l16.TimeFormat(), Equals, "15:04:05.000000")
	c.Check(l17.TimeFormat(), Equals, "")
}

func (a *LoggingSuite) TestWrite(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, INFO)
	n, err := l.Write([]byte("when in the course of human events\n"))
	c.Check(n, Equals, 35)
	c.Check(err, IsNil)
	c.Check(string(buf.Bytes()), Equals, "when in the course of human events\n")
}

func (a *LoggingSuite) TestPrint(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Print("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestPrintln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Println("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestPrintf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Printf("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2}          unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestDebug(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Debug("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestDebugln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Debugln("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestDebugf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Debugf("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} DEBUG    unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestInfo(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Info("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestInfoln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Infoln("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestInfof(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Infof("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} INFO     unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestWarn(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Warn("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestWarnln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Warnln("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestWarnf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Warnf("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} WARNING  unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestError(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Error("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestErrorln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Errorln("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestErrorf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Errorf("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} ERROR    unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestFatal(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Fatal("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: abcd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *LoggingSuite) TestFatalln(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Fatalln("ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: ab cd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *LoggingSuite) TestFatalf(c *C) {
	exitStatus := -1
	exiter = func(n int) { exitStatus = n }
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	l.Fatalf("%s / %s", "ab", "cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: ab / cd$")
	c.Check(exitStatus, Equals, 1)
}

func (a *LoggingSuite) TestPanic(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	c.Check(func() { l.Panic("ab", "cd") }, Panics, "abcd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: abcd$")
}

func (a *LoggingSuite) TestPanicln(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	c.Check(func() { l.Panicln("ab", "cd") }, Panics, "ab cd\n")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: ab cd$")
}

func (a *LoggingSuite) TestPanicf(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetFlags(log.Ldate | log.Lshortfile)
	l.SetPrefix("unittest")
	c.Check(func() { l.Panicf("%s / %s", "ab", "cd") }, Panics, "ab / cd")
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, "^[0-9]{4}/[0-9]{2}/[0-9]{2} CRITICAL unittest logging_test.go:[0-9]+: ab / cd$")
}

func (a *LoggingSuite) TestTrace(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, INFO)
	l.SetFlags(log.Ldate)
	l.SetPrefix("trace")
	l.Trace()
	lines := strings.Split(strings.TrimSpace(string(buf.Bytes())), "\n")
	c.Check(len(lines) >= 6, Equals, true)
	c.Check(lines[0], Matches, `[0-9]{4}/[0-9]{2}/[0-9]{2} trace github.com/rclancey/logging.\(\*Logger\).Trace\(\)`)
	c.Check(lines[1], Matches, `           trace     /.*/github.com/rclancey/logging/logging.go:[0-9]+`)
	c.Check(lines[2], Matches, `           trace github.com/rclancey/logging.\(\*LoggingSuite\).TestTrace\(\)`)
	c.Check(lines[3], Matches, `           trace     /.*/github.com/rclancey/logging/logging_test.go:[0-9]+`)
}

func (a *LoggingSuite) TestMakeDefault(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, INFO)
	l.SetFlags(log.Ldate | log.Llongfile)
	l.SetPrefix("yowza")
	l.MakeDefault()
	c.Check(defaultLogger, Equals, l)
	//c.Check(log.Writer(), Equals, l)
	//c.Check(log.Flags(), Equals, log.Ldate | log.Llongfile)
	//c.Check(log.Prefix(), Equals, "yowza")
}

func (a *LoggingSuite) TestClone(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	cl := l.Clone()
	cl.SetLevel(INFO)
	c.Check(l.Level(), Equals, DEBUG)
	c.Check(cl.Level(), Equals, INFO)
}

func (a *LoggingSuite) TestFormatting(c *C) {
	buf := bytes.NewBuffer([]byte{})
	l := NewLogger(buf, DEBUG)
	l.SetSourceFormat("")
	before := time.Now()
	l.Warnln("abcd")
	after := time.Now()
	log.Println("warn took", after.Sub(before))
	os.Stderr.Write(buf.Bytes())
	l = l.WithColor()
	buf.Reset()
	l.SetLevelColor(DEBUG, ColorBlue, ColorBlack, FontItalic)
	l.SetSourceFormat("%{package}/%{filename}:%{linenumber:05x}:%{foo:x}")
	l.SetTimeFormat("2006-01-02T15:04:05-07:00")
	tz, err := time.LoadLocation("America/Los_Angeles")
	c.Assert(err, IsNil)
	l.SetTimeZone(tz)
	l.SetTimeColor(ColorHotPink, ColorDefault, FontDefault)
	l.SetPrefix("unittest")
	before = time.Now()
	l.Debugln("abcd")
	after = time.Now()
	log.Println("debug took", after.Sub(before))
	os.Stderr.Write(buf.Bytes())
	c.Check(strings.TrimSpace(string(buf.Bytes())), Matches, `^.\[38;5;199;49m[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}-0[78]:00.\[0m .\[34;40;3mDEBUG   .\[0m .\[34;40;3munittest.\[0m .\[34;40;3mgithub\.com/rclancey/logging/logging_test\.go:[0-9a-f]{5}:%{foo:x}.\[0m .\[34;40;3mabcd.\[0m$`)
}
