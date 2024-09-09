package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGitCommand(createGitCommand, changeGitCommand *cobra.Command) *cobra.Command {
	command := &cobra.Command{
		Use:   "git",
		Short: "Git",
		Long:  `Git.`,
	}

	command.PersistentFlags().IntP("span-size", "s", 0, "Size of the span")
	viper.BindPFlag("spansize", command.PersistentFlags().Lookup("span-size"))

	command.AddCommand(createGitCommand)
	command.AddCommand(changeGitCommand)

	return command
}
