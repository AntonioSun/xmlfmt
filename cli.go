////////////////////////////////////////////////////////////////////////////
// Program: xmlfmtC
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2021, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// xmlfmtC

type rootT struct {
	cli.Helper
	Filei  *clix.Reader `cli:"*f,file" usage:"The xml file to read from (or stdin)"`
	Prefix string       `cli:"p,prefix" usage:"each element begins on a new line and this prefix"`
	Indent string       `cli:"i,indent" usage:"indent string for nested elements" dft:"  "`
	Nested bool         `cli:"n,nested" usage:"nested tags in comments"`
}

var root = &cli.Command{
	Name: "xmlfmtC",
	Desc: "XML Formatter\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2021, Antonio Sun",
	Text: "The xmlfmt will format the XML string without rewriting the document",
	Argv: func() interface{} { return new(rootT) },
	Fn:   xmlfmtC,

	NumOption: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Filei	*clix.Reader
//  	Prefix	string
//  	Indent	string
//  	Nested	bool
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "xmlfmtC"
//          version   = "0.1.0"
//          date = "2021-12-06"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func xmlfmtC(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
