// This file is not used. It's only there as a way to decalre all our tools dependencies.
// It's the recommended way of versionning tools:
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// +build tools

package main

import (
	_ "golang.org/x/lint/golint"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
