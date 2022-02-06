// xmlfmt - XML Formatter
//
// The xmlfmt will format the XML string without rewriting the document

package main

////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

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
