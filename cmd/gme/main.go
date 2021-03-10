package main

import (
	gmesh "github.com/full-stack-gods/gme.sh-cli/internal/pkg/gme-sh"
	"log"
)

func main() {
	cli := gmesh.New()
	if err := cli.Run(); err != nil {
		log.Fatalln("Error:", err)
		return
	}
}
