package check

import (
	"github.com/cagiti/yb/pkg/yubikey"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Cmd struct {
	Cmd  *cobra.Command
	Args []string
}

func NewCheckCmd() *cobra.Command {
	c := &Cmd{}
	cmd := &cobra.Command{
		Use:     "check",
		Short:   "checks the selected yubikey",
		Long:    "",
		Example: "",
		Aliases: []string{"c", "ch", "che", "chec", "check"},
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
	y := yubikey.NewYubikey()
	err := y.SelectYubikey()
	if err != nil {
		return err
	}
	err = y.Check()
	if err != nil {
		return err
	}
	return nil
}
