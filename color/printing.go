// the functions in this file allow to directly print ANSI coloured text to the
// standard output
//
// Colours are passed to the functions by strings that define a colour code.
//
// Colour code syntax:
// 	r ... red
//	g ... green
// 	y ... yellow
//	b ... blue
//	x ... black
//	m ... magenta
// 	c ... cyan
//	w ... white
//	d ... default
//	+ ... bold
//	* ... italic
//	~ ... reverse
// 	_ ... underlined
//	# ... blink
//	? ... concealed
//
//	background coloring is affected by using upper case letters
//
// examples:
//	+r ... bold red
//	rY ... red text on yellow background
package color

import (
	"fmt"
	"io"
	"strings"
)

// --- EXPORTED ---

func Printf(colorCode string, format string, a ...interface{}) {
	colored := colourText(colorCode, format)
	fmt.Printf(colored, a...)
}

func Print(colorCode string, a ...interface{}) {
	colored := fmt.Sprint(a...)
	colored = colourText(colorCode, colored)
	fmt.Print(colored)
}

func Println(colorCode string, a ...interface{}) {
	colored := fmt.Sprint(a...)
	fmt.Println(colored)
}

func Sprintf(colorCode, format string, a ...interface{}) string {
	return colourText(colorCode, fmt.Sprintf(format, a))
}

func Fprintf(colorCode string, w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, Sprintf(colorCode, format, a))
}

// returns the text colored according to the color code (exported version of colourText)
func ColorText(colorCode string, text string) string {
	return colourText(colorCode, text)
}

// --- INTERNAL ONLY ---

// colours a text based on the given colour code
func colourText(colorCode string, text string) string {
	var coloredText string

	if len(colorCode) < 1 {
		return text // nothing to do
	}

	text = strings.Replace(text, "%", "%%", -1)

	switch colorCode[0] {
	// foreground
	case 'r':
		coloredText = Red(text).String()
	case 'g':
		coloredText = Green(text).String()
	case 'y':
		coloredText = Yellow(text).String()
	case 'b':
		coloredText = Blue(text).String()
	case 'x':
		coloredText = Black(text).String()
	case 'm':
		coloredText = Magenta(text).String()
	case 'c':
		coloredText = Cyan(text).String()
	case 'w':
		coloredText = White(text).String()
	case 'd':
		coloredText = Default(text).String()
	// background
	case 'R':
		coloredText = BgRed(text).String()
	case 'G':
		coloredText = BgGreen(text).String()
	case 'Y':
		coloredText = BgYellow(text).String()
	case 'B':
		coloredText = BgBlue(text).String()
	// case 'X': -> not implemented
	case 'M':
		coloredText = BgMagenta(text).String()
	case 'C':
		coloredText = BgCyan(text).String()
	case 'W':
		coloredText = BgWhite(text).String()
	case 'D':
		coloredText = BgDefault(text).String()
	// specials
	case '+':
		coloredText = Bold(text).String()
	case '*':
		coloredText = Italic(text).String()
	case '~':
		coloredText = Reverse(text).String()
	case '_':
		coloredText = Underline(text).String()
	case '#':
		coloredText = Blink(text).String()
	case '?':
		coloredText = Concealed(text).String()
	// default -> panic
	default:
		panic("unknown color code")
	}

	coloredText = strings.Replace(coloredText, "%%", "%", -1)

	return colourText(colorCode[1:], coloredText)
}
