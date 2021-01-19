package me

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ankitpokhrel/jira-cli/internal/config"
)

// NewCmdMe is a me command.
func NewCmdMe() *cobra.Command {
	return &cobra.Command{
		Use:   "me",
		Short: "Displays configured jira user",
		Long:  "Displays configured jira user.",
		Run:   me,
	}
}

func me(*cobra.Command, []string) {
	fmt.Println(config.GetLogin())
}
