// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package command

import (
	"fmt"
	"strings"
)

type DidoCommand struct {
	Meta
}

func (c *DidoCommand) Run(args []string) int {
	fmt.Println("Se FODEU")
	return 0
}

func (c *DidoCommand) Help() string {
	helpText := `
Usage: tofu [global options] dido [options] [DIR]

  This command fucks you in every sense.
`
	return strings.TrimSpace(helpText)
}

func (c *DidoCommand) Synopsis() string {
	return "This command fucks you in every sense."
}
