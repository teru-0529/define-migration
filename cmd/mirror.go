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

// コピー先リポジトリルート
var repositoryRoot string

// mirrorCmd represents the mirror command
var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "migration file copy to sorce repository.",
	Long:  "migration file copy to sorce repository.",
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

		// INFO: コピー処理
		if err := migration.MirrorMigration(schemaName, repositoryRoot); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	mirrorCmd.Flags().StringVarP(&repositoryRoot, "repositoryRoot", "D", "", "repository root dir to mirror migration file.")

	mirrorCmd.MarkFlagRequired("repositoryRoot")
}
