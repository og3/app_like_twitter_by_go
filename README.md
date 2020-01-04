# app_like_twitter_by_go
これで作る
https://blog.kannart.co.jp/programming/2026/
## DB接続について
```
DBへの接続
gormでのDBの接続はほぼ以下の書式で行われます。

db, err := gorm.Open(“mysql”, “[ユーザー名]:@[ホスト名]/[DB名]?parseTime=true”)
if err != nil {
    panic(“You can’t open DB (dbGetAll())”)
} defer db.Close()

今回の注意点としては、?parseTime=trueです。

MySQLの場合は文字コードだかの関係で、これを付けないと現状接続できません。
```
mysqlだから特別なことをしないといけないわけじゃないらしい。。

## DB処理の雛形
- モデル
```
// モデルの宣言
type Tweet struct {
	gorm.Model
	// カラム名 型
	Content string
}
```
- DB初期化
```
// DBの初期化
func dbInit() {
  // MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/app_like_twitter_by_go?parseTime=true") 
	if err != nil {
		panic("DBが開けません (dbInit())")
	}
	//構造体に基づいてテーブルを作成
	db.AutoMigrate(&Tweet{})
  // コネクション解放
	defer db.Close()
}
```
この処理でDB作成をしてくれるのは初回のみの模様。。どっかに読み先情報をキャッシュしてるとか？

- post
```
// DB追加
func dbInsert(content string) {
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/app_like_twitter_by_go?parseTime=true")
  if err != nil {
    panic("DBを開けません (dbInsert())")
  }
  // Insert処理
  db.Create(&Tweet{Content: content})
  defer db.Close()
}
```
- all
```
// 全件取得
func dbGetAll() []Tweet {
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/app_like_twitter_by_go?parseTime=true")
  if err != nil {
   panic("DBを開けません (dbGetAll())")
  }
  var tweets []Tweet
  // Findでテーブル名を指定して取得した後、orderで登録順に並び替え
  db.Order("created_at desc").Find(&tweets)
  defer db.Close()
  return tweets
}
```
- find
```
//DB一つ取得
func dbGetOne(id int) Todo {
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo_app_by_go?parseTime=true")
  if err != nil {
      panic("データベース開けず！(dbGetOne())")
  }
  var todo Todo
  db.First(&todo, id)
  db.Close()
  return todo
}
```
- update
```
//DB更新
func dbUpdate(id int, text string, status string) {
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo_app_by_go?parseTime=true")
  if err != nil {
      panic("データベース開けず！（dbUpdate)")
  }
  var todo Todo
  db.First(&todo, id)
  todo.Text = text
  todo.Status = status
  db.Save(&todo)
  db.Close()
}
```
- delete
```
//DB削除
func dbDelete(id int) {
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo_app_by_go?parseTime=true")
  if err != nil {
      panic("データベース開けず！（dbDelete)")
  }
  var todo Todo
  // 削除対象の読み込み
  db.First(&todo, id)
  // 削除
  db.Delete(&todo)
  db.Close()
}
```

## メモ
- DB処理でerrの出力条件が!=nilだから実はエラー文が入ってるのか？pry的なもので中身を見てみたい。。

## 参考
- gorm
http://gorm.io/ja_JP/
