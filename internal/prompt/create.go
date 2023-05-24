package prompt

import (
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/utils"
	"github.com/manifoldco/promptui"
)

type Options struct {
	Language        string
	Framework       string
	Protocol        string
	TargetDirectory string
	ProjectName     string
}

func CreateOptions(options *Options) error {
	languagePrompt := promptui.Select{
		Label: "Programming Language",
		Items: []string{"Java", "Golang"},
	}

	_, options.Language, _ = languagePrompt.Run()

	frameworkPrompt := promptui.Select{
		Label: "Project Framework",
		Items: []string{"Spring"},
	}

	_, options.Framework, _ = frameworkPrompt.Run()

	protocolPrompt := promptui.Select{
		Label: "Choose protocol",
		Items: []string{"http", "gRPC"},
	}
	_, options.Protocol, _ = protocolPrompt.Run()

	directoryPrompt := promptui.Prompt{
		Label:     "Target Directory",
		Default:   utils.WorkDir(),
		AllowEdit: true,
	}

	options.TargetDirectory, _ = directoryPrompt.Run()

	projectNamePrompt := promptui.Prompt{
		Label:     "Project Name",
		Default:   "test",
		AllowEdit: true,
	}

	options.ProjectName, _ = projectNamePrompt.Run()

	return nil
}
