package main

import (
	"os"
)

func main() {
	println("os.Args[0]:", os.Args[0])
	exeName, err := os.Executable()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println("os.Executable():", exeName)
}
