# Pix Generator

<img src="gopher.png" width="200" />

Modern rewrite of [go-pix](https://github.com/fonini/go-pix)

## Requirements

- Golang 1.24.1 or higher
  - [Download](https://go.dev/dl/)

## Usage

```go
package main

import (
    "fmt"

    "github.com/i9si-sistemas/pix"
)

func main() {
   options := pix.Options{
        Name: "Gabriel Luiz",
        Key: "7a067b11-bce7-406f-af8a-2dcf82c429d6",
        City: "Caruaru",
        Amount: 20.67, // optional
        Description: "Invoice #5374", // optional
        TransactionID: "***", // optional
    }

    copyPaste, err := pix.New(options)
    if err != nil {
        panic(err)
    }

    fmt.Println(copyPaste) 

    optionsFromCode, err := pix.Read(copyPaste)
    if err != nil {
        panic(err)
    }

    fmt.Println(optionsFromCode)
}
```