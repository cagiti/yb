package cmd

import (
	"github.com/cagiti/yb/pkg/prompter"
)

type CommonOptions struct {
	prompter prompter.Prompter
}

func (c *CommonOptions) SetPrompter(p prompter.Prompter) {
	c.prompter = p
}

func (c *CommonOptions) Prompter() prompter.Prompter {
	if c.prompter == nil {
		c.prompter = prompter.NewPrompter()
	}
	return c.prompter
}
