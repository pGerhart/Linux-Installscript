package main

import (
	"installscript/cli"
	"installscript/helpers"
	"installscript/software"
	"os"

	"github.com/akamensky/argparse"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	args := parseArgs()
	if args.Verbose {
		log.SetLevel(log.DebugLevel)
	}
	log.Debug("Evaluated Distro: ", args.Distro)

	sw, err := software.ParseJSON()
	if err != nil {
		log.Fatal(err)
	}

	desiredPackages := cli.GetDesiredPackages(sw.PackageList())
	script, warning := sw.CreateInstallScript(desiredPackages, args.Distro)

	if args.IgnoreWarnings {
		script += warning
	} else if warning != "" {
		log.Warning("\n" + software.MissingDistrosHint() + warning)
	}

	helpers.WriteToFile(args.Outpath, script)
	err = helpers.MakeExecutable(args.Outpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Writing install script into file: ", args.Outpath)

}

func parseArgs() Args {
	parser := argparse.NewParser("InstallScript", "Creates an install Script for Linux.")
	outpath := parser.String("o", "outpath", &argparse.Options{
		Required: false,
		Help:     "Path in which the install script is written",
		Default:  "installScript.sh",
	})
	ignoreWarnings := parser.Flag("i", "ignoreWarnings", &argparse.Options{
		Required: false,
		Help:     "ignores Warnings and writes all Commands to the Script",
		Default:  false,
	})
	verbose := parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Set output level to debug",
		Default:  false,
	})
	customDistro := parser.String("d", "distro", &argparse.Options{
		Required: false,
		Help:     "Custom Distro to create script for.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}
	distro := *customDistro
	if distro == "" {
		distro, err = software.GetDistro()
		if err != nil {
			log.Fatal(err)
		}
	}

	return Args{
		Outpath:        *outpath,
		IgnoreWarnings: *ignoreWarnings,
		Verbose:        *verbose,
		Distro:         distro,
	}
}

type Args struct {
	Outpath        string
	IgnoreWarnings bool
	Verbose        bool
	Distro         string
}
