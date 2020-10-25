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
	distro, err := software.GetDistro()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Evaluated Distro:\t", distro)

	for _, pkg := range sw.Packages {
		cmd, err := pkg.EvaluateCommand(distro)
		if err != nil {
			log.Error(err)
		}
		log.Debug("<Package: ", pkg.Name, "\tCommand: ", cmd[:len(cmd)-1], ">")
	}
}
