/**
 * This file is for generic message types. Not specific messages.
 */

package output

import (
	"fmt"
	"github.com/kortschak/ct"
	padUtf8 "github.com/willf/pad/utf8"
)

var (
	err      = (ct.Fg(ct.Red)).Paint
	warn       = (ct.Fg(ct.Yellow)).Paint
	highlight  = (ct.Fg(ct.BoldCyan) | ct.Bold).Paint
)

var TAB = "  "

func Error(text string) {
	fmt.Println(err(padUtf8.Left("🚨  Error:", 12, " ")), text)
}

func ErrorAsObject(err error) {
	Error(fmt.Sprintf("%v", err))
}

func Search(text string) {
	fmt.Println("🔍  " + text)
}

func Write(text string) {
	fmt.Println("✍️  " + text)
}

func Space() {
	fmt.Print("\n")
}

func Warning(text string) {
	fmt.Println(warn("⚠️  " + text))
}
