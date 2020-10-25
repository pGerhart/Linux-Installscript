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
	VaryingCommands map[string][]string
	DefaultCommand  []string
}

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

// EvaluateCommand creates an command string from a package
func (pkg Package) EvaluateCommand(distro string) (string, error) {
	if len(pkg.VaryingCommands) == 0 {
		return cmdToString(pkg.DefaultCommand), nil
	}
	if val, ok := pkg.VaryingCommands[distro]; ok {
		return cmdToString(val), nil
	}
	return cmdToString(pkg.DefaultCommand), &DistroNotSupportedError{pkg, distro}
}

func cmdToString(cmd []string) string {
	var answer string
	for _, line := range cmd {
		answer += line + "\n"
	}
	return answer
}

// DistroNotSupportedError will be raised when there is no custom command
// for the wanted distro
type DistroNotSupportedError struct {
	Pkg    Package
	Distro string
}

func (e *DistroNotSupportedError) Error() string {
	return fmt.Sprintf("Package %s has no custom command for distro %s", e.Pkg.Name, e.Distro)
}

// GetDistro checks distro
func GetDistro() (string, error) {
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

func (sw Software) PackageList() []string {
	answer := make([]string, len(sw.Packages))
	for index, pkg := range sw.Packages {
		answer[index] = pkg.Name
	}
	return answer
}
