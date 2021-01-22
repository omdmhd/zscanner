package cmd

import (
	"fmt"
	"strconv"
	"zscanner/modules/dns"

	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "zscanner 1.0.0 is a network scanner",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			start := 0
			end := 3
			if len(args) > 1 {
				start, _ = strconv.Atoi(args[1])
			}
			if len(args) > 2 {
				end, _ = strconv.Atoi(args[1])
			}
			dns.Execute(args[0], start, end)
		} else {
			fmt.Println("specify the domain name")
		}
	},
}
