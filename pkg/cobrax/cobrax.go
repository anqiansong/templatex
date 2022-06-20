package cobrax

import (
	"fmt"

	"github.com/anqiansong/templatex/pkg/exec"
	"github.com/spf13/cobra"
)

var TemplatexCommand = &cobra.Command{
	Use:    "templatex",
	Short:  "Print all the sub-commands for templatex",
	Long:   "This command automatically generates a templatex command with all the sub-commands",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		root := cmd.Root()
		var list []string
		for _, sub := range root.Commands() {
			if sub.Name() == "help" || sub.Hidden || sub.Use == "templatex" || sub.Use == "completion" {
				continue
			}
			list = append(list, sub.Name())
		}
		data, _ := exec.Marshal(list)
		fmt.Println(data)
	},
}
