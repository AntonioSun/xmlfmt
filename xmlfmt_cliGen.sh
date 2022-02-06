 pwd=`pwd`
 cd $GOPATH/src/github.com/go-easygen/easygen/test
 easygen commandlineGoFlags.header,commandlineGoFlags.ityped.tmpl,commandlineGoFlags "$pwd/xmlfmt"_cli | gofmt > "$pwd/xmlfmt"_cliDef.go
