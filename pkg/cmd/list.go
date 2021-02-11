package cmd

import (
	"os"

	"github.com/go-piv/piv-go/piv"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ListCmd struct {
	Cmd  *cobra.Command
	Args []string
}

func NewListCmd() *cobra.Command {
	c := &ListCmd{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list the connected yubikeys",
		Long:    "",
		Example: "",
		Aliases: []string{"l", "li", "lis", "list"},
		Run: func(cmd *cobra.Command, args []string) {
			c.Cmd = cmd
			c.Args = args
			err := c.Run()
			if err != nil {
				logrus.Fatalf("unable to run command: %s", err)
			}
		},
		Args: cobra.MaximumNArgs(1),
	}

	return cmd
}

func (c *ListCmd) Run() error {
	// List all smartcards connected to the system.
	cards, err := piv.Cards()
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Ref", "Card"})
	t.AppendSeparator()

	for k, v := range cards {
		t.AppendRow(table.Row{k, v})
	}

	t.Render()

	return nil
}
