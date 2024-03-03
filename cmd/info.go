/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/teru-0529/define-migration/models/migration"
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
			return fmt.Errorf("schema[\"%s\"] is undefined for sourceFile", schemaName)
		}

		// INFO: 情報表示
		fmt.Printf("schemaName: %s\n", schemaName)
		migrationIO := migration.MigrationIO{
			Schema:   schemaName,
			Source:   sources.SourceUrl(schemaName),
			Database: postgres.DatabaseUrl(schemaName),
		}
		migrationIO.Show()

		return nil
	},
}

func init() {
}
