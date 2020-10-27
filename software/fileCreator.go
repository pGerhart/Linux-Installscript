package software

import "fmt"

var horizontalLine = "# ----------------------------------------------------\n"

func createPackageBlog(cmd, name string) string {
	answer := horizontalLine
	answer += fmt.Sprintf(`#%s`, name) + "\n\n"
	answer += fmt.Sprintf(`echo "installing package %s"`, name) + "\n"
	answer += fmt.Sprintf(`%s=$(%s)`, name, cmd) + "\n"
	answer += fmt.Sprintf(`%s_ret=$?`, name) + "\n"
	answer += fmt.Sprintf(`echo "$%s" > /dev/null`, name) + "\n"
	answer += fmt.Sprintf(`%s`, horizontalLine) + "\n\n"
	return answer
}

func MissingDistrosHint() string {
	answer := "# !!!!!!!!!!!!!ATTENTION!!!!!!!!!!!!!\n"
	answer += "#\n"
	answer += "# The following packages are not supported but those commands may work: \n"
	answer += "#\n"
	// answer += "# !!!!!!!!!!!!!ATTENTION!!!!!!!!!!!!!\n"
	answer += "\n"
	return answer
}

func checkSuccessfullInstall(name string) string {
	answer := fmt.Sprintf(`# %s Test`, name) + "\n"
	answer += fmt.Sprintf(`if [ $%s_ret -ne 0 ]; then`, name) + "\n"
	answer += fmt.Sprintf(`printf "%s# %s \n#\n" >> "$LOG_FILE"`, horizontalLine, name) + "\n"
	answer += fmt.Sprintf(`echo "Return code for package %s was not zero but $%s_ret" | tee -a "$LOG_FILE"`, name, name) + "\n"
	answer += fmt.Sprintf(`echo "${%s}" 2>&1 >> "$LOG_FILE"`, name) + "\n"
	answer += "fi \n\n"
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
