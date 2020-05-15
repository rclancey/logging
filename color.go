package logging

import (
	"strings"

	"github.com/pkg/errors"
)

// https://stackoverflow.com/questions/4842424/list-of-ansi-color-escape-sequences

type ColorCode string
type FontCode int

const (
	ColorDefault   = ColorCode("default")
	ColorBlack     = ColorCode("black")
	ColorRed       = ColorCode("red")
	ColorGreen     = ColorCode("green")
	ColorYellow    = ColorCode("yellow")
	ColorBlue      = ColorCode("blue")
	ColorMagenta   = ColorCode("magenta")
	ColorCyan      = ColorCode("cyan")
	ColorWhite     = ColorCode("white")
	ColorHotPink   = ColorCode("hot pink")
	ColorOrange    = ColorCode("orange")
	ColorPurple    = ColorCode("purple")
	ColorTurquoise = ColorCode("turquoise")
	ColorLightGray = ColorCode("light gray")
	ColorDarkGray  = ColorCode("dark gray")

	FontDefault   = FontCode(0)
	FontBold      = FontCode(2)
	FontLight     = FontCode(4)
	FontItalic    = FontCode(8)
	FontUnderline = FontCode(16)
	FontBlink     = FontCode(32)
	FontReverse   = FontCode(128)
)

var colorEscapes = map[ColorCode]string{
	ColorDefault:   "9",
	ColorBlack:     "0",
	ColorRed:       "1",
	ColorGreen:     "2",
	ColorYellow:    "3",
	ColorBlue:      "4",
	ColorMagenta:   "5",
	ColorCyan:      "6",
	ColorWhite:     "7",
	ColorHotPink:   "8;5;199",
	ColorOrange:    "8;5;208",
	ColorPurple:    "8;5;91",
	ColorTurquoise: "8;5;80",
	ColorLightGray: "8;5;250",
	ColorDarkGray:  "8;5;240",
}

var fontEscapes = map[FontCode]string{
	FontDefault:   "0",
	FontBold:      "1",
	FontLight:     "2",
	FontItalic:    "3",
	FontUnderline: "4",
	FontBlink:     "5",
	FontReverse:   "7",
}

type Colorizer struct {
	foreground ColorCode
	background ColorCode
	font FontCode
	escape string
}

func NewColorizer(foreground, background ColorCode, font FontCode) (*Colorizer, error) {
	c := &Colorizer{
		foreground: ColorDefault,
		background: ColorDefault,
		font: FontDefault,
		escape: "",
	}
	err := c.SetForeground(foreground)
	if err != nil {
		return nil, err
	}
	err = c.SetBackground(background)
	if err != nil {
		return nil, err
	}
	err = c.SetFont(font)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Colorizer) SetForeground(fg ColorCode) error {
	return c.update(fg, c.background, c.font)
}

func (c *Colorizer) GetForeground() ColorCode {
	return c.foreground
}

func (c *Colorizer) SetBackground(bg ColorCode) error {
	return c.update(c.foreground, bg, c.font)
}

func (c *Colorizer) GetBackground() ColorCode {
	return c.background
}

func (c *Colorizer) SetFont(font FontCode) error {
	return c.update(c.foreground, c.background, font)
}

func (c *Colorizer) GetFont() FontCode {
	return c.font
}

func (c *Colorizer) update(fg, bg ColorCode, font FontCode) error {
	if fg == ColorDefault && bg == ColorDefault && font == FontDefault {
		c.foreground = fg
		c.background = bg
		c.font = font
		c.escape = ""
		return nil
	}
	escapeCodes := []string{}
	esc, ok := colorEscapes[fg]
	if ok {
		escapeCodes = append(escapeCodes, "3" + esc)
	} else {
		return errors.Errorf("unknown foreground color '%s'", string(fg))
	}
	esc, ok = colorEscapes[bg]
	if ok {
		escapeCodes = append(escapeCodes, "4" + esc)
	} else {
		return errors.Errorf("unknown background color '%s'", string(bg))
	}
	for i := 1; i <= 1024; i *= 2 {
		if int(font) & i != 0 {
			esc, ok = fontEscapes[FontCode(i)]
			if ok {
				escapeCodes = append(escapeCodes, esc)
			}
		}
	}
	c.foreground = fg
	c.background = bg
	c.font = font
	c.escape = "\033[" + strings.Join(escapeCodes, ";") + "m"
	return nil
}

func (c *Colorizer) Colorize(message string) string {
	if c.escape == "" {
		return message
	}
	return c.escape + message + "\033[0m"
}

