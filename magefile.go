//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Run App Locally
func RunLocal() error {
	return sh.RunV("go", "run", "cmd/app/main.go")
}
