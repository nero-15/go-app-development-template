# go-app-development-template
go言語でアプリケーション開発する際のテンプレート

## 概要
すぐに開発が始められる環境を構築するためのフローをまとめています。

## 環境
- go1.14.3

## やり方

1. 初期設定
```
mkdir {myapp}
go mod init　{myapp}
```

2. 該当パッケージインストール
```
go get github.com/labstack/echo/v4
```
[echo](https://echo.labstack.com/guide/){:target="_blank"}
