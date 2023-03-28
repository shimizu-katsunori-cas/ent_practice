# entの技術調査

## 使い方
### スキーマの作成
* スキーマの定義
```
ent new User
```
→Userに関するスキーマを定義

* スキーマの生成
```
ent generate ./ent/schema
```
→スキーマファイルを修正する度に実行が必要

* スキーマの修正
ent_practice/ent/schema以下のファイルを修正

## 調査方針
- [ ] 生のクエリ作成
  - [ ] select・update・insert・deleteが利用できるか
  - [ ] AND, ORが利用可能か
  - [ ] in句が利用可能か
  - [ ] LIKEが利用可能か
  - [ ] どの型は値を渡せるのか？(int, string, slice, struct)
  - [ ] 演算子が利用可能か(<, >, =)
  - [ ] if文が書けるか
  - [ ] プレースホルダが利用可能か(named parameter)(:)
  - [ ] 生のクエリからコードに変換する大変さ、テストが楽かどうか
  - [ ] schemaからモデルを生成してくれるかどうか
