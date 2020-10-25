package main

import (
	"installscript/cli"
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
	distro, err := software.GetDistro()
	if err != nil {
		log.Fatal(err)
	}

	desiredPackages := cli.GetDesiredPackages(sw.PackageList())
	script, warning := sw.CreateInstallScript(desiredPackages, distro)

	log.Debug("Evaluated Distro:", distro)
	log.Info("Created Script:\n" + script)
	if warning != "" {
		log.Warning("\n" + warning)
	}
}
