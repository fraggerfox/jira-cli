package open

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"

	"github.com/ankitpokhrel/jira-cli/internal/cmdutil"
	"github.com/ankitpokhrel/jira-cli/internal/config"
)

const (
	helpText = `Open opens issue in a browser. If the issue key is not given, it will open the project page.`
	examples = `$ jira open
$ jira open ISSUE-1`
)

// NewCmdOpen is an open command.
func NewCmdOpen() *cobra.Command {
	return &cobra.Command{
		Use:     "open [ISSUE_KEY]",
		Short:   "Open issue in a browser",
		Long:    helpText,
		Example: examples,
		Aliases: []string{"browse", "navigate"},
		Annotations: map[string]string{
			"cmd:main":  "true",
			"help:args": "[ISSUE_KEY]\tIssue key, eg: ISSUE-1",
		},
		Run: open,
	}
}

func open(_ *cobra.Command, args []string) {
	server := config.GetServer()
	project := config.GetProject()

	var url string

	if len(args) == 0 {
		url = fmt.Sprintf("%s/browse/%s", server, project)
	} else {
		url = fmt.Sprintf("%s/browse/%s", server, args[0])
	}

	cmdutil.ExitIfError(browser.OpenURL(url))
}
