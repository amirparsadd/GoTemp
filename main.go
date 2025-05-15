package main

import (
	"fmt"
	"gotemp/converter"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Args
	// 0 - Execution (unrelated)
	// 1 - Value
	// 2 - From
	// 3 - To

	converter.RegisterBaseUnits()

	if(len(os.Args) != 4) {
		exitWithError("Invalid Args Length.")
	}

	var value, valueErr = strconv.ParseFloat(os.Args[1], 64)

	if(valueErr != nil) {
		exitWithError("Invalid Value. Must Be Float64")
	}

	var from, fromOk = converter.UnitRegistry[os.Args[2]]
	var to, toOk = converter.UnitRegistry[os.Args[3]]

	var allowedUnitsString = strings.Join(converter.GetAllowedUnits(), ", ")

	if(!fromOk) {
		exitWithError("Invalid From Unit! Allowed Units: " + allowedUnitsString)
	}

	if(!toOk) {
		exitWithError("Invalid To Unit! Allowed Units: " + allowedUnitsString)
	}

	var final, err = converter.Convert(value, from.Symbol(), to.Symbol())

	if(err != nil) {
		exitWithError("Unexpected Error: " + err.Error())
	}

	fmt.Printf("%v%v %v %v%v\n", value, from.Symbol(), getArrow(), final, to.Symbol())
}

// getArrow returns the appropriate arrow symbol, attempting to detect terminal capabilities.
func getArrow() string {
	var term = os.Getenv("TERM")
	var lang = os.Getenv("LANG")
	var lcAll = os.Getenv("LC_ALL")

	if strings.Contains(term, "xterm") || strings.Contains(term, "tmux") {
		return "→"
	}
	if strings.Contains(lang, "UTF-8") || strings.Contains(lcAll, "UTF-8") {
		return "→"
	}
	return "->"
}

func exitWithError(text string) {
	fmt.Println(text)
	os.Exit(1)
}