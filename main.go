package main

import (
  _"github.com/go-sql-driver/mysql" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
)

// 各種処理
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("views/*.html")

  dbInit()

  // インデックスページのルーティング
  router.GET("/", func(c *gin.Context) {
   tweets := dbGetAll()
   c.HTML(200, "index.html", gin.H{"tweets": tweets})
  })

  // POSTデータを受け取ってDBに登録する
  router.POST("/new", func(c *gin.Context) {
   content := c.PostForm("content")
   dbInsert(content)
   c.Redirect(302, "/")
  })

  router.Run()
}

// モデルの宣言
type Tweet struct {
  gorm.Model
  // カラム名 型
  Content string
}

// DBの初期化
func dbInit() {
  // MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
  db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/app_like_twitter_by_go?parseTime=true") 
  if err != nil {
   panic("DBを開けません (dbInit())")
  }
  //宣言に基づいてテーブルを作成
  db.AutoMigrate(&Tweet{})
  // コネクション解放
  defer db.Close()
}

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
