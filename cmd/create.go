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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create migration file.",
	Long:  "create migration file.",
	RunE: func(cmd *cobra.Command, args []string) error {

		// INFO: 引数の入力確認
		if len(args) < 2 {
			return errors.New("please input schemaName for args[0] and fileDescription for args[1]")
		}
		schemaName := args[0]
		fileDescription := args[1]

		// INFO: スキーマの存在確認
		hasSource := sources.Exist(schemaName)
		if !hasSource {
			return fmt.Errorf("schema[\"%s\"] is undefined for sourceFile", schemaName)
		}

		// INFO: マイグレーションファイルの作成
		migration.GenerateFile(schemaName, fileDescription)

		return nil
	},
}

func init() {
}
