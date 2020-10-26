package main

import (
	"installscript/cli"
	"installscript/constants"
	"installscript/helpers"
	"installscript/software"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	sw, err := software.ParseJSON("testSoftware.json")
	if err != nil {
		log.Fatal(err)
	}
	/*
		distro, err := software.GetDistro()
		if err != nil {
			log.Fatal(err)
		}
	*/
	distro := "Debian"

	desiredPackages := cli.GetDesiredPackages(sw.PackageList())
	script, warning := sw.CreateInstallScript(desiredPackages, distro)

	log.Debug("Evaluated Distro: ", distro)
	helpers.WriteToFile(constants.OUTFILE, script)
	err = helpers.MakeExecutable(constants.OUTFILE)
	if err != nil {
		log.Fatal(err)
	}
	if warning != "" {
		log.Warning("\n" + warning)
	}
}
