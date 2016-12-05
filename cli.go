////////////////////////////////////////////////////////////////////////////
// Program: xmlfmt
// Purpose: XML Formatter
// Authors: Antonio Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// xmlfmt

type rootT struct {
	cli.Helper
	Filei  *clix.Reader `cli:"*f,file"usage:"The xml file to read from (or stdin)"`
	Prefix string       `cli:"p,prefix"usage:"each element begins on a new line and this prefix"`
	Indent string       `cli:"i,indent"usage:"indent string for nested elements" dft:"  "`
}

var root = &cli.Command{
	Name: "xmlfmt",
	Desc: "XML Formatter\nbuilt on " + buildTime,
	Text: "The xmlfmt will format the XML string without rewriting the document",
	Argv: func() interface{} { return new(rootT) },
	Fn:   xmlfmt,

	NumOption: cli.AtLeast(1),
}

// func main() {
// 	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
// 	//NOTE: You can set any writer implements io.Writer
// 	// default writer is os.Stdout
// 	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	fmt.Println("")
// }

// func xmlfmt(ctx *cli.Context) error {
// 	ctx.JSON(ctx.RootArgv())
// 	ctx.JSON(ctx.Argv())
// 	fmt.Println()

// 	return nil
// }
