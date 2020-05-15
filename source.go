package logging

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type SourceRecord struct {
	PC uintptr
	FullPath string
	FileName string
	BasePath string
	LineNumber int
	Package string
	QualifiedFunction string
	Receiver string
	Function string
}

func NewSourceRecord(skip int) *SourceRecord {
	pc, fn, ln, ok := runtime.Caller(skip+1)
	if !ok {
		return nil
	}
	sr := &SourceRecord{
		PC: pc,
		FullPath: fn,
		FileName: filepath.Base(fn),
		LineNumber: ln,
	}
	fnc := runtime.FuncForPC(pc)
	name := fnc.Name()
	pkgpath := strings.Split(name, "/")
	fname := strings.Split(pkgpath[len(pkgpath) - 1], ".")
	pkgpath[len(pkgpath) - 1] = fname[0]
	sr.Package = strings.Join(pkgpath, "/")
	sr.BasePath = filepath.Join(filepath.FromSlash(sr.Package), sr.FileName)
	sr.QualifiedFunction = strings.Join(fname[1:], ".")
	sr.Function = fname[len(fname) - 1]
	if len(fname) > 2 && strings.HasPrefix(fname[1], "(") && strings.HasSuffix(fname[1], ")") {
		sr.Receiver = fname[1][1:len(fname[1])-1]
	}
	return sr
}

type SourceFormatter struct {
	format string
}

var fmtre = regexp.MustCompile(`%{([a-z]+)(?::(.*?))?}`)

func NewSourceFormatter(layout string) *SourceFormatter {
	idx := map[string]string{
		"pc": "[1]",
		"fullpath": "[2]",
		"filename": "[3]",
		"basepath": "[4]",
		"linenumber": "[5]",
		"package": "[6]",
		"receiver": "[7]",
		"function": "[8]",
	}
	verbs := map[string]string{
		"linenumber": "d",
	}
	format := ""
	prev := 0
	for _, m := range fmtre.FindAllStringSubmatchIndex(layout, -1) {
		start, end := m[0], m[1]
		if start > prev {
			format += layout[prev:start]
		}
		prev = end
		name := layout[m[2]:m[3]]
		pos, ok := idx[name]
		if !ok {
			format += "%%{" + name
			if m[4] != -1 {
				format += ":" + layout[m[4]:m[5]]
			}
			format += "}"
			continue
		}
		format += "%"
		if m[4] != -1 {
			format += layout[m[4]:m[5]-1]
			format += pos
			format += layout[m[5]-1:m[5]]
		} else {
			format += pos
			v, ok := verbs[name]
			if ok {
				format += v
			} else {
				format += "s"
			}
		}
	}
	format += layout[prev:]
	return &SourceFormatter{format: format}
}

func (sf *SourceFormatter) Format(skip int) string {
	sr := NewSourceRecord(skip + 1)
	if sr == nil {
		return ""
	}
	return fmt.Sprintf(sf.format, sr.PC, sr.FullPath, sr.FileName, sr.BasePath, sr.LineNumber, sr.Package, sr.Receiver, sr.Function)
}
