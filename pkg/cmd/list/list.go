package list

import (
	"os"

	"github.com/cagiti/yb/pkg/yubikey"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Cmd struct {
	Cmd  *cobra.Command
	Args []string
}

func NewListCmd() *cobra.Command {
	c := &Cmd{}
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

func (c *Cmd) Run() error {
	// List all connected yubikeys.
	ybs, err := yubikey.ListYubikeys()
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Ref", "Card"})
	t.AppendSeparator()

	for k, v := range ybs {
		t.AppendRow(table.Row{k, v})
	}

	t.Render()

	return nil
}
