package llogger

import (
	"fmt"
	"os"
)

var TraceLevelColor = LevelColor{r: 250, g: 250, b: 250}
var DebugLevelColor = LevelColor{r: 0, g: 255, b: 0}
var ErrorLevelColor = LevelColor{r: 255, g: 0, b: 0}
var InfoLevelColor = LevelColor{r: 250, g: 250, b: 250}
var WarnLevelColor = LevelColor{r: 255, g: 255, b: 0}

type LevelColor struct {
	r int
	g int
	b int
}

func (lc *LevelColor) setColor() {
	fmt.Fprintf(os.Stdout, "\u001B[48;2;0;0;0m\u001B[38;2;%d;%d;%dm", lc.r, lc.g, lc.b)
}

func (lc *LevelColor) resetColor() {
	fmt.Fprintf(os.Stdout, "\u001B[0m")
}

func getColorByLevel(level int) LevelColor {
	switch level {
	case TraceLevel:
		return TraceLevelColor
	case DebugLevel:
		return DebugLevelColor
	case ErrorLevel:
		return ErrorLevelColor
	case WarnLevel:
		return WarnLevelColor
	default:
		return InfoLevelColor
	}
}
