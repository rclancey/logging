package logging

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

var defaultLogger = NewLogger(os.Stderr, WARNING)

func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

func Writer() io.Writer {
	return defaultLogger.Writer()
}

func SetLevel(level LogLevel) {
	defaultLogger.SetLevel(level)
}

func Level() LogLevel {
	return defaultLogger.Level()
}

func SetLevelColor(level LogLevel, fg, bg ColorCode, font FontCode) {
	defaultLogger.SetLevelColor(level, fg, bg, font)
}

func LevelColor(level LogLevel) *Colorizer {
	return defaultLogger.LevelColor(level)
}

func SetTimeFormat(timeFormat string) {
	defaultLogger.SetTimeFormat(timeFormat)
}

func TimeFormat() string {
	return defaultLogger.TimeFormat()
}

func SetTimeZone(timeZone *time.Location) {
	defaultLogger.SetTimeZone(timeZone)
}

func TimeZone() *time.Location {
	return defaultLogger.TimeZone()
}

func SetTimeColor(fg, bg ColorCode, font FontCode) {
	defaultLogger.SetTimeColor(fg, bg, font)
}

func TimeColor() *Colorizer {
	return defaultLogger.TimeColor()
}

func SetSourceFormat(format string) {
	defaultLogger.SetSourceFormat(format)
}

func SourceFormat() *SourceFormatter {
	return defaultLogger.SourceFormat()
}

func SetSourceColor(fg, bg ColorCode, font FontCode) {
	defaultLogger.SetSourceColor(fg, bg, font)
}

func SourceColor() *Colorizer {
	return defaultLogger.SourceColor()
}

func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}

func Prefix() string {
	return defaultLogger.Prefix()
}

func SetMessageColor(fg, bg ColorCode, font FontCode) {
	defaultLogger.SetMessageColor(fg, bg, font)
}

func MessageColor() *Colorizer {
	return defaultLogger.MessageColor()
}

func SetFlags(flags int) {
	defaultLogger.SetFlags(flags)
}

func RawLogSync(ctx context.Context, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogSync(deepen(ctx), level, args...)
}

func RawLoglnSync(ctx context.Context, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLoglnSync(deepen(ctx), level, args...)
}

func RawLogfSync(ctx context.Context, level LogLevel, format string, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogfSync(deepen(ctx), level, format, args...)
}

func RawLog(ctx context.Context, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLog(deepen(ctx), level, args...)
}

func RawLogln(ctx context.Context, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogln(deepen(ctx), level, args...)
}

func RawLogf(ctx context.Context, level LogLevel, format string, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogf(deepen(ctx), level, format, args...)
}

func Print(ctx context.Context, args ...interface{}) {
	RawLog(deepen(ctx), NONE, args...)
}

func Println(ctx context.Context, args ...interface{}) {
	RawLogln(deepen(ctx), NONE, args...)
}

func Printf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(deepen(ctx), NONE, format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	RawLog(deepen(ctx), DEBUG, args...)
}

func Debugln(ctx context.Context, args ...interface{}) {
	RawLogln(deepen(ctx), DEBUG, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(deepen(ctx), DEBUG, format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	RawLog(deepen(ctx), INFO, args...)
}

func Infoln(ctx context.Context, args ...interface{}) {
	RawLogln(deepen(ctx), INFO, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	RawLogf(deepen(ctx), INFO, format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	RawLog(deepen(ctx), WARNING, args...)
}

func Warnln(ctx context.Context, args ...interface{}) {
	RawLogln(deepen(ctx), WARNING, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(deepen(ctx), WARNING, format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	RawLog(deepen(ctx), ERROR, args...)
}

func Errorln(ctx context.Context, args ...interface{}) {
	RawLogln(deepen(ctx), ERROR, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(deepen(ctx), ERROR, format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	RawLogSync(deepen(ctx), CRITICAL, args...)
	exiter(1)
}

func Fatalln(ctx context.Context, args ...interface{}) {
	RawLoglnSync(deepen(ctx), CRITICAL, args...)
	exiter(1)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	RawLogfSync(deepen(ctx), CRITICAL, format, args...)
	exiter(1)
}

func Panic(ctx context.Context, args ...interface{}) {
	RawLogSync(deepen(ctx), CRITICAL, args...)
	panic(fmt.Sprint(args...))
}

func Panicln(ctx context.Context, args ...interface{}) {
	RawLoglnSync(deepen(ctx), CRITICAL, args...)
	panic(fmt.Sprintln(args...))
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	RawLogfSync(deepen(ctx), CRITICAL, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func StackTrace(ctx context.Context) {
	l := FromContext(ctx)
	l.RawStackTrace(deepen(ctx), l.Prefix())
}

func Trace(ctx context.Context, fnc TraceFunc, args ...interface{}) error {
	l := FromContext(ctx)
	return l.Trace(deepen(ctx), fnc, args...)
}

func Traceln(ctx context.Context, fnc TraceFunc, args ...interface{}) error {
	l := FromContext(ctx)
	return l.Traceln(deepen(ctx), fnc, args...)
}

func Tracef(ctx context.Context, fnc TraceFunc, format string, args ...interface{}) error {
	l := FromContext(ctx)
	return l.Tracef(deepen(ctx), fnc, format, args...)
}
