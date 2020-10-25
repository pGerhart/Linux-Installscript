package software

import (
	log "github.com/sirupsen/logrus"
)

// CreateInstallScript evaluates all the variables to the final script
func (sw Software) CreateInstallScript(desiredPackages []string, distro string) string {
	answer := "#!/bin/bash \n"
	answer += sw.EvaluateUpdateCommand(distro)
	answer += sw.EvaluateVariables()

	var missingDistro string

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
		answer += missingDistrosHint()
		answer += missingDistro

	}
	return answer
}

func createPackageBlog(cmd, name string) string {
	answer := "# ----------------------------------------------------\n"
	answer += "# " + name + "\n"
	answer += "\n"
	answer += cmd
	answer += "\n"
	answer += "# ----------------------------------------------------\n\n\n"
	return answer
}

func missingDistrosHint() string {
	answer := "# !!!!!!!!!!!!!ATTENTION!!!!!!!!!!!!!\n"
	answer += "#\n"
	answer += "# The following packages are not supported but those commands may work: \n"
	answer += "#\n"
	answer += "# !!!!!!!!!!!!!ATTENTION!!!!!!!!!!!!!\n"
	return answer
}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
