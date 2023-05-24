package commands

import (
	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/commands/app"
	"fmt"
	"os"

	"bitbucket.org/dreamplug-backend/ProjectGeneratorCli/internal/ansi"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ProjectGeneratorCli",
	Short: ansi.ColoredBoldStatus("A CLI to help you create new java project structure from templates"),
	Long:  ansi.ColoredBoldStatus("The command-line tool to interact with project generator."),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(app.NewCreateCmd().Reqs.Cmd)
}
