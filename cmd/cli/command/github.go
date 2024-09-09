package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGitHubCommand(createGitHubCommand, changeGitHubCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "github",
		Short: "GitHub",
		Long:  `GitHub.`,
	}

	command.PersistentFlags().IntP("span-size", "s", 0, "Size of the span")
	viper.BindPFlag("spansize", command.PersistentFlags().Lookup("span-size"))

	command.AddCommand(createGitHubCommand)
	command.AddCommand(changeGitHubCommand)

	return command
}
