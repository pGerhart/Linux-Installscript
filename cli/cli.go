package cli

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
)

func GetDesiredSoftware() []string {
	answers := struct {
		SelectedSoftware []string
		InputCorrect     bool
	}{}
	for !answers.InputCorrect {
		err := survey.Ask(softwareQuestions, &answers)
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Info(answers.SelectedSoftware, answers.InputCorrect)
	}
	return []string(answers.SelectedSoftware)
}

var softwareQuestions = []*survey.Question{
	{
		Name: "selectedSoftware",
		Prompt: &survey.MultiSelect{
			Message: "Select software to install on your Machine",
			Options: []string{"Syncthing", "Borg", "KeepassXC"},
		},
	},
	{
		Name: "inputCorrect",
		Prompt: &survey.Confirm{
			Message: "Is your input correct?",
		},
	},
}
