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

func RawLog(ctx context.Context, skip int, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLog(skip + 1, level, args...)
}

func RawLogln(ctx context.Context, skip int, level LogLevel, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogln(skip + 1, level, args...)
}

func RawLogf(ctx context.Context, skip int, level LogLevel, format string, args ...interface{}) {
	l := FromContext(ctx)
	l.RawLogf(skip + 1, level, format, args...)
}

func Print(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, NONE, args...)
}

func Println(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, NONE, args...)
}

func Printf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, NONE, format, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, DEBUG, args...)
}

func Debugln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, DEBUG, args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, DEBUG, format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, INFO, args...)
}

func Infoln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, INFO, args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, INFO, format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, WARNING, args...)
}

func Warnln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, WARNING, args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, WARNING, format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, ERROR, args...)
}

func Errorln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, ERROR, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, ERROR, format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, CRITICAL, args...)
	exiter(1)
}

func Fatalln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, CRITICAL, args...)
	exiter(1)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, CRITICAL, format, args...)
	exiter(1)
}

func Panic(ctx context.Context, args ...interface{}) {
	RawLog(ctx, 1, CRITICAL, args...)
	panic(fmt.Sprint(args...))
}

func Panicln(ctx context.Context, args ...interface{}) {
	RawLogln(ctx, 1, CRITICAL, args...)
	panic(fmt.Sprintln(args...))
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	RawLogf(ctx, 1, CRITICAL, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Trace(ctx context.Context) {
	l := FromContext(ctx)
	l.RawTrace(1, l.Prefix())
}
