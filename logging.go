package logging

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	//"github.com/pkg/errors"
)

// required for testing fatal
var exiter = func(n int) { os.Exit(n) }

type Logger struct {
	w io.Writer
	colorize bool
	level LogLevel
	levelColor map[LogLevel]*Colorizer
	timeFormat string
	timeZone *time.Location
	timeColor *Colorizer
	sourceFormat *SourceFormatter
	sourceColor *Colorizer
	prefix string
	prefixColor *Colorizer
	messageColor *Colorizer
}

func NewLogger(w io.Writer, level LogLevel) *Logger {
	l := &Logger{
		w: w,
		colorize: false,
		level: level,
		levelColor: map[LogLevel]*Colorizer{},
		timeFormat: "2006/01/02 15:04:05",
		timeZone: time.Local,
		timeColor: nil,
		sourceFormat: nil,
		sourceColor: nil,
		prefix: "",
		prefixColor: nil,
		messageColor: nil,
	}
	l.SetLevelColor(DEBUG,    ColorLightGray, ColorDefault, FontDefault)
	l.SetLevelColor(INFO,     ColorBlue,      ColorDefault, FontDefault)
	l.SetLevelColor(WARNING,  ColorYellow,    ColorDefault, FontDefault)
	l.SetLevelColor(ERROR,    ColorRed,       ColorDefault, FontDefault)
	l.SetLevelColor(CRITICAL, ColorRed,       ColorDefault, FontBold | FontBlink)
	l.SetSourceFormat("%{filename}:%{linenumber}:")
	return l
}

func (l *Logger) Clone() *Logger {
	c := *l
	return &c
}

func (l *Logger) WithOutput(w io.Writer) *Logger {
	l = l.Clone()
	l.SetOutput(w)
	return l
}

func (l *Logger) SetOutput(w io.Writer) {
	l.w = w
}

func (l *Logger) Writer() io.Writer {
	return l.w
}

func (l *Logger) WithColor() *Logger {
	l.Clone()
	l.colorize = true
	return l
}

func (l *Logger) WithoutColor() *Logger {
	l.Clone()
	l.colorize = false
	return l
}

func (l *Logger) Colorize() {
	l.colorize = true
}

func (l *Logger) WithLevel(level LogLevel) *Logger {
	l = l.Clone()
	l.SetLevel(level)
	return l
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Level() LogLevel {
	return l.level
}

func (l *Logger) WithLevelColor(level LogLevel, fg, bg ColorCode, font FontCode) *Logger {
	lc := map[LogLevel]*Colorizer{}
	for k, v := range l.levelColor {
		lc[k] = v
	}
	l = l.Clone()
	l.levelColor = lc
	l.SetLevelColor(level, fg, bg, font)
	return l
}

func (l *Logger) SetLevelColor(level LogLevel, fg, bg ColorCode, font FontCode) {
	c, err := NewColorizer(fg, bg, font)
	if err == nil {
		l.levelColor[level] = c
	}
}

func (l *Logger) LevelColor(level LogLevel) *Colorizer {
	return l.levelColor[level]
}

func (l *Logger) WithTimeFormat(timeFormat string) *Logger {
	l = l.Clone()
	l.SetTimeFormat(timeFormat)
	return l
}

func (l *Logger) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

func (l *Logger) TimeFormat() string {
	return l.timeFormat
}

func (l *Logger) WithTimeZone(timeZone *time.Location) *Logger {
	l = l.Clone()
	l.SetTimeZone(timeZone)
	return l
}

func (l *Logger) SetTimeZone(timeZone *time.Location) {
	l.timeZone = timeZone
}

func (l *Logger) TimeZone() *time.Location {
	return l.timeZone
}

func (l *Logger) WithTimeColor(fg, bg ColorCode, font FontCode) *Logger {
	l = l.Clone()
	l.SetTimeColor(fg, bg, font)
	return l
}

func (l *Logger) SetTimeColor(fg, bg ColorCode, font FontCode) {
	c, err := NewColorizer(fg, bg, font)
	if err == nil {
		l.timeColor = c
	}
}

func (l *Logger) TimeColor() *Colorizer {
	return l.timeColor
}

func (l *Logger) WithSourceFormat(format string) *Logger {
	l = l.Clone()
	l.SetSourceFormat(format)
	return l
}

func (l *Logger) SetSourceFormat(format string) {
	if format == "" {
		l.sourceFormat = nil
	} else {
		l.sourceFormat = NewSourceFormatter(format)
	}
}

func (l *Logger) SourceFormat() *SourceFormatter {
	return l.sourceFormat
}

func (l *Logger) WithSourceColor(fg, bg ColorCode, font FontCode) *Logger {
	l = l.Clone()
	l.SetSourceColor(fg, bg, font)
	return l
}

func (l *Logger) SetSourceColor(fg, bg ColorCode, font FontCode) {
	c, err := NewColorizer(fg, bg, font)
	if err == nil {
		l.sourceColor = c
	}
}

func (l *Logger) SourceColor() *Colorizer {
	return l.sourceColor
}

func (l *Logger) WithPrefix(prefix string) *Logger {
	l = l.Clone()
	l.SetPrefix(prefix)
	return l
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) WithPrefixColor(fg, bg ColorCode, font FontCode) *Logger {
	l = l.Clone()
	l.SetPrefixColor(fg, bg, font)
	return l
}

func (l *Logger) SetPrefixColor(fg, bg ColorCode, font FontCode) {
	c, err := NewColorizer(fg, bg, font)
	if err == nil {
		l.prefixColor = c
	}
}

func (l *Logger) PrefixColor() *Colorizer {
	return l.prefixColor
}

func (l *Logger) WithMessageColor(fg, bg ColorCode, font FontCode) *Logger {
	l = l.Clone()
	l.SetMessageColor(fg, bg, font)
	return l
}

func (l *Logger) SetMessageColor(fg, bg ColorCode, font FontCode) {
	c, err := NewColorizer(fg, bg, font)
	if err == nil {
		l.messageColor = c
	}
}

func (l *Logger) MessageColor() *Colorizer {
	return l.messageColor
}

func (l *Logger) WithFlags(flags int) *Logger {
	l = l.Clone()
	l.SetFlags(flags)
	return l
}

func (l *Logger) SetFlags(flags int) {
	if flags & log.Ldate != 0 {
		if flags & log.Ltime != 0 {
			if flags & log.Lmicroseconds != 0 {
				l.SetTimeFormat("2006/01/02 15:04:05.000000")
			} else {
				l.SetTimeFormat("2006/01/02 15:04:05")
			}
		} else {
			l.SetTimeFormat("2006/01/02")
		}
	} else if flags & log.Ltime != 0 {
		if flags & log.Lmicroseconds != 0 {
			l.SetTimeFormat("15:04:05.000000")
		} else {
			l.SetTimeFormat("15:04:05")
		}
	} else {
		l.SetTimeFormat("")
	}
	if flags & log.LUTC != 0 {
		l.SetTimeZone(time.UTC)
	} else {
		l.SetTimeZone(time.Local)
	}
	if flags & log.Lshortfile != 0 {
		l.SetSourceFormat("%{filename}:%{linenumber}:")
	} else if flags & log.Llongfile != 0 {
		l.SetSourceFormat("%{fullpath}:%{linenumber}:")
	} else {
		l.SetSourceFormat("")
	}
}

func (l *Logger) getColorizer(defaultColorizer, colorizer *Colorizer) *Colorizer {
	if !l.colorize {
		return nil
	}
	if colorizer != nil {
		return colorizer
	}
	return defaultColorizer
}

func (l *Logger) RawWrite(ctx context.Context, level LogLevel, message string, trace ...string) (int, error) {
	if level > l.level {
		return 0, nil
	}
	skip := getDepth(ctx)
	sr := NewSourceRecord(skip + 1)
	return l.RawWriteWithSource(ctx, level, sr, message, trace...)
}

func (l *Logger) RawWriteAsync(ctx context.Context, level LogLevel, message string, trace ...string) {
	if level > l.level {
		return
	}
	skip := getDepth(ctx)
	sr := NewSourceRecord(skip + 1)
	go func() {
		l.RawWriteWithSource(ctx, level, sr, message, trace...)
	}()
}

func (l *Logger) RawWriteWithSource(ctx context.Context, level LogLevel, sr *SourceRecord, message string, trace ...string) (int, error) {
	if level > l.level {
		return 0, nil
	}
	dc := l.getColorizer(l.levelColor[level], nil)
	line := ""
	if l.timeFormat != "" {
		t := time.Now()
		if l.timeZone != nil {
			t = t.In(l.timeZone)
		}
		c := l.getColorizer(dc, l.timeColor)
		if c != nil {
			line += c.Colorize(t.Format(l.timeFormat))
		} else {
			line += t.Format(l.timeFormat)
		}
		line += " "
	}
	if dc != nil {
		line += dc.Colorize(level.PaddedString(8))
	} else {
		line += level.PaddedString(8)
	}
	line += " "
	if l.prefix != "" {
		c := l.getColorizer(dc, l.prefixColor)
		if c != nil {
			line += c.Colorize(l.prefix)
		} else {
			line += l.prefix
		}
		line += " "
	}
	if len(trace) > 0 {
		line += dc.Colorize(trace[0]) + " "
	}
	if l.sourceFormat != nil {
		c := l.getColorizer(dc, l.sourceColor)
		if c != nil {
			line += c.Colorize(l.sourceFormat.FormatRecord(sr))
		} else {
			line += l.sourceFormat.FormatRecord(sr)
		}
		line += " "
	}
	c := l.getColorizer(dc, l.messageColor)
	if c != nil {
		line += c.Colorize(strings.TrimSpace(message))
	} else {
		line += strings.TrimSpace(message)
	}
	line += "\n"
	return l.w.Write([]byte(line))
}

func (l *Logger) RawStackTrace(ctx context.Context, prefix string) {
	skip := getDepth(ctx)
	padding := ""
	if l.timeFormat != "" {
		t := time.Now()
		if l.timeZone != nil {
			t = t.In(l.timeZone)
		}
		padding += t.Format(l.timeFormat) + " "
	}
	first := true
	for {
		sr := NewSourceRecord(skip)
		if sr == nil {
			return
		}
		skip += 1
		line := fmt.Sprintf("%s%s %s.%s()\n", padding, prefix, sr.Package, sr.QualifiedFunction)
		l.w.Write([]byte(line))
		if first {
			padding = strings.Repeat(" ", len(padding))
		}
		line = fmt.Sprintf("%s%s     %s:%d\n", padding, prefix, sr.FullPath, sr.LineNumber)
		l.w.Write([]byte(line))
	}
}

func (l *Logger) RawLogSync(ctx context.Context, level LogLevel, args ...interface{}) {
	if l.level >= level {
		l.RawWrite(deepen(ctx), level, fmt.Sprint(args...))
	}
}

func (l *Logger) RawLoglnSync(ctx context.Context, level LogLevel, args ...interface{}) {
	if l.level >= level {
		l.RawWrite(deepen(ctx), level, fmt.Sprintln(args...))
	}
}

func (l *Logger) RawLogfSync(ctx context.Context, level LogLevel, format string, args ...interface{}) {
	if l.level >= level {
		l.RawWrite(deepen(ctx), level, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) RawLog(ctx context.Context, level LogLevel, args ...interface{}) {
	if l.level >= level {
		l.RawWriteAsync(deepen(ctx), level, fmt.Sprint(args...))
	}
}

func (l *Logger) RawLogln(ctx context.Context, level LogLevel, args ...interface{}) {
	if l.level >= level {
		l.RawWriteAsync(deepen(ctx), level, fmt.Sprintln(args...))
	}
}

func (l *Logger) RawLogf(ctx context.Context, level LogLevel, format string, args ...interface{}) {
	if l.level >= level {
		l.RawWriteAsync(deepen(ctx), level, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) Write(data []byte) (int, error) {
	return l.RawWrite(withDepth(nil, 3), LOG, string(data))
	//return l.w.Write(data)
}

func (l *Logger) Print(args ...interface{}) {
	l.RawLog(withDepth(nil, 1), NONE, args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.RawLogln(withDepth(nil, 1), NONE, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.RawLogf(withDepth(nil, 1), NONE, format, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.RawLog(withDepth(nil, 1), DEBUG, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.RawLogln(withDepth(nil, 1), DEBUG, args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.RawLogf(withDepth(nil, 1), DEBUG, format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.RawLog(withDepth(nil, 1), INFO, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.RawLogln(withDepth(nil, 1), INFO, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.RawLogf(withDepth(nil, 1), INFO, format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.RawLog(withDepth(nil, 1), WARNING, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.RawLogln(withDepth(nil, 1), WARNING, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.RawLogf(withDepth(nil, 1), WARNING, format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.RawLog(withDepth(nil, 1), ERROR, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.RawLogln(withDepth(nil, 1), ERROR, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.RawLogf(withDepth(nil, 1), ERROR, format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.RawLogSync(withDepth(nil, 1), CRITICAL, args...)
	exiter(1)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.RawLoglnSync(withDepth(nil, 1), CRITICAL, args...)
	exiter(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.RawLogfSync(withDepth(nil, 1), CRITICAL, format, args...)
	exiter(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.RawLogSync(withDepth(nil, 1), CRITICAL, args...)
	panic(fmt.Sprint(args...))
}

func (l *Logger) Panicln(args ...interface{}) {
	l.RawLoglnSync(withDepth(nil, 1), CRITICAL, args...)
	panic(fmt.Sprintln(args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.RawLogfSync(withDepth(nil, 1), CRITICAL, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) StackTrace() {
	l.RawStackTrace(withDepth(nil, 1), l.prefix)
}

func (l *Logger) MakeDefault() {
	defaultLogger = l
	log.SetFlags(0)
	log.SetPrefix("")
	log.SetOutput(l)
}
