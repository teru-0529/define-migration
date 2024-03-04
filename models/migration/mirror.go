package migration

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// マイグレーションファイルのコピー
func MirrorMigration(schemaName string, distRoot string) error {

	srcDir := filepath.Join("src", schemaName, sourceDir)
	distDir := filepath.Join(distRoot, sourceDir)

	fmt.Println(schemaName)
	fmt.Println(srcDir)
	fmt.Println(distDir)

	// INFO: フォルダが存在しない場合終了
	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("Warning: migration file of schema[\"%s\"] is empty.", schemaName)
		return nil
	}

	// INFO: フォルダが存在しない場合作成する
	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		if err := os.MkdirAll(distDir, 0777); err != nil {
			return fmt.Errorf("cannot create directory: %s", err.Error())
		}
	}

	// 配下のファイルをリスト化
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		newFile, err := os.Create(filepath.Join(distDir, entry.Name()))
		if err != nil {
			return err
		}

		oldFile, err := os.Open(filepath.Join(srcDir, entry.Name()))
		if err != nil {
			return err
		}

		io.Copy(newFile, oldFile)
		fmt.Printf("Info: copy file[\"%s\"]\n", entry.Name())
	}

	return nil
}
