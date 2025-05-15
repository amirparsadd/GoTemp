package cmd

import (
	"os"
	"strings"
)

func TerminalSupportsFancyText() bool {
	var term = os.Getenv("TERM")
	var lang = os.Getenv("LANG")
	var lcAll = os.Getenv("LC_ALL")

	if strings.Contains(term, "xterm") || strings.Contains(term, "tmux") {
		return true
	}
	if strings.Contains(lang, "UTF-8") || strings.Contains(lcAll, "UTF-8") {
		return true
	}
	return false
}

func GetArrow() string {
	const FANCY_ARROW = "→"
	const LEGACY_ARROW = "=>"

	if(TerminalSupportsFancyText()) {
		return FANCY_ARROW
	}

	return LEGACY_ARROW
}