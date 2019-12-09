# xmlfmt

About XML formatting, see https://github.com/go-xmlfmt/xmlfmt.

## Download binaries

- The latest binary executables are available under  
https://bintray.com/antoniosun/bin/xmlfmt, or directly under  
https://bintray.com/version/files/antoniosun/bin/xmlfmt  
as the result of the Continuous-Integration process.
- I.e., they are built during every git push, automatically by [travis-ci](https://travis-ci.org/), right from the source code, truly WYSIWYG.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `xmlfmt-linux-amd64` file. If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- You may want to rename it to a shorter name instead, e.g., `xmlfmt`, after downloading it. 


## Debian package

Available at https://bintray.com/antoniosun/deb/xmlfmt,  
or directly at  https://dl.bintray.com/antoniosun/deb:

```
echo "deb [trusted=yes] https://dl.bintray.com/antoniosun/deb all main" | sudo tee /etc/apt/sources.list.d/antoniosun-debs.list
sudo apt-get update

sudo chmod 644 /etc/apt/sources.list.d/antoniosun-debs.list
apt-cache policy xmlfmt

sudo apt-get install -y xmlfmt
```



## Install Source

To install the source code instead:

```
go get github.com/AntonioSun/xmlfmt
```


## Author(s) & Contributor(s)

- [Antonio SUN](https://github.com/AntonioSun)

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
