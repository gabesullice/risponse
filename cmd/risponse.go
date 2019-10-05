package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	server "github.com/gabesullice/risponse/lib/server"
)

var (
	baseDir string
	port    int
)

func init() {
	flag.StringVar(&baseDir, "d", ".", "Specifies an alternate directory from which to serve resources.")
	flag.IntVar(&port, "p", 8081, "Specifies an alternate directory from which to serve resources.")
	flag.Parse()
}

func main() {
	if baseDir != "." {
		if err := os.Chdir(baseDir); err != nil {
			log.Fatalln(err)
		}
	}
	server.ListenAndServe(fmt.Sprintf(":%d", port), server.LoadConfigFromFile("./config.json"))
}
