package cli

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
)

func GetDesiredPackages(possiblePackages []string) []string {
	answers := struct {
		SelectedSoftware []string
	}{}
	err := survey.Ask(createSoftwareQuestions(possiblePackages), &answers)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return []string(answers.SelectedSoftware)
}

func createSoftwareQuestions(possibleSoftware []string) []*survey.Question {
	var softwareQuestions = []*survey.Question{
		{
			Name: "selectedSoftware",
			Prompt: &survey.MultiSelect{
				Message: "Select software to install on your Machine",
				Options: possibleSoftware,
			},
		},
	}
	return softwareQuestions
}
