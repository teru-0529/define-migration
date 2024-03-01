/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "migration file copy to sorce repository.",
	Long:  "migration file copy to sorce repository.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mirror called")
	},
}

func init() {
}
