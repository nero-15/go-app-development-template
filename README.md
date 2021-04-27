# go-app-development-template
go言語でアプリケーション開発する際のテンプレート

## 概要
すぐに開発が始められる環境を構築するためのフローをまとめています。

## 環境
- go1.14.3
- mysql 8.0

## やり方

1. ファイルの追加
このリポジトリのファイルをプロジェクトにzipファイル等でコピー

2. 該当パッケージインストール
```
go get github.com/labstack/echo/v4
go get gopkg.in/ini.v1
go get github.com/go-sql-driver/mysql
```
[echo](https://echo.labstack.com/guide/) <br>
[gopkg.in/ini.v1](https://pkg.go.dev/gopkg.in/ini.v1) <br>
[go-sql-driver](github.com/go-sql-driver/mysql)

3. config.iniファイルを作成
```
cd struttura/config
mv config.ini.default config.ini
```
DBを使う場合はconfig.iniの該当の情報を追加

host = DBのホスト名
name = DB名
password = パスワード

4. 開発を始めよう！

