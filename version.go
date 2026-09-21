package modtest

import (
	_ "embed"
)

//go:embed VERSION.txt
var Version string
