/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"errors"
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
			return errors.New("please input schemaName for args[0]")
		}
		schemaName := args[0]

		// INFO: スキーマの存在確認
		hasSource := sources.Exist(schemaName)
		if !hasSource {
			return fmt.Errorf("schema[\"%s\"] is not exist for sourceFile", schemaName)
		}

		// INFO: 情報表示
		fmt.Printf("schemaName: %s\n", schemaName)
		fmt.Printf("[source  ] %s\n", sources.SourceUrl(schemaName))
		fmt.Printf("[database] %s\n", postgres.DatabaseUrl(schemaName))
		fmt.Println(useLocal)

		return nil
	},
}

func init() {
}
