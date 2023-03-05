# IVIS NAS Internal Web Server
[![ko](https://img.shields.io/badge/lang-ko-red.svg)](https://github.com/picel/IVIS_NAS/blob/main/README.md)
[![ja](https://img.shields.io/badge/lang-ja-blue.svg)](https://github.com/picel/IVIS_NAS/blob/main/README.jp.md)

## 開発環境
- Go
- MXRoute Mail API (DirectAdmin API)
- SMB
- AFP
- NGINX (Reverse Proxy)

## ルート情報
- GET /
    - NASユーザー追加ページ
- POST /process
    - NASユーザー追加処理ページ
    - UNIXアカウント作成
    - SMBアカウント作成
    - MXRoute Mailユーザー自動追加
- POST /loginCheck
    - Helperサーバーからユーザー認証要求用API
    - Helperサーバーからユーザー認証要求時、UNIXアカウントの存在有無を確認して認証結果を返す。
    - /etc/shadowファイルを読み込んで認証処理。

## 実行画面
- NASユーザー追加ページ
![NASユーザー追加ページ](https://user-images.githubusercontent.com/30901178/222890273-194bfde2-8ca2-4c21-b972-72a6d42de615.png)
- NASユーザー登録結果ページ
![NASユーザー登録結果ページ](https://user-images.githubusercontent.com/30901178/222890305-ff4eb233-0a31-48ec-b486-c6921b474bb3.png)

## 注意事項
- keyファイルはgitにアプロードしない。
    - keyファイルは/etc/goweb/keysディレクトリに配置。
    - MXRouteURL, MXRouteID, MXRoutePW定義