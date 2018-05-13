package log

import (
	"fmt"
	"runtime"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textMagenta
	textCyan
	textWhite
)
const (
	backgroundBlack = iota + 40
	backgroundRed
	backgroundGreen
	backgroundYellow
	backgroundBlue
	backgroundMagenta
	backgroundCyan
	backgroundWhite
)

//黑色
func Black(str string) string {
	return textColor(textBlack, str)
}

//红色
func Red(str string) string {
	return textColor(textRed, str)
}

//绿色
func Green(str string) string {
	return textColor(textGreen, str)
}

//黄色
func Yellow(str string) string {
	return textColor(textYellow, str)
}

//蓝色
func Blue(str string) string {
	return textColor(textBlue, str)
}

//紫红色
func Magenta(str string) string {
	return textColor(textMagenta, str)
}

//青蓝色
func Cyan(str string) string {
	return textColor(textCyan, str)
}

//白色
func White(str string) string {
	return textColor(textWhite, str)
}
func BlackBg(str string) string {
	return backgroundColor(backgroundBlack, str)
}

func RedBg(str string) string {
	return backgroundColor(backgroundRed, str)
}

func GreenBg(str string) string {
	return backgroundColor(backgroundGreen, str)
}

func YellowBg(str string) string {
	return backgroundColor(backgroundYellow, str)
}

func BlueBg(str string) string {
	return backgroundColor(backgroundBlue, str)
}

func MagentaBg(str string) string {
	return backgroundColor(backgroundMagenta, str)
}

func CyanBg(str string) string {
	return backgroundColor(backgroundCyan, str)
}

func WhiteBg(str string) string {
	return backgroundColor(backgroundWhite, str)
}
func textColor(color int, str string) string {
	if isWindows() {
		return str
	}
	switch color {
	case textBlack:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textBlack, str)
	case textRed:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textRed, str)
	case textGreen:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textGreen, str)
	case textYellow:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textYellow, str)
	case textBlue:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textBlue, str)
	case textMagenta:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textMagenta, str)
	case textCyan:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textCyan, str)
	case textWhite:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", textWhite, str)
	default:
		return str
	}
}
func backgroundColor(color int, str string) string {
	if isWindows() {
		return str
	}
	switch color {
	case backgroundBlack:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundBlack, str)
	case backgroundRed:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundRed, str)
	case backgroundGreen:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundGreen, str)
	case backgroundYellow:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundYellow, str)
	case backgroundBlue:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundBlue, str)
	case backgroundMagenta:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundMagenta, str)
	case backgroundCyan:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundCyan, str)
	case backgroundWhite:
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", backgroundWhite, str)
	default:
		return str
	}
}
func isWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
