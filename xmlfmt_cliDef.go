// xmlfmt - XML Formatter
//
// The xmlfmt will format the XML string without rewriting the document

package main

////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

import (
//  	"fmt"
//  	"os"

//  	"github.com/go-easygen/go-flags"
)

// Template for main starts here

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "xmlfmt"
//          version   = "0.1.0"
//          date = "2022-02-07"

//  	// Opts store all the configurable options
//  	Opts OptsT
//  )
//
//  var parser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	Opts.Version = showVersion
//  	Opts.Verbflg = func() {
//  		Opts.Verbose++
//  	}
//
//  	if _, err := parser.Parse(); err != nil {
//  		fmt.Println()
//  		parser.WriteHelp(os.Stdout)
//  		os.Exit(1)
//  	}
//  	fmt.Println()
//  	//DoXmlfmt()
//  }
//
//  func showVersion() {
//   	fmt.Fprintf(os.Stderr, "xmlfmt - XML Formatter, version %s\n", version)
//  	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
//   	fmt.Fprintf(os.Stderr, "Copyright (C) 2022, Antonio Sun\n\n")
//  	fmt.Fprintf(os.Stderr, "The xmlfmt will format the XML string without rewriting the document\n\n")
//  	os.Exit(0)
//  }
// Template for main ends here

// DoXmlfmt implements the business logic of command `xmlfmt`
//  func DoXmlfmt() error {
//  	return nil
//  }

// Template for type define starts here

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Filei   string `short:"f" long:"file" env:"XMLFMT_FILEI" description:"The xml file to read from (or \"-\" for stdin)" required:"true"`
	Prefix  string `short:"p" long:"prefix" env:"XMLFMT_PREFIX" description:"Each element begins on a new line and this prefix"`
	Indent  string `short:"i" long:"indent" env:"XMLFMT_INDENT" description:"Indent string for nested elements" default:"  "`
	Nested  bool   `short:"n" long:"nested" env:"XMLFMT_NESTED" description:"Nested tags in comments"`
	Verbflg func() `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
	Verbose int
	Version func() `short:"V" long:"version" description:"Show program version and exit"`
}

// Template for type define ends here
