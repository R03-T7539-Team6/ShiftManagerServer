# ShiftManagerServer
> ここのディレクトリにサーバーサイドのコードをおいて、完全に分離します。

シフト管理ツールのAPI



## エンドポイント（仮）
```
GET /users/
GET /users/{userId}
POST /users/
PUT /users/{userId}
DELETE /users/{userId}
POST /login
```

## ディレクトリ構成について
- service : modelの動作を実装
- server : ginの設定（ルーティングなど）
- controller : アクションを定義しserviceを呼び出す
- entity : modelのstructを分離する
- db : データベース設定
- vendor : Goで用いるパッケージなど

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
