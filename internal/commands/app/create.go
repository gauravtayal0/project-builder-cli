package app

import (
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/prompt"
	"github.com/spf13/cobra"

	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/ansi"
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/javaAppGenerator"
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/requests"
)

type createCmd struct {
	Reqs requests.Base
}

func NewCreateCmd() *createCmd {
	gc := &createCmd{}
	gc.Reqs.Cmd = &cobra.Command{
		Use:     "create",
		Short:   ansi.ColoredBoldStatus("create a create application"),
		Long:    ansi.ColoredBoldStatus("This command will generate a new application"),
		Example: ``,
		RunE:    createCmdRunner,
	}

	return gc
}

// createCmdRunner is the handler for "create project".
func createCmdRunner(cmd *cobra.Command, args []string) error {

	co := &prompt.Options{}
	prompt.CreateOptions(co)

	if co.Language == "Java" && co.Framework == "Spring" {
		sp := javaAppGenerator.SpringAppGenerator{
			TargetDirectory: co.TargetDirectory,
			ProjectName:     co.ProjectName,
		}
		sp.CreateProjectStructure()
	}

	return nil
}
