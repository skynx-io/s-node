package main

import (
	"fmt"
	// "log"

	"skynx.io/s-lib/pkg/version"
	"skynx.io/s-node/internal/app/node/cmd"
)

func main() {
	// if err := cmd.ConsoleInit(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Printf("%s %s ", version.NODE_NAME, version.GetVersion())

	cmd.Execute()
}
