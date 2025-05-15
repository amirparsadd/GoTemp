package main

import (
	"fmt"
	"gotemp/registry"
	"gotemp/cmd"
	"os"
)

func main() {
	// Args
	// 0 - Execution (unrelated)
	// 1 - Value
	// 2 - From
	// 3 - To

	registry.RegisterBaseUnits()

	exit(cmd.CommandLineEntry(os.Args))
}

func exit(result cmd.CommandLineResult) {
	if(result.ResultNumber == cmd.FAIL) {
		fmt.Println(result.ErrorMessage)
	}
	os.Exit(int(result.ResultNumber))
}