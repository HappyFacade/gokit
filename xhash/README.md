# GoKit - xhash

Hash kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xhash"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xhash)

## Example

### get md5 of string

    h := xhash.Md5("12345678")
    fmt.Println(h.Hex())
    fmt.Println(h.B64())

### get md5 of file

    h, err := xhash.FileMd5("12345678")
    if err != nil {
        panic(err)
    }
    fmt.Println(h.Hex())
    fmt.Println(h.B64())

## LICENSE

Copyright 2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)