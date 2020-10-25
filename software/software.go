package software

import "errors"

// Package represents a software and the commands such that
// it can be intstalled on Supported Distros
type Package struct {
	Name            string
	VaryingCommands map[string][]string
	DefaultCommand  []string
}

// Software reads an Json Config
type Software struct {
	Packages []Package
}

// EvaluateCommand creates an command string from a package
func (pkg Package) EvaluateCommand(distro string) (string, error) {
	if len(pkg.VaryingCommands) == 0 {
		return cmdToString(pkg.DefaultCommand), nil
	}
	if val, ok := pkg.VaryingCommands[distro]; ok {
		return cmdToString(val), nil
	}
	return cmdToString(pkg.DefaultCommand), errors.New("")
}

func cmdToString(cmd []string) string {
	var answer string
	for _, line := range cmd {
		answer += line + "\n"
	}
	return answer
}
