// xmlfmt - XML Formatter

// The xmlfmt will format the XML string without rewriting the document

package main

////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh -v xmlfmt_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "xmlfmt"
	version  = "0.1.0"
	date     = "2022-02-06"

	// Opts store all the configurable options
	Opts OptsT
)

var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	Opts.Version = showVersion
	Opts.Verbflg = func() {
		Opts.Verbose++
	}

	if _, err := parser.Parse(); err != nil {
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println()
	//DoXmlfmt()
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "xmlfmt - XML Formatter\n")
	fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Antonio Sun\n\n")
	fmt.Fprintf(os.Stderr, "The xmlfmt will format the XML string without rewriting the document\n\nBuilt on %s\nVersion %s\n",
		date, version)
	os.Exit(0)
}

// DoXmlfmt implements the business logic of command `xmlfmt`
func DoXmlfmt() error {
	return nil
}

