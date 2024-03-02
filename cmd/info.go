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
			return errors.New("please input schema name for args[0]")
		}
		schemaName := args[0]

		sourceUrl, hasSource := sources.SourceUrl(schemaName, sourceType)
		// INFO: スキーマの存在確認
		if !hasSource {
			return fmt.Errorf("schema[\"%s\"] is not exist for schema file", schemaName)
		}

		// INFO: 情報表示
		fmt.Printf("schema name: %s\n", schemaName)
		fmt.Printf("[source  ] %s\n", sourceUrl)
		fmt.Printf("[database] %s\n", postgres.DatabaseUrl(schemaName))

		return nil
	},
}

func init() {
}
