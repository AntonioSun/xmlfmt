// xmlfmt - XML Formatter

// The xmlfmt will format the XML string without rewriting the document

package main

////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2016-2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh xmlfmt_cliGen.sh

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-easygen/go-flags"
	"github.com/go-xmlfmt/xmlfmt"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "xmlfmt"
	version  = "1.1.2"
	date     = "2022-08-12"

	// opts store all the configurable options
	opts optsT
)

var parser = flags.NewParser(&opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	opts.Version = showVersion
	opts.Verbflg = func() {
		opts.Verbose++
	}

	if _, err := parser.Parse(); err != nil {
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println()
	DoXmlfmt()
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "xmlfmt - XML Formatter, version %s\n", version)
	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
	fmt.Fprintf(os.Stderr, "Copyright (C) 2016-2022, Antonio Sun\n\n")
	fmt.Fprintf(os.Stderr, "The xmlfmt will format the XML string without rewriting the document\n")
	os.Exit(0)
}

// DoXmlfmt implements the business logic of command `xmlfmt`
func DoXmlfmt() error {
	var data []byte
	var err error
	if opts.Filei == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(opts.Filei)
	}
	abortOn("Input", err)

	fmt.Println(xmlfmt.FormatXML(string(data),
		opts.Prefix, opts.Indent, opts.Nested))

	return nil
}

// abortOn will quit on anticipated errors gracefully without stack trace
func abortOn(errCase string, e error) {
	if e != nil {
		fmt.Printf("[%s] %s error: %v\n", progname, errCase, e)
		os.Exit(1)
	}
}
