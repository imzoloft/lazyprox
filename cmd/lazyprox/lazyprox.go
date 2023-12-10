package main

import (
	inout "github.com/imzoloft/lazyprox/internal/io"
	"github.com/imzoloft/lazyprox/internal/io/entry"
	"github.com/imzoloft/lazyprox/internal/proxy/validator"
)

func main() {
	entry.SetupFlag()

	if err := validator.ValidateProxy(); err != nil {
		inout.FatalError(err.Error())
	}
	inout.DisplayOverview()
}
