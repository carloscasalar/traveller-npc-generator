//go:build tools
// +build tools

package main

import (
	_ "github.com/alvaroloes/enumer"
	_ "github.com/moznion/go-errgen/cmd/errgen"
	_ "github.com/moznion/gonstructor/cmd/gonstructor"
)
