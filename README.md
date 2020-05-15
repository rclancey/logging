# Logging package for Go

[![Coverage Status](https://coveralls.io/repos/github/rclancey/logging/badge.svg?branch=master)](https://coveralls.io/github/rclancey/logging)
[![GoDoc](http://godoc.org/github.com/rclancey/logging?status.svg)](http://godoc.org/github.com/rclancey/logging)

Thus package provides a more complex logger than the default provided by the standard library `log` package

## Installation

Use the `go` command:

	$ go get github.com/rclancey/logging

## Example

```go
errlog := logging.NewLogger(os.Stdout, logging.DEBUG)
errlog.Colorize()
errlog.SetLevelColor(logging.INFO, logging.ColorCyan, logging.ColorDefault, logging.FontDefault)
errlog.SetLevelColor(logging.LOG, logging.ColorMagenta, logging.ColorDefault, logging.FontDefault)
errlog.SetLevelColor(logging.NONE, logging.ColorHotPink, logging.ColorDefault, logging.FontDefault)
errlog.SetTimeFormat("2006-01-02 15:04:05.000")
errlog.SetTimeColor(logging.ColorDefault, logging.ColorDefault, logging.FontItalic | logging.FontLight)
errlog.SetSourceFormat("%{basepath}:%{linenumber}:")
errlog.SetSourceColor(logging.ColorGreen, logging.ColorDefault, logging.FontDefault)
errlog.SetPrefixColor(logging.ColorOrange, logging.ColorDefault, logging.FontDefault)
errlog.SetMessageColor(logging.ColorDefault, logging.ColorDefault, logging.FontDefault)
errlog.MakeDefault()
errlog.Infoln("Server starting...")
```

## Documentation

[Documentation](http://godoc.org/github.com/rclancey/logging) is hosted at GoDoc project.

## Copyright

Copyright (C) 2019-2020 by Ryan Clancey

Package released under MIT License.
See [LICENSE](https://github.com/rclancey/logging/blob/master/LICENSE) for details.
