package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func GenerateFile(schemaName string, fileDescription string) {

	dir := filepath.Join("src", schemaName, sourceDir)

	// INFO: 連番取得
	num := generateNumber(dir)
	fmt.Println("generate migration file")

	// INFO: 新規ファイル作成
	for _, token := range []string{"up", "down"} {
		filePath := filepath.Join(dir, fmt.Sprintf("%06d_%s.%s.sql", num, fileDescription, token))
		generateEmptyFile(filePath)
		fmt.Printf("  : %s\n", filePath)
	}
}

// 空ファイルを作成する（フォルダが無ければ作成する）
func generateEmptyFile(fileName string) error {
	dir := filepath.Dir(fileName)

	// INFO: フォルダが存在しない場合作成する
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return fmt.Errorf("cannot create directory: %s", err.Error())
		}
	}

	// INFO: 出力用ファイルのオープン
	d := []byte("")
	err := os.WriteFile(fileName, d, 0644)
	if err != nil {
		return fmt.Errorf("cannot create file: %s", err.Error())
	}
	return nil
}

// 最新のマイグレーション番号を生成する
func generateNumber(dir string) int {
	// 配下のファイルをリスト化
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Warning: %s\n", err)
		return 1
	}

	// 000000_{*}.up.sqlのファイルを抽出
	correctFileName := regexp.MustCompile(`^[0-9]{6}_(.*).up.sql$`)
	entries = lo.Filter(entries, func(item os.DirEntry, num int) bool { return correctFileName.MatchString(item.Name()) })
	if len(entries) == 0 {
		return 1
	}

	// リストを降順でソート
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() > entries[j].Name() })

	// 1件目の要素の最初の6文字を抽出し、数値型に変換して返す
	num, err := strconv.Atoi(entries[0].Name()[:6])
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return num + 1
}
