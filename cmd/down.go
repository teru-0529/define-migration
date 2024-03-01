/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "migration down.",
	Long:  "migration down.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("down called")
	},
}

func init() {
}
