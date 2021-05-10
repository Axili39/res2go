![Build & Test](https://github.com/Axili39/res2go/workflows/Build%20&%20Test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/Axili39/res2go/badge.svg?branch=master)](https://coveralls.io/github/Axili39/res2go?branch=master)

Resource Package Generator for Golang
=====================================

Very simple generator used to embed some file into executable. Useful to store template files.

```shell
go get github.com/Axili39/res2go
go build
go install
```

usage:

```shell
Usage of ./res2go:
  -o string
        output file (default "resources.go")
  -package string
        package name (default "resources")
  -prefix string
        add prefix to all exported function

examples:
res2go ./resources

res2go ./templates/index.html ./templates/style.js

res2go -package myresources index.html

res2go -package myresources -o rsrc/rsrc.go resources/*.templates

res2go -package main -o rsrc.go -prefix Rsrc resources/*
```

Example in Go code:

```go
//go:generate res2go -package resources -o resources/resources.go resources/*.templates

import (
  "./resources"
)

func main() {
  resources.Init()

  myfile = resources.Files["resources/example.template"]
  fmt.Println(string(myfile))

}
```
