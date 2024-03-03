/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/teru-0529/define-migration/models/migration"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "migration up.",
	Long:  "migration up.",
	RunE: func(cmd *cobra.Command, args []string) error {

		// INFO: 対象スキーマの選定
		var schemas []string
		if len(args) > 0 {
			schemas = sources.CheckSchemas(args)
		} else {
			schemas = sources.Schemas()
		}

		// INFO: migration実行
		for _, schema := range schemas {
			migrationIO := migration.MigrationIO{
				Schema:   schema,
				Source:   sources.SourceUrl(schema),
				Database: postgres.DatabaseUrl(schema),
			}
			err := migration.MigrateUp(migrationIO)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
}
