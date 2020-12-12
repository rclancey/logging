package logging

import (
	"context"
	"encoding/base32"
	"fmt"
	"math/rand"
	//"runtime"
	//"strconv"
	//"strings"
	"time"
)

const (
	traceIdKey = ctxKey("traceId")
	DefaultTraceID = "xxxxxxxxxxxxx"
)

type TraceFunc func (ctx context.Context) error

var b32enc = base32.StdEncoding.WithPadding(base32.NoPadding)

func genId() string {
	idBytes := make([]byte, 8)
	n, err := rand.Read(idBytes)
	if err != nil || n < len(idBytes) {
		return DefaultTraceID
	}
	return b32enc.EncodeToString(idBytes)
}

func withTraceId(ctx context.Context, id string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, traceIdKey, id)
}

func getTraceId(ctx context.Context) string {
	if ctx == nil {
		return DefaultTraceID
	}
	val := ctx.Value(traceIdKey)
	if val != nil {
		id, isa := val.(string)
		if isa {
			return id
		}
	}
	return DefaultTraceID
}

func (l *Logger) RawTrace(ctx context.Context, fnc TraceFunc, msg string) error {
	parentId := getTraceId(ctx)
	childId := genId()
	childCtx := withDepth(withTraceId(ctx, childId), 0)
	start := time.Now()
	err := fnc(childCtx)
	end := time.Now()
	dur := end.Sub(start).Seconds()
	trace := fmt.Sprintf("%s %s %09.6fs", parentId, childId, dur)
	l.RawWrite(deepen(ctx), TRACE, msg, trace)
	return err
}

func (l *Logger) Trace(ctx context.Context, fnc TraceFunc, args ...interface{}) error {
	msg := fmt.Sprint(args...)
	return l.RawTrace(deepen(ctx), fnc, msg)
}

func (l *Logger) Traceln(ctx context.Context, fnc TraceFunc, args ...interface{}) error {
	msg := fmt.Sprintln(args...)
	return l.RawTrace(deepen(ctx), fnc, msg)
}

func (l *Logger) Tracef(ctx context.Context, fnc TraceFunc, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return l.RawTrace(deepen(ctx), fnc, msg)
}

/*
func Trace(ctx context.Context, name string, f func(ctx context.Context) error, skips... int) error {
	maxDepth := getMaxDepth(ctx)
	depth := getDepth(ctx)
	if maxDepth >= 0 && depth >= maxDepth {
		return f(ctx)
	}
	logger := getLogger(ctx)
	if logger == nil {
		return f(ctx)
	}
	var xname string
	skip := 1
	if len(skips) > 0 {
		skip += skips[0]
	}
	pc, _, ln, ok := runtime.Caller(skip)
	if ok {
		fnc := runtime.FuncForPC(pc)
		pkgpath := strings.Split(fnc.Name(), "/")
		fname := strings.Split(pkgpath[len(pkgpath) - 1], ".")
		xname = strings.Join(fname[1:], ".") + ":" + strconv.Itoa(ln)
	} else {
		xname = "??"
	}
	if name != "" {
		xname += ":" + name
	}
	parentId := getParentId(ctx)
	id := genId()
	begin := getBegin(ctx)
	myCtx := withDepth(withParentId(ctx, id), depth + 1)
	start := time.Now()
	if begin {
		logger.Infoln("BEGIN   ", id, parentId, xname)
	}
	err := f(myCtx)
	end := time.Now()
	if err != nil {
		logger.Infof("ERROR    %s %s %09.6fs %s %s", id, parentId, end.Sub(start).Seconds(), xname, err)
	} else {
		logger.Infof("COMPLETE %s %s %09.6fs %s", id, parentId, end.Sub(start).Seconds(), xname)
	}
	return err
}
*/
