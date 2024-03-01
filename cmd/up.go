/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "migration up.",
	Long:  "migration up.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called")
	},
}

func init() {
}
