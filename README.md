# ShiftManagerServer
> ここのディレクトリにサーバーサイドのコードをおいて、完全に分離します。

シフト管理ツールのAPI

APIのドキュメントは`HackMD`に書いたのでの参照してください（未完成）</br>
[ShiftManager API ドキュメント](https://hackmd.io/@R1zb_r8nS2SNRHkRjV6SoA/SJ0ZtA23O)

## ディレクトリ構成について
- service : modelの動作を実装
- server : ginの設定（ルーティングなど）
- controller : アクションを定義しserviceを呼び出す
- entity : modelのstructを分離す）
- db : データベース設定
- vendor : Goで用いるパッケージなど

## 認証について
使いやすいのでJWT認証にします。
それぞれのヘッダーにAccessTokenを付与してください

## Database
開発環境のデータベースを動かし方
```bash
$ go run main.go

$ docker-compose up -d

$ qspl -h 172.29.48.1 -U gorm gorm
gorm=####
```

## 開発ルール
`develop`のブランチを開発環境、`main`を本番環境と同じにする。

開発は、次の流れに従う
1. issueを作る。または、確認する。
2. 担当するissueに自分を割り当てる。
3. issueごとにブランチを切る
4. コードが完成したら、Pull Requestを出す
5. Pull Requestを誰かがレビュー（今回は省略）
6. `develop`ブランチにマージする
7. issueを閉じる

--- 

基本的に次の流れでコマンドを打つ
```bash
% git checkout -b issue/?
% git add .
% git commit -m "commit message"
% git push origin HEAD
```
