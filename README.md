[![Build Status](https://travis-ci.org/lycis/ansi.png?branch=master)](https://travis-ci.org/lycis/ansi)
go-ansi 
=============

Installing
-------
```shell
    $ go get github.com/str1ngs/ansi/color
```
Example
-------
```go
    package main

    import (
        "fmt"
        . "github.com/str1ngs/ansi/color"
    )

    func main() {
        fmt.Printf("%s\n", Blue("techno color"))
    }
```
