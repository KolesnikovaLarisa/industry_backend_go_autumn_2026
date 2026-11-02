package main

import "strings"

func greet(name string) string {

	trimmedName := strings.TrimSpace(name)

	if trimmedName == "" {
		trimmedName = "World"
	}

	return "Hello, " + trimmedName + "!"
}
