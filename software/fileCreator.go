package software

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
	// answer += "# !!!!!!!!!!!!!ATTENTION!!!!!!!!!!!!!\n"
	answer += "\n"
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
