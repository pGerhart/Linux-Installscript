package main

import (
	"installscript/software"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	//log.Debug(cli.GetDesiredSoftware())
	sw, err := software.ParseJSON("testSoftware.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(sw)
}
