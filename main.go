/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/bomgar/fbmirrors/cmd"
)

var version string = "dev"

func main() {
	cmd.Execute(version)
}
