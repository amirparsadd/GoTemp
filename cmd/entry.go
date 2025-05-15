package cmd

import (
	"fmt"
	"gotemp/converter"
	"gotemp/registry"
	"strconv"
	"strings"
)

type CommandLineResultNumber int

const (
	SUCCESS CommandLineResultNumber = iota
	FAIL
)

type CommandLineResult struct {
	ResultNumber CommandLineResultNumber
	ErrorMessage string
}

func CommandLineEntry(args []string) CommandLineResult {
	if(len(args) != 4) {
		return CommandLineResult{FAIL, "Invalid Args Length."}
	}

	var value, valueErr = strconv.ParseFloat(args[1], 64)

	if(valueErr != nil) {
		return CommandLineResult{FAIL, "Invalid Value. Must Be Float64"}
	}

	var from, fromOk = registry.UnitRegistry[args[2]]
	var to, toOk = registry.UnitRegistry[args[3]]

	var allowedUnitsString = strings.Join(registry.GetAllowedUnits(), ", ")

	if(!fromOk) {
		return CommandLineResult{FAIL, "Invalid From Unit! Allowed Units: " + allowedUnitsString}
	}

	if(!toOk) {
		return CommandLineResult{FAIL, "Invalid To Unit! Allowed Units: " + allowedUnitsString}
	}

	var final, err = converter.Convert(value, from.Symbol(), to.Symbol())

	if(err != nil) {
		return CommandLineResult{FAIL, "Unexpected Error: " + err.Error()}
	}

	fmt.Printf("%v%v %v %v%v\n",
		value,         // Starting value
		from.Symbol(), // Starting value Symbol
		GetArrow(),    // Arrow. Changes based on terminal capabilities
		final,         // Final value
		to.Symbol())   // Final value symbol

	return CommandLineResult{ SUCCESS, "" }
}