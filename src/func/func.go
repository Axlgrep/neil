package main

import (
	"fmt"
	"strings"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	addPng := func(name string) string { return name + ".png" }
	addJpg := func(name string) string { return name + ".jpg" }
	fmt.Println(addPng("filename"), addJpg("filename"))

	cusfunc := MakeAddSuffix(".axl")
	fmt.Println(cusfunc("filename"), cusfunc("filename.axl"))
}
