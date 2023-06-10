# InvFinSDK

SDK for InvFin endpoints

# Go API client for Inversiones & Finanzas (inversionesyfinanzas.xyz)

<!-- ![GitHub release (latest by date)](https://img.shields.io/github/v/release/)
![GitHub](https://img.shields.io/github/license/) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/)  -->
<!-- [![Go Report Card](https://goreportcard.com/badge/github.com/)](https://goreportcard.com/report/github.com/) -->


## Overview

- API documentation [InvFin](https://inversionesyfinanzas.xyz/api/api-documentacion/)
- API version: 1
- Package version: 0.1.0


## Support Methods

- [Company](https://inversionesyfinanzas.xyz/api/api-documentacion/#lista-de-terminos)
- [Exchange](https://inversionesyfinanzas.xyz/api/api-documentacion/#lista-de-exchanges)
- [Term](https://inversionesyfinanzas.xyz/api/api-documentacion/#lista-de-terminos)
- [Industry](https://inversionesyfinanzas.xyz/api/api-documentacion/#lista-de-terminos)
- [Sector](https://inversionesyfinanzas.xyz/api/api-documentacion/#lista-de-terminos)
- [Superinvestor](https://inversionesyfinanzas.xyz/api/api-documentacion/#superinversores)


## Installation
This project requires Go to be installed.

On OS X with Homebrew you can just run `brew install go`.

For other OS you can go to the [docs](https://go.dev/doc/install)

Once Go installed you can run:

```sh
go get -u github.com/InvFin/Go-SDK
```

## Example:

```go
package main

import (
	"log"

	"github.com/InvFin/Go-SDK/api"
)

func main() {
	APIClient, err := api.NewAPI("YOUR_KEY")
	if err != nil {
		log.Println("Error init api client: " + err.Error())
	}
	// Get all companies
	allCompanies := APIClient.Company.GetAllCompanies()
	log.Println(allCompanies)
}
```