package software

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// Package represents a software and the commands such that
// it can be intstalled on Supported Distros
type Package struct {
	Name            string
	log "github.com/sirupsen/logrus"
)

// Software reads an Json Config
type Software struct {
	Packages       []Package
	Variables      map[string]string
	UpdateCommands map[string]string
}

// EvaluateVariables creates the Variables for the Bash Script
func (software Software) EvaluateVariables() string {
	var answer string
	for key, value := range software.Variables {
		answer += fmt.Sprintf(`%s='%s'`, key, value)
		answer += "\n"
	}
	answer += "\n"
	return answer
}

// EvaluateUpdateCommand creates the UpdateCommand for the Bash Script
func (software Software) EvaluateUpdateCommand(distro string) string {
	if cmd, found := software.UpdateCommands[distro]; found {
		return cmd + "\n"
	}
	return software.UpdateCommands["Default"] + "\n"
}

func (sw Software) PackageList() []string {
	answer := make([]string, len(sw.Packages))
	for index, pkg := range sw.Packages {
		answer[index] = pkg.Name
	}
	return answer
}

// CreateInstallScript evaluates all the variables to the final script
func (sw Software) CreateInstallScript(desiredPackages []string, distro string) (string, string) {
	answer := "#!/bin/bash \n"
	answer += sw.EvaluateUpdateCommand(distro)
	answer += sw.EvaluateVariables()

	var warning, missingDistro string

	for _, pkg := range sw.Packages {
		if _, found := find(desiredPackages, pkg.Name); !found {
			continue
		}

		cmd, err := pkg.EvaluateCommand(distro)
		if err != nil {
			log.Error(err)
			missingDistro += createPackageBlog(cmd, pkg.Name)
		} else {
			answer += createPackageBlog(cmd, pkg.Name)
		}
	}
	if missingDistro != "" {
		warning = missingDistrosHint() + missingDistro
	}

	return answer, warning
}

// GetDistro checks distro
func GetDistro() (string, error) {
	if runtime.GOOS != "Linux" {
		return "", &OSNotSupportedError{runtime.GOOS}
	}
	var out bytes.Buffer
	cmd := exec.Command("lsb_release", "-a")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`Distributor\sID:\s.*\n`)
	distro := re.FindString(out.String())
	distro = strings.Trim(distro, "Distributor ID:")
	distro = strings.Trim(distro, "\n")
	distro = strings.Trim(distro, "\t")
	return distro, nil
}
