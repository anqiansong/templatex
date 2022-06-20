package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/anqiansong/templatex/pkg/cobrax"
	"github.com/anqiansong/templatex/pkg/exec"
	"github.com/spf13/cobra"
)

func main() {
	var cmd = cobra.Command{
		Use: "custom",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use: "join",
			Run: func(cmd *cobra.Command, args []string) {
				var v []string
				if err := exec.Unmarshal(args[0], &v); err != nil {
					log.Fatalln(err)
				}

				fmt.Println(strings.Join(v[0:len(v)-1], v[len(v)-1]))
			},
		},
	)
	cmd.AddCommand(cobrax.TemplatexCommand)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
