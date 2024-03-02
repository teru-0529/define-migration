/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show migration source and database info.",
	Long:  "show migration source and database info.",
	RunE: func(cmd *cobra.Command, args []string) error {

		// INFO: 引数の入力確認
		if len(args) < 1 {
			fmt.Println("please input schema name for args[0].")
			return nil
		}
		schemaName := args[0]

		// INFO: スキーマの存在確認
		if !schemas.Exist(schemaName) {
			fmt.Printf("schema[\"%s\"] is not exist for schema file.\n", schemaName)
			return nil
		}

		fmt.Println(schemaName)
		fmt.Println(postgres)
		fmt.Println(github)
		fmt.Println(schemas)
		return nil
	},
}

func init() {
}
