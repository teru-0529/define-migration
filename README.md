# define-migration

テーブルマイグレーションアプリ

## 参考

* go buildでファイル名指定
  `go build -o xxx.exe`
  https://qiita.com/kanuma1984/items/80b1f4c35ba847bb84a2

* go buildで-ldflags
  `go build -ldflags '-X main.version=1.0.0 -X main.date=2024-03-02T12:12:12'`
  https://kazuhira-r.hatenablog.com/entry/2021/03/08/003752

* goreleaser ローカルビルド
  `goreleaser check`
  `goreleaser release --snapshot --clean`
  https://zenn.dev/kou_pg_0131/articles/goreleaser-usage

* goreleaserでファイル名変える

  ``` yaml
  builds:
    binary: dx2-migration
  ```

  https://goreleaser.com/customization/builds/

* postgres　バージョン
  https://hub.docker.com/_/postgres

* golang-migrateでのPostgreSQLマイグレーション備忘録
  https://zenn.dev/keyamin/articles/24695c455c1591
